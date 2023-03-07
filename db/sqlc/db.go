// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addAccountMoneyStmt, err = db.PrepareContext(ctx, addAccountMoney); err != nil {
		return nil, fmt.Errorf("error preparing query AddAccountMoney: %w", err)
	}
	if q.createAccountStmt, err = db.PrepareContext(ctx, createAccount); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccount: %w", err)
	}
	if q.createEntryStmt, err = db.PrepareContext(ctx, createEntry); err != nil {
		return nil, fmt.Errorf("error preparing query CreateEntry: %w", err)
	}
	if q.createTransferStmt, err = db.PrepareContext(ctx, createTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTransfer: %w", err)
	}
	if q.deleteAccountStmt, err = db.PrepareContext(ctx, deleteAccount); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAccount: %w", err)
	}
	if q.deleteEntryStmt, err = db.PrepareContext(ctx, deleteEntry); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteEntry: %w", err)
	}
	if q.deleteTransferStmt, err = db.PrepareContext(ctx, deleteTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTransfer: %w", err)
	}
	if q.getAccountStmt, err = db.PrepareContext(ctx, getAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccount: %w", err)
	}
	if q.getAccountForUpdateStmt, err = db.PrepareContext(ctx, getAccountForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccountForUpdate: %w", err)
	}
	if q.getEntryStmt, err = db.PrepareContext(ctx, getEntry); err != nil {
		return nil, fmt.Errorf("error preparing query GetEntry: %w", err)
	}
	if q.getEntryForUpdateStmt, err = db.PrepareContext(ctx, getEntryForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetEntryForUpdate: %w", err)
	}
	if q.getTransferStmt, err = db.PrepareContext(ctx, getTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query GetTransfer: %w", err)
	}
	if q.getTransferForUpdateStmt, err = db.PrepareContext(ctx, getTransferForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetTransferForUpdate: %w", err)
	}
	if q.listAccountStmt, err = db.PrepareContext(ctx, listAccount); err != nil {
		return nil, fmt.Errorf("error preparing query ListAccount: %w", err)
	}
	if q.listEntryStmt, err = db.PrepareContext(ctx, listEntry); err != nil {
		return nil, fmt.Errorf("error preparing query ListEntry: %w", err)
	}
	if q.listTransferStmt, err = db.PrepareContext(ctx, listTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query ListTransfer: %w", err)
	}
	if q.updateAccountStmt, err = db.PrepareContext(ctx, updateAccount); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAccount: %w", err)
	}
	if q.updateEntryStmt, err = db.PrepareContext(ctx, updateEntry); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateEntry: %w", err)
	}
	if q.updateTransferStmt, err = db.PrepareContext(ctx, updateTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTransfer: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addAccountMoneyStmt != nil {
		if cerr := q.addAccountMoneyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addAccountMoneyStmt: %w", cerr)
		}
	}
	if q.createAccountStmt != nil {
		if cerr := q.createAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccountStmt: %w", cerr)
		}
	}
	if q.createEntryStmt != nil {
		if cerr := q.createEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createEntryStmt: %w", cerr)
		}
	}
	if q.createTransferStmt != nil {
		if cerr := q.createTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTransferStmt: %w", cerr)
		}
	}
	if q.deleteAccountStmt != nil {
		if cerr := q.deleteAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAccountStmt: %w", cerr)
		}
	}
	if q.deleteEntryStmt != nil {
		if cerr := q.deleteEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteEntryStmt: %w", cerr)
		}
	}
	if q.deleteTransferStmt != nil {
		if cerr := q.deleteTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTransferStmt: %w", cerr)
		}
	}
	if q.getAccountStmt != nil {
		if cerr := q.getAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountStmt: %w", cerr)
		}
	}
	if q.getAccountForUpdateStmt != nil {
		if cerr := q.getAccountForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountForUpdateStmt: %w", cerr)
		}
	}
	if q.getEntryStmt != nil {
		if cerr := q.getEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEntryStmt: %w", cerr)
		}
	}
	if q.getEntryForUpdateStmt != nil {
		if cerr := q.getEntryForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEntryForUpdateStmt: %w", cerr)
		}
	}
	if q.getTransferStmt != nil {
		if cerr := q.getTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTransferStmt: %w", cerr)
		}
	}
	if q.getTransferForUpdateStmt != nil {
		if cerr := q.getTransferForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTransferForUpdateStmt: %w", cerr)
		}
	}
	if q.listAccountStmt != nil {
		if cerr := q.listAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAccountStmt: %w", cerr)
		}
	}
	if q.listEntryStmt != nil {
		if cerr := q.listEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listEntryStmt: %w", cerr)
		}
	}
	if q.listTransferStmt != nil {
		if cerr := q.listTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTransferStmt: %w", cerr)
		}
	}
	if q.updateAccountStmt != nil {
		if cerr := q.updateAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAccountStmt: %w", cerr)
		}
	}
	if q.updateEntryStmt != nil {
		if cerr := q.updateEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateEntryStmt: %w", cerr)
		}
	}
	if q.updateTransferStmt != nil {
		if cerr := q.updateTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTransferStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                       DBTX
	tx                       *sql.Tx
	addAccountMoneyStmt      *sql.Stmt
	createAccountStmt        *sql.Stmt
	createEntryStmt          *sql.Stmt
	createTransferStmt       *sql.Stmt
	deleteAccountStmt        *sql.Stmt
	deleteEntryStmt          *sql.Stmt
	deleteTransferStmt       *sql.Stmt
	getAccountStmt           *sql.Stmt
	getAccountForUpdateStmt  *sql.Stmt
	getEntryStmt             *sql.Stmt
	getEntryForUpdateStmt    *sql.Stmt
	getTransferStmt          *sql.Stmt
	getTransferForUpdateStmt *sql.Stmt
	listAccountStmt          *sql.Stmt
	listEntryStmt            *sql.Stmt
	listTransferStmt         *sql.Stmt
	updateAccountStmt        *sql.Stmt
	updateEntryStmt          *sql.Stmt
	updateTransferStmt       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                       tx,
		tx:                       tx,
		addAccountMoneyStmt:      q.addAccountMoneyStmt,
		createAccountStmt:        q.createAccountStmt,
		createEntryStmt:          q.createEntryStmt,
		createTransferStmt:       q.createTransferStmt,
		deleteAccountStmt:        q.deleteAccountStmt,
		deleteEntryStmt:          q.deleteEntryStmt,
		deleteTransferStmt:       q.deleteTransferStmt,
		getAccountStmt:           q.getAccountStmt,
		getAccountForUpdateStmt:  q.getAccountForUpdateStmt,
		getEntryStmt:             q.getEntryStmt,
		getEntryForUpdateStmt:    q.getEntryForUpdateStmt,
		getTransferStmt:          q.getTransferStmt,
		getTransferForUpdateStmt: q.getTransferForUpdateStmt,
		listAccountStmt:          q.listAccountStmt,
		listEntryStmt:            q.listEntryStmt,
		listTransferStmt:         q.listTransferStmt,
		updateAccountStmt:        q.updateAccountStmt,
		updateEntryStmt:          q.updateEntryStmt,
		updateTransferStmt:       q.updateTransferStmt,
	}
}