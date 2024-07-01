package stock

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leetcode-golang-classroom/golang-stock-api/internal/types"
)

type Handler struct {
	stockStore types.StockStore
}

func NewHandler(stockStore types.StockStore) *Handler {
	return &Handler{stockStore: stockStore}
}

func (handler *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/api/stocks/{id}", handler.MiddlewareGetStock).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/stocks", handler.MiddlewareGetAllStock).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/stocks", handler.MiddlewareCreateStock).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/stocks/{id}", handler.MiddlewareUpdateStock).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/api/stocks/{id}", handler.MiddlewareDeleteStock).Methods(http.MethodDelete, http.MethodOptions)
}
