package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"../entities"
)

// var DB *gorm.DB

// const conn = "root:QSGGSQ139@tcp(127.0.0.1:3306)/StudentDB"

var secretKey = []byte("secret_key")

var User []entities.Users

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type SampUser struct {
	UserName string `json:"userName"`
	Password string `json:"passWord"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")

	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}
	var user SampUser
	_ = json.NewDecoder(r.Body).Decode(&user)

	stmt := fmt.Sprintf(`call Std_admin_login('%v','%v')`, user.UserName, user.Password)
	rows, err := db.Queryx(stmt)
	if err != nil {
		log.Fatal(err)
	} else {
		for rows.Next() {
			var userId string
			var userName string
			var passWord string
			var firstName string
			var lastName string
			var contactNumber string
			var email string

			err2 := rows.Scan(&userId, &userName, &passWord, &firstName, &lastName, &contactNumber, &email)
			if err2 != nil {
				log.Fatal(err)
			} else {
				user := entities.Users{userId, userName, passWord, firstName, lastName, contactNumber, email}

				User = append(User, user)
			}
		}
	}

	if User == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println(user)
	fmt.Println(User)
	for _, value := range User {
		if user.Password != value.Password {

			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
	// fmt.Println("stopped here")
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	fmt.Println("token is created")
	User = nil
}

func LoggedUser(w http.ResponseWriter, r *http.Request) {

	// userName, _ := ExtractTokenUsername(r)
	// for _, value := range User {
	// 	if userName != value.UserName {
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	} else {
	// 		w.WriteHeader(http.StatusAccepted)
	// 	}
	// }

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")

	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))

}

func LogOut(w http.ResponseWriter, r *http.Request) {

	c := http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &c)
	w.Write([]byte("Logged Out!\n"))
	fmt.Println("LoggedOut")
}

// func ExtractTokenUsername(r *http.Request) (string, error) {

// 	// get our token string from Cookie
// 	biscuit, err := r.Cookie("token")

// 	var tokenString string
// 	if err != nil {
// 		tokenString = ""
// 	} else {
// 		tokenString = biscuit.Value
// 	}

// 	// abort
// 	if tokenString == "" {
// 		return "", nil
// 	}

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(secretKey), nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		username := fmt.Sprintf("%s", claims["username"]) // convert to string
// 		if err != nil {
// 			return "", err
// 		}
// 		return username, nil
// 	}
// 	return "", nil
// }
