package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/julebarn/DIS_Project/server/auth"
	"github.com/julebarn/DIS_Project/server/db"
)

func main() {

	handler := http.NewServeMux()
	handler.Handle("/", http.FileServer(http.Dir("./build")))

	handler.HandleFunc("/api/event/details/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hej")

		id := r.PathValue("id")

		fmt.Println(id)

		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		event, err := db.New(db.Conn(r.Context())).GetEvent(r.Context(), int32(idInt))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		organizers, err := db.New(db.Conn(r.Context())).GetOrganizers(r.Context(), event.ID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		var res []struct {
			ID   int32  `json:"id"`
			Name string `json:"name"`
		}
		for _, o := range organizers {
			res = append(res, struct {
				ID   int32  `json:"id"`
				Name string `json:"name"`
			}{
				ID:   o.ID,
				Name: o.Username,
			})
		}

		json.NewEncoder(w).Encode(EventResponse{
			ID:          event.ID,
			Name:        event.Name,
			Place:       event.Place,
			Description: event.Description,
			StartTime:   event.StartTime,
			EndTime:     event.EndTime,
			ClubID:      event.ClubID,
			Organizers:  res,
		})
	})

	handler.HandleFunc("/api/event/future", func(w http.ResponseWriter, r *http.Request) {
		event, err := db.New(db.Conn(r.Context())).GetFutureEvents(r.Context())
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		var res []EventResponse
		for _, e := range event {
			res = append(res, EventResponse{
				ID:          e.ID,
				Name:        e.Name,
				Place:       e.Place,
				Description: e.Description,
				StartTime:   e.StartTime,
				EndTime:     e.EndTime,
				ClubID:      e.ClubID,
			})
		}

		json.NewEncoder(w).Encode(res)

	})

	handler.HandleFunc("/api/club/details/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		fmt.Println(id)

		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		DB := db.New(db.Conn(r.Context()))

		club, err := DB.GetClub(r.Context(), int32(idInt))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		managers, err := DB.GetManagers(r.Context(), club.ID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		ClubResponse := ClubResponse{
			ID:          club.ID,
			Name:        club.Name,
			Description: club.Description,
		}

		for _, m := range managers {
			ClubResponse.Managers = append(ClubResponse.Managers, struct {
				ID   int32  `json:"id"`
				Name string `json:"name"`
			}{
				ID:   m.ID,
				Name: m.Username,
			})
		}

		json.NewEncoder(w).Encode(ClubResponse)
	})

	handler.HandleFunc("/api/club/list", func(w http.ResponseWriter, r *http.Request) {
		club, err := db.New(db.Conn(r.Context())).GetClubs(r.Context())
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		var res []db.Club
		for _, c := range club {
			res = append(res, c)
		}

		json.NewEncoder(w).Encode(res)
	})

	handler.HandleFunc("/api/user/list", func(w http.ResponseWriter, r *http.Request) {
		user, err := db.New(db.Conn(r.Context())).GetAllUsers(r.Context())
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		var res []struct {
			ID       int32  `json:"id"`
			Username string `json:"username"`
		}
		for _, u := range user {
			res = append(res, struct {
				ID       int32  `json:"id"`
				Username string `json:"username"`
			}{
				ID:       u.ID,
				Username: u.Username,
			})
		}

		json.NewEncoder(w).Encode(res)
	})

	handler.Handle("/api/event/create",
		auth.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("create event")
			var req CreateEventRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			fmt.Println(req.Name)
			fmt.Println(req.Place)
			fmt.Println(req.Description)
			fmt.Println(req.StartTime)
			fmt.Println(req.EndTime)
			fmt.Println(req.ClubID)

			//2006-01-02T15:04

			// time location for copenhagen
			location := time.FixedZone("CET", 1*60*60)

			startTime, startErr := time.ParseInLocation("2006-01-02T15:04", req.StartTime, location)
			endTime, endErr := time.ParseInLocation("2006-01-02T15:04", req.EndTime, location)

			if startErr != nil || endErr != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			userID := r.Context().Value(auth.UserIDKey).(int32)

			EventParams := db.CreateEventParams{
				Name:        req.Name,
				Place:       req.Place,
				Description: req.Description,
				StartTime:   pgtype.Timestamp{Time: startTime.UTC(), Valid: true},
				EndTime:     pgtype.Timestamp{Time: endTime.UTC(), Valid: true},
				ClubID:      req.ClubID,
				UserID:      int32(userID),
			}

			// Create event
			err = db.New(db.Conn(r.Context())).CreateEvent(r.Context(), EventParams)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		})))

	handler.Handle("/api/club/create",
		auth.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("create event")
			var req CreateEventRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			fmt.Println(req.Name)
			fmt.Println(req.Description)

			userID := r.Context().Value(auth.UserIDKey).(int32)

			// Create event
			err = db.New(db.Conn(r.Context())).CreateClub(r.Context(), db.CreateClubParams{
				Name:        req.Name,
				Description: req.Description,
				UserID:      int32(userID),
			})
			if err != nil {
				fmt.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		})))

	handler.Handle("/api/club/isOrganizer",
		auth.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userID := r.Context().Value(auth.UserIDKey).(int32)

			club, err := db.New(db.Conn(r.Context())).GetClubByManagers(r.Context(), userID)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(club)
		})))

	handler.Handle("/api/club/addManager",
		auth.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			var req struct {
				ClubID int32 `json:"club"`
				UserID int32 `json:"manager"`
			}

			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			fmt.Println(req.ClubID)
			fmt.Println(req.UserID)

			err = db.New(db.Conn(r.Context())).AddManager(r.Context(), db.AddManagerParams{
				ClubID: req.ClubID,
				UserID: req.UserID,
			})
			if err != nil {
				fmt.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}


			
		})))

	handler = auth.EndpointsHandler(handler)

	http.ListenAndServe(":8080", newhanderloger(handler))
}

// maybe dont use pgtype in the request struct -mads
// ps. i wrote the code useing pgtype
type CreateEventRequest struct {
	Name        string      `json:"name"`
	Place       string      `json:"place"`
	Description string      `json:"description"`
	StartTime   string      `json:"start_time"`
	EndTime     string      `json:"end_time"`
	ClubID      pgtype.Int4 `json:"club_id"`
}

type EventResponse struct {
	ID          int32            `json:"id"`
	Name        string           `json:"name"`
	Place       string           `json:"place"`
	Description string           `json:"description"`
	StartTime   pgtype.Timestamp `json:"start_time"`
	EndTime     pgtype.Timestamp `json:"end_time"`
	ClubID      pgtype.Int4      `json:"club_id"`
	Organizers  []struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	} `json:"organizers"`
}

type ClubResponse struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Managers    []struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	} `json:"managers"`
}

func newhanderloger(next http.Handler) *handerloger {
	return &handerloger{
		handler: next,
	}
}

type handerloger struct {
	handler http.Handler
}

func (h *handerloger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.Method)

	for _, Cookie := range r.Cookies() {
		fmt.Print(Cookie.Name)
		fmt.Print(": ")
		fmt.Print(Cookie.Value)
		fmt.Println()
	}

	h.handler.ServeHTTP(w, r)
}
