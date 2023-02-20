// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	messagesFieldNames          = builder.RawFieldNames(&Messages{})
	messagesRows                = strings.Join(messagesFieldNames, ",")
	messagesRowsExpectAutoSet   = strings.Join(stringx.Remove(messagesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	messagesRowsWithPlaceHolder = strings.Join(stringx.Remove(messagesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheMessagesIdPrefix = "cache:messages:id:"
)

type (
	messagesModel interface {
		Insert(ctx context.Context, data *Messages) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Messages, error)
		Update(ctx context.Context, data *Messages) error
		Delete(ctx context.Context, id int64) error
		FindLatestMsg(ctx context.Context, uid int64, preMsgTime string) (*[]Messages, error)
	}

	defaultMessagesModel struct {
		sqlc.CachedConn
		table string
	}

	Messages struct {
		Id       int64          `db:"id"`
		UserId   sql.NullInt64  `db:"user_id"`
		ToUserId sql.NullInt64  `db:"to_user_id"`
		Content  sql.NullString `db:"content"`
		CreateAt sql.NullTime   `db:"create_at"`
		DeleteAt sql.NullTime   `db:"delete_at"`
		UpdateAt sql.NullTime   `db:"update_at"`
	}
)

func newMessagesModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultMessagesModel {
	return &defaultMessagesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`messages`",
	}
}

func (m *defaultMessagesModel) Delete(ctx context.Context, id int64) error {
	messagesIdKey := fmt.Sprintf("%s%v", cacheMessagesIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, messagesIdKey)
	return err
}

func (m *defaultMessagesModel) FindOne(ctx context.Context, id int64) (*Messages, error) {
	messagesIdKey := fmt.Sprintf("%s%v", cacheMessagesIdPrefix, id)
	var resp Messages
	err := m.QueryRowCtx(ctx, &resp, messagesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", messagesRows, m.table)
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

func (m *defaultMessagesModel) FindLatestMsg(ctx context.Context, uid int64, preMsgTime string) (*[]Messages, error) {
	var resp []Messages
	query := fmt.Sprintf("select %s from %s where (`user_id` = ? or `to_user_id` = ?) and create_at >= ? order by create_at asc limit 30", messagesRows, m.table)
	println(query)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, uid, uid, preMsgTime)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		println("error1")
		return nil, ErrNotFound
	default:
		println(err.Error())
		return nil, err
	}
}

func (m *defaultMessagesModel) Insert(ctx context.Context, data *Messages) (sql.Result, error) {
	messagesIdKey := fmt.Sprintf("%s%v", cacheMessagesIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, messagesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ToUserId, data.Content, data.DeleteAt)
	}, messagesIdKey)
	return ret, err
}

func (m *defaultMessagesModel) Update(ctx context.Context, data *Messages) error {
	messagesIdKey := fmt.Sprintf("%s%v", cacheMessagesIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, messagesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ToUserId, data.Content, data.DeleteAt, data.Id)
	}, messagesIdKey)
	return err
}

func (m *defaultMessagesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMessagesIdPrefix, primary)
}

func (m *defaultMessagesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", messagesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultMessagesModel) tableName() string {
	return m.table
}