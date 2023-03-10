// Code generated by goctl. DO NOT EDIT.

package follows

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
	followsFieldNames          = builder.RawFieldNames(&Follows{})
	followsRows                = strings.Join(followsFieldNames, ",")
	followsRowsExpectAutoSet   = strings.Join(stringx.Remove(followsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	followsRowsWithPlaceHolder = strings.Join(stringx.Remove(followsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheFollowsIdPrefix = "cache:follows:id:"
)

type (
	followsModel interface {
		Insert(ctx context.Context, data *Follows) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Follows, error)
		FindAllByUserId(ctx context.Context, uid int64) ([]*Follows, error)     // 查询关注列表
		FindAllByToUserId(ctx context.Context, touid int64) ([]*Follows, error) // 查询被关注列表
		FindOneById(ctx context.Context, uid int64, touid int64) error          //查询是否存在
		Update(ctx context.Context, data *Follows) error
		Delete(ctx context.Context, id int64) error
		DeleteByuid(ctx context.Context, uid int64, touid int64) error // 取消关注

	}

	defaultFollowsModel struct {
		sqlc.CachedConn
		table string
	}

	Follows struct {
		Id        int64        `db:"id"`
		CreatedAt sql.NullTime `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
		UserId    int64        `db:"user_id"`    // 粉丝用户ID
		ToUserId  int64        `db:"to_user_id"` // 被关注用户ID
	}
)

func newFollowsModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFollowsModel {
	return &defaultFollowsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`follows`",
	}
}

func (m *defaultFollowsModel) Delete(ctx context.Context, id int64) error {
	followsIdKey := fmt.Sprintf("%s%v", cacheFollowsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, followsIdKey)
	return err
}
func (m *defaultFollowsModel) DeleteByuid(ctx context.Context, uid int64, touid int64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ? and `to_user_id` = ? ", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, uid, touid)
	return err
}

func (m *defaultFollowsModel) FindOne(ctx context.Context, id int64) (*Follows, error) {
	followsIdKey := fmt.Sprintf("%s%v", cacheFollowsIdPrefix, id)
	var resp Follows
	err := m.QueryRowCtx(ctx, &resp, followsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", followsRows, m.table)
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
func (m defaultFollowsModel) FindOneById(ctx context.Context, uid int64, touid int64) error {
	var resp Follows
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `to_user_id` = ?", followsRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, resp, query, uid, touid)
	switch err {
	case nil:
		return nil
	case sqlc.ErrNotFound:
		return ErrNotFound
	default:
		return err
	}
}
func (m *defaultFollowsModel) FindAllByUserId(ctx context.Context, uid int64) ([]*Follows, error) {
	var resp []*Follows
	query := fmt.Sprintf("select %s from %s where `user_id` = ? ", followsRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultFollowsModel) FindAllByToUserId(ctx context.Context, touid int64) ([]*Follows, error) {
	var resp []*Follows
	query := fmt.Sprintf("select %s from %s where `to_user_id` = ? ", followsRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, touid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFollowsModel) Insert(ctx context.Context, data *Follows) (sql.Result, error) {
	followsIdKey := fmt.Sprintf("%s%v", cacheFollowsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, followsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.ToUserId)
	}, followsIdKey)
	return ret, err
}

func (m *defaultFollowsModel) Update(ctx context.Context, data *Follows) error {
	followsIdKey := fmt.Sprintf("%s%v", cacheFollowsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, followsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.ToUserId, data.Id)
	}, followsIdKey)
	return err
}

func (m *defaultFollowsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheFollowsIdPrefix, primary)
}

func (m *defaultFollowsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", followsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFollowsModel) tableName() string {
	return m.table
}
