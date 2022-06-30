package main

import (
	"log"
	"net/http"

	"./controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// var DB *gorm.DB

// func Connect() {
// 	connection, err := gorm.Open(mysql.Open("root:QSGGSQ139@tcp(127.0.0.1:3306)/StudentDB"), &gorm.Config{})

// 	if err != nil {
// 		panic("could not connect to the database")
// 	}

// 	DB = connection

// 	connection.AutoMigrate(&entities.Users{})
// }

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/students", controllers.GetStudents).Methods("GET")
	r.HandleFunc("/students/{studentId}", controllers.GetStudent).Methods("GET")
	r.HandleFunc("/students", controllers.CreateStudents).Methods("POST")
	r.HandleFunc("/students/{studentId}", controllers.UpdateStudents).Methods("PUT")
	r.HandleFunc("/students/{studentId}", controllers.DeleteStudents).Methods("DELETE")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/user", controllers.LoggedUser).Methods("GET")
	r.HandleFunc("/logout", controllers.LogOut).Methods("POST")
	// r.HandleFunc("/file/{studentId}", controllers.DownloadFile).Methods("GET")
	// r.HandleFunc("/refresh", controllers.Refresh).Methods("POST")
	// app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))

	// Setup(app)

	// app.Listen(":8000")
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(handlers.AllowedMethods([]string{"HEAD", "GET", "POST", "PUT", "DELETE"}), handlers.AllowedOrigins([]string{"*"}), handlers.AllowedHeaders([]string{"authentication", "Content-Type"}), handlers.AllowCredentials())(r)))
}

// func Setup(app *fiber.App) {

// 	// app.Post("/api/register", controllers.Register)
// 	app.Post("/api/login", controllers.Login)
// 	// app.Get("/api/user", controllers.User)
// 	// app.Post("/api/logout", controllers.Logout)

// }
