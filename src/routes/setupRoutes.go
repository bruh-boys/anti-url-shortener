package routes

import (
	"net/http"
	"os"

	"github.com/bruh-boys/anti-url-shortener/src/controllers"
)

func SetupRoutes() error {

	http.Handle("/", http.FileServer(http.Dir("./src/static")))
	http.HandleFunc("/no-bitly", controllers.NoBitly)
	port, err := os.LookupEnv("PORT")
	if !err {
		port = "8080"
	}

	return http.ListenAndServe(":"+port, nil)

}
