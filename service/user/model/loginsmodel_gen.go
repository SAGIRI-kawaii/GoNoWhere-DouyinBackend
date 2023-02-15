// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	loginsFieldNames          = builder.RawFieldNames(&Logins{})
	loginsRows                = strings.Join(loginsFieldNames, ",")
	loginsRowsExpectAutoSet   = strings.Join(stringx.Remove(loginsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	loginsRowsWithPlaceHolder = strings.Join(stringx.Remove(loginsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLoginsIdPrefix   = "cache:logins:id:"
	cacheLoginsNamePrefix = "cache:logins:name:"
)

type (
	loginsModel interface {
		Insert(ctx context.Context, data *Logins) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Logins, error)
		FindOneByName(ctx context.Context, name string) (*Logins, error)
		Update(ctx context.Context, data *Logins) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLoginsModel struct {
		sqlc.CachedConn
		table string
	}

	Logins struct {
		Id        int64        `db:"id"` // 自增主键
		CreateAt  time.Time    `db:"create_at"`
		UpdateAt  time.Time    `db:"update_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
		Name      string       `db:"name"`
		PassWord  string       `db:"pass_word"`
	}
)

func newLoginsModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultLoginsModel {
	return &defaultLoginsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`logins`",
	}
}

func (m *defaultLoginsModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	loginsIdKey := fmt.Sprintf("%s%v", cacheLoginsIdPrefix, id)
	loginsNameKey := fmt.Sprintf("%s%v", cacheLoginsNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, loginsIdKey, loginsNameKey)
	return err
}

func (m *defaultLoginsModel) FindOne(ctx context.Context, id int64) (*Logins, error) {
	loginsIdKey := fmt.Sprintf("%s%v", cacheLoginsIdPrefix, id)
	var resp Logins
	err := m.QueryRowCtx(ctx, &resp, loginsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", loginsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLoginsModel) FindOneByName(ctx context.Context, name string) (*Logins, error) {
	loginsNameKey := fmt.Sprintf("%s%v", cacheLoginsNamePrefix, name)
	var resp Logins
	err := m.QueryRowIndexCtx(ctx, &resp, loginsNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", loginsRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLoginsModel) Insert(ctx context.Context, data *Logins) (sql.Result, error) {
	loginsIdKey := fmt.Sprintf("%s%v", cacheLoginsIdPrefix, data.Id)
	loginsNameKey := fmt.Sprintf("%s%v", cacheLoginsNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, loginsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.PassWord)
	}, loginsIdKey, loginsNameKey)
	return ret, err
}

func (m *defaultLoginsModel) Update(ctx context.Context, newData *Logins) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	loginsIdKey := fmt.Sprintf("%s%v", cacheLoginsIdPrefix, data.Id)
	loginsNameKey := fmt.Sprintf("%s%v", cacheLoginsNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, loginsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DeletedAt, newData.Name, newData.PassWord, newData.Id)
	}, loginsIdKey, loginsNameKey)
	return err
}

func (m *defaultLoginsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheLoginsIdPrefix, primary)
}

func (m *defaultLoginsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", loginsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLoginsModel) tableName() string {
	return m.table
}