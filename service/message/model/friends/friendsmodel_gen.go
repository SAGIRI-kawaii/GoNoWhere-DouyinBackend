// Code generated by goctl. DO NOT EDIT.

package friends

import (
	"context"
	"database/sql"
	"fmt"
	"mini-douyin/service/message/model"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	friendsFieldNames          = builder.RawFieldNames(&Friends{})
	friendsRows                = strings.Join(friendsFieldNames, ",")
	friendsRowsExpectAutoSet   = strings.Join(stringx.Remove(friendsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	friendsRowsWithPlaceHolder = strings.Join(stringx.Remove(friendsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheFriendsIdPrefix = "cache:friends:id:"
)

type (
	friendsModel interface {
		Insert(ctx context.Context, data *Friends) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Friends, error)
		Update(ctx context.Context, data *Friends) error
		Delete(ctx context.Context, id int64) error
		DeleteById(ctx context.Context, id1 int64, id2 int64) error
		FindOneByBothway(ctx context.Context, id1 int64, id2 int64) (*Friends, error)
	}

	defaultFriendsModel struct {
		sqlc.CachedConn
		table string
	}

	Friends struct {
		Id       int64         `db:"id"`
		UserId   sql.NullInt64 `db:"user_id"`
		ToUserId sql.NullInt64 `db:"to_user_id"`
		CreateAt sql.NullTime  `db:"create_at"`
	}
)

func newFriendsModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFriendsModel {
	return &defaultFriendsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`friends`",
	}
}

func (m *defaultFriendsModel) Delete(ctx context.Context, id int64) error {
	friendsIdKey := fmt.Sprintf("%s%v", cacheFriendsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, friendsIdKey)
	return err
}
func (m *defaultFriendsModel) DeleteById(ctx context.Context, id1 int64, id2 int64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ? and `to_user_id` = ? ", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, id1, id2)
	return err

}

func (m *defaultFriendsModel) FindOne(ctx context.Context, id int64) (*Friends, error) {
	friendsIdKey := fmt.Sprintf("%s%v", cacheFriendsIdPrefix, id)
	var resp Friends
	err := m.QueryRowCtx(ctx, &resp, friendsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", friendsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultFriendsModel) FindOneByBothway(ctx context.Context, id1 int64, id2 int64) (*Friends, error) {
	var resp Friends
	query := fmt.Sprintf("select  %s from %s where `user_id` = ? and `to_user_id` = ? ", friendsRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, resp, query, id1, id2)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFriendsModel) Insert(ctx context.Context, data *Friends) (sql.Result, error) {
	friendsIdKey := fmt.Sprintf("%s%v", cacheFriendsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, friendsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ToUserId)
	}, friendsIdKey)
	return ret, err
}

func (m *defaultFriendsModel) Update(ctx context.Context, data *Friends) error {
	friendsIdKey := fmt.Sprintf("%s%v", cacheFriendsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, friendsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ToUserId, data.Id)
	}, friendsIdKey)
	return err
}

func (m *defaultFriendsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheFriendsIdPrefix, primary)
}

func (m *defaultFriendsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", friendsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFriendsModel) tableName() string {
	return m.table
}
