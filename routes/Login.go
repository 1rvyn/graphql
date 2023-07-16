package routes

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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

		// Here, you should fetch the user from your database. For the purpose of this example,
		// let's just use a dummy user.
		user := User{
			Username: "test",
			Password: "$2a$10$N9qo8uLOickgx2ZMRZoHK.ApicY1A9OP7T3Q/SB0x61A5x8C9XKa", // password is "password"
		}

		// Compare the password with the hashed password stored in the database
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": creds.Username,
			// You can add more claims here, like "role": user.Role
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
	}
}
