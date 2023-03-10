// Code generated by goctl. DO NOT EDIT.

package favorites

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
	favoritesFieldNames          = builder.RawFieldNames(&Favorites{})
	favoritesRows                = strings.Join(favoritesFieldNames, ",")
	favoritesRowsExpectAutoSet   = strings.Join(stringx.Remove(favoritesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	favoritesRowsWithPlaceHolder = strings.Join(stringx.Remove(favoritesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheFavoritesIdPrefix = "cache:favorites:id:"
)

type (
	favoritesModel interface {
		Insert(ctx context.Context, data *Favorites) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Favorites, error)
		Update(ctx context.Context, data *Favorites) error
		Delete(ctx context.Context, id int64) error
		DeleteByUserId2VideoId(ctx context.Context, UserId int64, VideoId int64) error
		FindVideoListByUserId(ctx context.Context, UserId int64) ([]int64, error)
		// FindAuthor(ctx context.Context, VideoId int64) (*Users, error)
	}

	defaultFavoritesModel struct {
		sqlc.CachedConn
		table string
	}

	Favorites struct {
		Id       int64     `db:"id"`        // 自增主键
		CreateAt time.Time `db:"create_at"` // 创建时间
		UserId   int64     `db:"user_id"`   // 点赞用户ID
		VideoId  int64     `db:"video_id"`  // 被点赞视频ID
	}
)

func newFavoritesModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFavoritesModel {
	return &defaultFavoritesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`favorites`",
	}
}

func (m *defaultFavoritesModel) Delete(ctx context.Context, id int64) error {
	favoritesIdKey := fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, favoritesIdKey)
	return err
}

func (m *defaultFavoritesModel) FindOne(ctx context.Context, id int64) (*Favorites, error) {
	favoritesIdKey := fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, id)
	var resp Favorites
	err := m.QueryRowCtx(ctx, &resp, favoritesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", favoritesRows, m.table)
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

func (m *defaultFavoritesModel) Insert(ctx context.Context, data *Favorites) (sql.Result, error) {
	favoritesIdKey := fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, favoritesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.VideoId)
	}, favoritesIdKey)
	return ret, err
}

func (m *defaultFavoritesModel) Update(ctx context.Context, data *Favorites) error {
	favoritesIdKey := fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, favoritesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.VideoId, data.Id)
	}, favoritesIdKey)
	return err
}

func (m *defaultFavoritesModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, primary)
}

func (m *defaultFavoritesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", favoritesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFavoritesModel) tableName() string {
	return m.table
}

func (m *defaultFavoritesModel) DeleteByUserId2VideoId(ctx context.Context, UserId int64, VideoId int64) error {
	favoritesIdKey := fmt.Sprintf("%s%v%v", cacheFavoritesIdPrefix, UserId, VideoId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ? and `video_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, UserId, VideoId)
	}, favoritesIdKey)
	return err
}

func (m *defaultFavoritesModel) FindVideoListByUserId(ctx context.Context, UserId int64) ([]int64, error) {
	// favoritesIdKey := fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, UserId)
	// var resp  Favorites
	var c []int64
	// err := m.QueryRowCtx(ctx, &c, favoritesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select  DISTINCT `video_id` from %s  where `user_id` = ?", m.table)
		// return conn.QueryRowsCtx(ctx, v, query, UserId)
		err := m.QueryRowsNoCacheCtx(ctx, &c, query, UserId)
	// })
	switch err {
	case nil:
		return c, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// func (m *defaultFavoritesModel) FindAuthor(ctx context.Context, VideoId int64) (*Users, error) {
// 	favoritesIdKey := fmt.Sprintf("%s%v", cacheFavoritesIdPrefix, VideoId)
// 	// var resp  Favorites
// 	var c *Users
// 	err := m.QueryRowCtx(ctx, c, favoritesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {

// 		query1 := fmt.Sprintf("select users.name from videos join users on users.user_id == videos.user_id  where video_id =?", favoritesRows)
// 		println(query1)
// 		return conn.QueryRowsCtx(ctx, v, query1, VideoId)

// 	})
// 	switch err {
// 	case nil:
// 		return c, nil
// 	case sqlc.ErrNotFound:
// 		return nil, ErrNotFound
// 	default:
// 		return nil, err
// 	}
// }
