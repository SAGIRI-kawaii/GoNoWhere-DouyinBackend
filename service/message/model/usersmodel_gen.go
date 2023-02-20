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
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUsersIdPrefix     = "cache:users:id:"
	cacheUsersUserIdPrefix = "cache:users:userId:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		FindOneByUserId(ctx context.Context, userId int64) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id int64) error
		AddFollowerByUserId(ctx context.Context, uid int64) error
		AddFollowByUserId(ctx context.Context, uid int64) error
		ReduceFollowerByUserId(ctx context.Context, uid int64) error
		ReduceFollowByUserId(ctx context.Context, uid int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id              int64          `db:"id"`
		CreateAt        time.Time      `db:"create_at"`
		DeletedAt       sql.NullTime   `db:"deleted_at"`
		Name            string         `db:"name"`
		FollowCount     int64          `db:"follow_count"`
		FollowerCount   int64          `db:"follower_count"`
		UserId          int64          `db:"user_id"`
		Avatar          sql.NullString `db:"avatar"`
		BackgroundImage sql.NullString `db:"background_image"`
		Signature       sql.NullString `db:"signature"`
		TotalFavorited  sql.NullInt64  `db:"total_favorited"`
		WorkCount       sql.NullInt64  `db:"work_count"`
		FavoriteCount   sql.NullInt64  `db:"favorite_count"`
	}
)

func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usersIdKey, usersUserIdKey)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, usersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
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

func (m *defaultUsersModel) FindOneByUserId(ctx context.Context, userId int64) (*Users, error) {
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, userId)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, usersUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId); err != nil {
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

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.FollowCount, data.FollowerCount, data.UserId, data.Avatar, data.BackgroundImage, data.Signature, data.TotalFavorited, data.WorkCount, data.FavoriteCount)
	}, usersIdKey, usersUserIdKey)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, newData *Users) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DeletedAt, newData.Name, newData.FollowCount, newData.FollowerCount, newData.UserId, newData.Avatar, newData.BackgroundImage, newData.Signature, newData.TotalFavorited, newData.WorkCount, newData.FavoriteCount, newData.Id)
	}, usersIdKey, usersUserIdKey)
	return err
}
func (m *defaultUsersModel) AddFollowerByUserId(ctx context.Context, uid int64) error {
	data, err := m.FindOneByUserId(ctx, uid)
	if err != nil {
		return err
	}
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.FollowCount, data.FollowerCount+1, data.UserId, data.Id)
	}, usersIdKey, usersUserIdKey)
	return err
}
func (m *defaultUsersModel) AddFollowByUserId(ctx context.Context, uid int64) error {
	data, err := m.FindOneByUserId(ctx, uid)
	if err != nil {
		return err
	}
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.FollowCount+1, data.FollowerCount, data.UserId, data.Id)
	}, usersIdKey, usersUserIdKey)
	return err
}
func (m *defaultUsersModel) ReduceFollowByUserId(ctx context.Context, uid int64) error {
	data, err := m.FindOneByUserId(ctx, uid)
	if err != nil {
		return err
	}
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.FollowCount-1, data.FollowerCount, data.UserId, data.Id)
	}, usersIdKey, usersUserIdKey)
	return err
}
func (m *defaultUsersModel) ReduceFollowerByUserId(ctx context.Context, uid int64) error {
	data, err := m.FindOneByUserId(ctx, uid)
	if err != nil {
		return err
	}
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.FollowCount, data.FollowerCount-1, data.UserId, data.Id)
	}, usersIdKey, usersUserIdKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}