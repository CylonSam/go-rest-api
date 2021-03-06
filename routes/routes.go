package routes

import (
	"log"
	"net/http"

	"github.com/CylonSam/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
