package types

type CreateStockResponse struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type UpdateStockResponse struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
type DeleteStockResponse struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
type GetStockResponse struct {
	StockID int64  `json:"stock_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Price   int64  `json:"price" validate:"required"`
	Company string `json:"company" validate:"required"`
}

func ConvertStockToResponse(stock Stock) GetStockResponse {
	var response GetStockResponse
	response.StockID = stock.StockID
	response.Name = stock.Name
	response.Company = stock.Company
	response.Price = stock.Price
	return response
}

type GetAllStocksResponse struct {
	Stocks     []GetStockResponse `json:"stocks"`
	Pagination PaginationResult   `json:"pagination"`
}

func ConvertStocksToResponse(stocksResult StocksResult) GetAllStocksResponse {
	var response GetAllStocksResponse
	for _, stock := range stocksResult.Stocks {
		response.Stocks = append(response.Stocks, ConvertStockToResponse(stock))
	}
	response.Pagination = stocksResult.Pagination
	return response
}
