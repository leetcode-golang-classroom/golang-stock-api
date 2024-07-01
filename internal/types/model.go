package types

type Stock struct {
	StockID int64  `json:"stock_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Price   int64  `json:"price" validate:"required"`
	Company string `json:"company" validate:"required"`
}

type StocksResult struct {
	Stocks     []Stock          `json:"stocks"`
	Pagination PaginationResult `json:"pagination"`
}
