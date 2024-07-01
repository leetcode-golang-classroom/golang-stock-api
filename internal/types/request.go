package types

type CreateStockRequest struct {
	Name    string `json:"name" validate:"required"`
	Price   int64  `json:"price" validate:"required"`
	Company string `json:"company" validate:"required"`
}

type QueryPagination struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit,omitempty"`
}

type PaginationResult struct {
	QueryPagination
	NextOffset int64 `json:"next_offset,omitempty"`
}

type UpdateStockRequest struct {
	Name    string `json:"name,omitempty"`
	Price   int64  `json:"price" validate:"required"`
	Company string `json:"company,omitempty"`
}
