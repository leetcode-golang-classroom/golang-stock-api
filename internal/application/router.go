package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leetcode-golang-classroom/golang-stock-api/internal/stock"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	// set default health check
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(map[string]string{"message": "status ok"})
		if err != nil {
			log.Fatal(err)
		}
	})
	return router
}

func (app *App) loadRoutes() {
	stockStore := stock.NewStore(app.db)
	storeHandler := stock.NewHandler(stockStore)
	storeHandler.RegisterRoute(app.router)
}
