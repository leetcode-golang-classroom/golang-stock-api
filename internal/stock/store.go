package stock

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/leetcode-golang-classroom/golang-stock-api/internal/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (store *Store) CreateStock(ctx context.Context, stock types.CreateStockRequest) (int64, error) {
	queryBuilder := sq.Insert("stocks").Columns("name", "price", "company")
	queryBuilder = queryBuilder.Values(stock.Name, stock.Price, stock.Company)
	queryBuilder = queryBuilder.Suffix("RETURNING stockid;").PlaceholderFormat(sq.Dollar)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return -1, fmt.Errorf("failed to create prepare statement %w", err)
	}
	rows, err := store.db.QueryContext(ctx, query, args...)
	if err != nil {
		return -1, fmt.Errorf("failed to insert stock data %w", err)
	}
	var id int64
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return -1, err
		}
	}
	return id, nil
}

func (store *Store) GetStock(ctx context.Context, stockID int64) (types.Stock, error) {
	queryBuilder := sq.Select("stockid", "name", "price", "company")
	queryBuilder = queryBuilder.From("stocks").Where(sq.Eq{"stockid": stockID}).PlaceholderFormat(sq.Dollar)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return types.Stock{}, fmt.Errorf("failed to create prepare statement %w", err)
	}
	rows, err := store.db.QueryContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.Stock{}, fmt.Errorf("no stock record found with stockid=%d %w", stockID, err)
		}
		return types.Stock{}, fmt.Errorf("failed to query stock with stockid=%d %w", stockID, err)
	}
	var stock types.Stock
	for rows.Next() {
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			return types.Stock{}, err
		}
	}
	if stock.StockID == 0 {
		return types.Stock{}, fmt.Errorf("no stock record found with stockid=%d", stockID)
	}
	return stock, nil
}

func (store *Store) GetAllStocks(ctx context.Context, pagination types.QueryPagination) (types.StocksResult, error) {
	queryBuilder := sq.Select("stockid", "name", "price", "company").From("stocks").OrderBy("stockid ASC")
	// whereCondition := []sq.Sqlizer{}

	// queryBuilder = queryBuilder.Where(sq.And(whereCondition))
	var limit uint64 = 10
	if pagination.Limit > 0 {
		limit = uint64(pagination.Limit)
	}
	var offset uint64 = 0
	if pagination.Offset > 0 {
		offset = uint64(pagination.Offset)
	}
	queryBuilder = queryBuilder.Offset(offset)
	queryBuilder = queryBuilder.Limit(limit)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return types.StocksResult{}, fmt.Errorf("failed to create preparement statement %w", err)
	}
	rows, err := store.db.QueryContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.StocksResult{}, fmt.Errorf("no stock record found %w", err)
		}
		return types.StocksResult{}, fmt.Errorf("failed to query rows %w", err)
	}
	var result types.StocksResult
	for rows.Next() {
		var stock types.Stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			return types.StocksResult{}, err
		}
		result.Stocks = append(result.Stocks, stock)
	}
	result.Pagination.Limit = pagination.Limit
	result.Pagination.Offset = pagination.Offset
	if len(result.Stocks) > 0 {
		result.Pagination.NextOffset = result.Pagination.Offset + result.Pagination.Limit
	}

	return result, nil
}

func (store *Store) UpdateStock(ctx context.Context, stockID int64, updateInfo types.UpdateStockRequest) (int64, error) {
	queryBuilder := sq.Update("stocks").Where(sq.Eq{"stockid": stockID})
	queryBuilder = queryBuilder.Set("price", updateInfo.Price).PlaceholderFormat(sq.Dollar)
	if len(updateInfo.Name) > 0 {
		queryBuilder = queryBuilder.Set("name", updateInfo.Name)
	}
	if len(updateInfo.Company) > 0 {
		queryBuilder = queryBuilder.Set("company", updateInfo.Company)
	}
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return -1, fmt.Errorf("failed to create preparestmt %w", err)
	}
	result, err := store.db.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, fmt.Errorf("failed to update stock %w", err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("failed to get affected rows %w", err)
	}
	return affectedRows, nil
}

func (store *Store) DeleteStock(ctx context.Context, stockID int64) (int64, error) {
	queryBuilder := sq.Delete("stocks").Where(sq.Eq{"stockid": stockID}).PlaceholderFormat(sq.Dollar)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return -1, fmt.Errorf("failed to create preparestmt %w", err)
	}
	result, err := store.db.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, fmt.Errorf("failed to delete stock %w", err)
	}
	deletedRows, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("failed to get affected rows %w", err)
	}
	return deletedRows, nil
}
