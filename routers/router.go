package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {

	router := mux.NewRouter()

	http.ListenAndServe(":9000", router)

}
