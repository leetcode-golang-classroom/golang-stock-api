package types

import "context"

type StockStore interface {
	CreateStock(ctx context.Context, stock CreateStockRequest) (int64, error)
	GetStock(ctx context.Context, stockID int64) (Stock, error)
	GetAllStocks(ctx context.Context, pagination QueryPagination) (StocksResult, error)
	UpdateStock(ctx context.Context, stockID int64, updateInfo UpdateStockRequest) (int64, error)
	DeleteStock(ctx context.Context, stockID int64) (int64, error)
}
