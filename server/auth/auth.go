package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julebarn/DIS_Project/server/db"
	"golang.org/x/crypto/bcrypt"
)

// top secret key do not share with anyone \s
var secretKey = []byte("secret-key")

type key int
const UserIDKey key = 0


// Middleware checks if the user is authenticated
// if the user is authenticated the user id is added to the (request) context 
// uder the key UserIDKey
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, userID := isAuth(r)

		if !auth {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		r = r.WithContext(ctx)
		
		next.ServeHTTP(w, r)

	})
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func EndpointsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if strings.HasSuffix(r.URL.Path, "/login") {
			var req loginRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			user, err := db.New(db.Conn()).GetUser(ctx, req.Username)
			if err != nil {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(user.Username), []byte(req.Password))
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			cookie, err := tokenCookie(user.ID)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, cookie)

		} else if strings.HasSuffix(r.URL.Path, "/register") {

			var req registerRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			user, err := db.New(db.Conn()).CreateUser(ctx, db.CreateUserParams{
				Username:     req.Username,
				Passwordhash: string(hash),
			})

			cookie, err := tokenCookie(user.ID)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, cookie)
			
		} else if strings.HasSuffix(r.URL.Path, "/logout") {
			
			// for reference you can logout with out being logged in 😂
			
			http.SetCookie(w, &http.Cookie{
				Name:   "token",
				Value: "",
				MaxAge: -1,
				Path:   "/",
			})

		} else if strings.HasSuffix(r.URL.Path, "/refresh") {

			auth, userid := isAuth(r)
			if !auth {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			cookie, err := tokenCookie(userid)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, cookie)

		} else {
			http.Error(w, "Not found", http.StatusNotFound)
		}
	})
}

func tokenCookie(userID int32) (*http.Cookie, error) {
	expirationTime := time.Now().Add(30 * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":  userID,
		"exp": expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return nil, err
	}

	// using cookie to store token is not perfect, but it is good enough for now
	// it might be possible to the Authorization header instead (i'm not sure about this)
	// it could also be to use the session store (in the browser) to store the token
	// but i don't want to to be in js for now - mads
	cookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    "/",
	}
	return cookie, nil
}


func isAuth(r *http.Request) (isAuth bool, userid int32) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return false, 0
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return false, 0
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, 0
	}

	userid, ok = claims["ID"].(int32)
	if !ok {
		return false, 0
	}

	return true, userid
}