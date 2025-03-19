package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	dbPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new Store
func NewStore(dbPool *pgxpool.Pool) *Store {
	return &Store{
		dbPool:  dbPool,
		Queries: New(dbPool),
	}
}

// execTx executes a function within a database transaction
// func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.dbPool.BeginTx(context.Background(), nil)
// }
