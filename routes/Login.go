package routes

import (
	"crypto/subtle"
	"encoding/json"
	"net/http"
	"time"

	"github.com/1rvyn/graphql-service/database"
	"github.com/1rvyn/graphql-service/models"
	"github.com/1rvyn/graphql-service/utils"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(router *mux.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the JSON body from the request
		var creds Credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Fetch the user from the database
		var user models.Employee
		result := database.Database.Db.Where("username = ?", creds.Username).First(&user)
		if result.Error != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		// Compare the provided password with the hashed password stored in the database
		hashedPassword := utils.HashPassword(creds.Password)

		if subtle.ConstantTimeCompare(hashedPassword, user.Password) == 0 {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": creds.Username,
			"exp":      jwt.TimeFunc().Add(time.Hour * 24).Unix(),
			"position": user.Position, // this can allow us to be very custom in the middleware ie managers get higer perms that software engineers
		})

		// Create the JWT string
		tokenString, err := token.SignedString([]byte("your-secret-key"))
		if err != nil {
			http.Error(w, "Error in generating token", http.StatusInternalServerError)
			return
		}

		// Finally, we set the client cookie for "token" as the JWT we just generated
		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: tokenString,
		})

		// return a success message
		w.Write([]byte("Login successful"))
	}
}
