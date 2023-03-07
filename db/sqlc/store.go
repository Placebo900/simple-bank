package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all funcitons to execute db entries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db, Queries: New(db)}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	if err = fn(q); err != nil {
		if rbErr := tx.Rollback(); err != nil {
			return fmt.Errorf("txError: %v, rbError: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var res TransferTxResult
	var err error
	err = store.execTx(ctx, func(q *Queries) error {
		if res.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams(arg)); err != nil {
			return err
		}

		if res.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{AccountID: arg.FromAccountID, Amount: -arg.Amount}); err != nil {
			return err
		}

		if res.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{AccountID: arg.ToAccountID, Amount: arg.Amount}); err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {
			res.FromAccount, res.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, arg.ToAccountID, arg.Amount)
		} else {
			res.ToAccount, res.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.FromAccountID, -arg.Amount)
		}

		return err
	})
	return res, err
}

func addMoney(ctx context.Context, q *Queries, from, to, amount int64) (a, b Account, err error) {
	if a, err = q.AddAccountMoney(ctx, AddAccountMoneyParams{ID: from, Amount: -amount}); err != nil {
		return
	}

	if b, err = q.AddAccountMoney(ctx, AddAccountMoneyParams{ID: to, Amount: amount}); err != nil {
		return
	}
	return
}
