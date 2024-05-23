package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/julebarn/DIS_Project/server/auth"
	"github.com/julebarn/DIS_Project/server/db"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./build")))
	http.Handle("api/auth", auth.EndpointsHandler())

	http.HandleFunc("/api/event/details/{id}", func(w http.ResponseWriter, r *http.Request) {

		id := r.URL.Query().Get("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		event, err := db.New(db.Conn()).GetEvent(r.Context(), int32(idInt))
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(EventResponse{
			ID:          event.ID,
			Name:        event.Name,
			Place:       event.Place,
			Description: event.Description,
			StartTime:   event.StartTime,
			EndTime:     event.EndTime,
			ClubID:      event.ClubID,
		})
	})

	http.HandleFunc("/api/event/future", func(w http.ResponseWriter, r *http.Request) {
		event, err :=  db.New(db.Conn()).GetFutureEvents(r.Context())
		if err != nil {
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

	http.Handle("/api/event/create",
		auth.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var req CreateEventRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			userID := r.Context().Value(auth.UserIDKey).(int)

			// Create event
			err = db.New(db.Conn()).CreateEvent(r.Context(), db.CreateEventParams{
				Name:        req.Name,
				Place:       req.Place,
				Description: req.Description,
				StartTime:   req.StartTime,
				EndTime:     req.EndTime,
				ClubID:      req.ClubID,
				UserID:      int32(userID),
			})
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}



		})))

	http.ListenAndServe(":8080", nil)
}


// maybe dont use pgtype in the request struct -mads 
// ps. i wrote the code useing pgtype
type CreateEventRequest struct {
	Name        string           `json:"name"`
	Place       string           `json:"place"`
	Description string           `json:"description"`
	StartTime   pgtype.Timestamp `json:"start_time"`
	EndTime     pgtype.Timestamp `json:"end_time"`
	ClubID      pgtype.Int4      `json:"club_id"`
}


type EventResponse struct {
	ID          int32            `json:"id"`
	Name        string           `json:"name"`
	Place       string           `json:"place"`
	Description string           `json:"description"`
	StartTime   pgtype.Timestamp `json:"start_time"`
	EndTime     pgtype.Timestamp `json:"end_time"`
	ClubID      pgtype.Int4      `json:"club_id"`
}

