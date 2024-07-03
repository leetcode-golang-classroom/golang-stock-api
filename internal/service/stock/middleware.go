package stock

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/leetcode-golang-classroom/golang-stock-api/internal/types"
	"github.com/leetcode-golang-classroom/golang-stock-api/internal/util"
)

func (handler *Handler) MiddlewareGetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockID, ok := params["id"]
	if !ok {
		util.WriteError(w, http.StatusBadRequest, fmt.Errorf("stock id not provided"))
		return
	}
	id, err := strconv.ParseInt(stockID, 10, 64)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, fmt.Errorf("stock id is not in correct format %w", err))
		return
	}
	stock, err := handler.stockStore.GetStock(r.Context(), id)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to query stock %w", err))
		return
	}
	util.FailOnError(util.WriteJSON(w, http.StatusOK, types.ConvertStockToResponse(stock)), "failed to response json")
}

func (handler *Handler) MiddlewareGetAllStock(w http.ResponseWriter, r *http.Request) {
	pagination := types.QueryPagination{
		Offset: 0,
		Limit:  10,
	}
	query := r.URL.Query()
	if query.Has("limit") {
		limit, err := strconv.ParseInt(query.Get("limit"), 10, 64)
		if err != nil {
			util.WriteError(w, http.StatusBadRequest, fmt.Errorf("limit parse err: %w", err))
			return
		}
		pagination.Limit = limit
	}
	if query.Has("offset") {
		offset, err := strconv.ParseInt(query.Get("offset"), 10, 64)
		if err != nil {
			util.WriteError(w, http.StatusBadRequest, fmt.Errorf("offset parse err: %w", err))
			return
		}
		pagination.Offset = offset
	}
	result, err := handler.stockStore.GetAllStocks(r.Context(), pagination)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to query stock %w", err))
		return
	}
	util.FailOnError(util.WriteJSON(w, http.StatusOK, types.ConvertStocksToResponse(result)), "failed to response json")
}

func (handler *Handler) MiddlewareCreateStock(w http.ResponseWriter, r *http.Request) {
	var stock types.CreateStockRequest
	// decode data into stock
	if err := util.ParseJSON(r, &stock); err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// validate input
	if err := util.Validdate.Struct(stock); err != nil {
		var valErrs validator.ValidationErrors
		if errors.As(err, &valErrs) {
			util.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload:%v", valErrs))
		}
		return
	}
	stockID, err := handler.stockStore.CreateStock(r.Context(), stock)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create stock: %w", err))
		return
	}
	res := types.CreateStockResponse{
		ID:      stockID,
		Message: "stock created successfully",
	}
	util.FailOnError(util.WriteJSON(w, http.StatusCreated, res), "failed on response json")
}

func (handler *Handler) MiddlewareUpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockID, ok := params["id"]
	if !ok {
		util.WriteError(w, http.StatusBadRequest, fmt.Errorf("stock id not provided"))
		return
	}
	id, err := strconv.ParseInt(stockID, 10, 64)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, fmt.Errorf("stock id is not in correct format %w", err))
		return
	}
	var updateRequest types.UpdateStockRequest
	if err := util.ParseJSON(r, &updateRequest); err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// validate input
	if err := util.Validdate.Struct(updateRequest); err != nil {
		var valErrs validator.ValidationErrors
		if errors.As(err, &valErrs) {
			util.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload:%v", valErrs))
		}
		return
	}
	resultRows, err := handler.stockStore.UpdateStock(r.Context(), id, updateRequest)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to update stock %w", err))
		return
	}
	msg := fmt.Sprintf("Stock updated successfully. Total rows/records affows %v", resultRows)
	res := types.UpdateStockResponse{
		ID:      id,
		Message: msg,
	}
	util.FailOnError(util.WriteJSON(w, http.StatusCreated, res), "failed to response json")
}

func (handler *Handler) MiddlewareDeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockID, ok := params["id"]
	if !ok {
		util.WriteError(w, http.StatusBadRequest, fmt.Errorf("stock id not provided"))
		return
	}
	id, err := strconv.ParseInt(stockID, 10, 64)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, fmt.Errorf("stock id is not in correct format %w", err))
		return
	}
	deletedRows, err := handler.stockStore.DeleteStock(r.Context(), id)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to delete stock %w", err))
		return
	}
	msg := fmt.Sprintf("Stock deleted successfully. Total rows/records affows %v", deletedRows)
	res := types.DeleteStockResponse{
		ID:      id,
		Message: msg,
	}
	util.FailOnError(util.WriteJSON(w, http.StatusOK, res), "failed to response json")
}
