package internal

import (
	"database/sql"
)

type GooseDB interface {
	//PingContext(ctx context.Context) error
	//Ping() error
	//Close() error
	//SetMaxIdleConns(n int)
	//SetMaxOpenConns(n int)
	//SetConnMaxLifetime(d time.Duration)
	//SetConnMaxIdleTime(d time.Duration)
	//Stats() sql.DBStats
	//PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	//Prepare(query string) (*sql.Stmt, error)
	//ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	//QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Query(query string, args ...interface{}) (GooseRows, error)
	//QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryRow(query string, args ...interface{}) *sql.Row
	//BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Begin() (*sql.Tx, error)
	//Driver() driver.Driver
	//Conn(ctx context.Context) (*sql.Conn, error)
}

/*
	type GooseTx interface {
		Commit() error
		Rollback() error
		//PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
		//Prepare(query string) (*sql.Stmt, error)
		//StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
		//Stmt(stmt *sql.Stmt) *sql.Stmt
		//ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
		Exec(query string, args ...interface{}) (sql.Result, error)
		//QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
		//Query(query string, args ...any) (*sql.Rows, error)
		//QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
		//QueryRow(query string, args ...any) *sql.Row
	}
*/
type GooseRows interface {
	Next() bool
	//NextResultSet() bool
	Err() error
	//Columns() ([]string, error)
	//ColumnTypes() ([]*sql.ColumnType, error)
	Scan(dest ...interface{}) error
	Close() error
}
