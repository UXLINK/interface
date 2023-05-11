package {{.pkg}}
import (
	"context"
    "database/sql"
    "fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ {{.upperStartCamelObject}}Model = (*custom{{.upperStartCamelObject}}Model)(nil)

type (
	// {{.upperStartCamelObject}}Model is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.upperStartCamelObject}}Model.
	{{.upperStartCamelObject}}Model interface {
		{{.lowerStartCamelObject}}Model
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
	    InsertWithSession(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result, error)
	    UpdateWithSession(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) error
	}

	custom{{.upperStartCamelObject}}Model struct {
		*default{{.upperStartCamelObject}}Model
	}
)

// New{{.upperStartCamelObject}}Model returns a model for the database table.
func New{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf{{end}}) {{.upperStartCamelObject}}Model {
	return &custom{{.upperStartCamelObject}}Model{
		default{{.upperStartCamelObject}}Model: new{{.upperStartCamelObject}}Model(conn{{if .withCache}}, c{{end}}),
	}
}

func (m *custom{{.upperStartCamelObject}}Model) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *custom{{.upperStartCamelObject}}Model) InsertWithSession(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
	return session.ExecCtx(ctx, query)
}

func (m *default{{.upperStartCamelObject}}Model) UpdateWithSession(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) error {
	query := fmt.Sprintf("update %s set %s where `XXXX` = ?", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	_, err := session.ExecCtx(ctx, query)
	return err
}