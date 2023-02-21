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
	commentsFieldNames          = builder.RawFieldNames(&Comments{})
	commentsRows                = strings.Join(commentsFieldNames, ",")
	commentsRowsExpectAutoSet   = strings.Join(stringx.Remove(commentsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	commentsRowsWithPlaceHolder = strings.Join(stringx.Remove(commentsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheCommentsIdPrefix = "cache:comments:id:"
)

type (
	commentsModel interface {
		Insert(ctx context.Context, data *Comments) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Comments, error)
		Update(ctx context.Context, data *Comments) error
		Delete(ctx context.Context, id int64) error
		FindList(ctx context.Context, id int64) ([]*Comments, error)
	}

	defaultCommentsModel struct {
		sqlc.CachedConn
		table string
	}

	Comments struct {
		Id        int64        `db:"id"` // 自增主键
		CreateAt  time.Time    `db:"create_at"`
		UpdateAt  time.Time    `db:"update_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
		UserId    int64        `db:"user_id"`
		VideoId   int64        `db:"video_id"` // 被评论视频ID
		Content   string       `db:"content"`  // 评论内容
	}
)

func newCommentsModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultCommentsModel {
	return &defaultCommentsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`comments`",
	}
}

func (m *defaultCommentsModel) Delete(ctx context.Context, id int64) error {
	commentsIdKey := fmt.Sprintf("%s%v", cacheCommentsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, commentsIdKey)
	return err
}

func (m *defaultCommentsModel) FindOne(ctx context.Context, id int64) (*Comments, error) {
	commentsIdKey := fmt.Sprintf("%s%v", cacheCommentsIdPrefix, id)
	var resp Comments
	err := m.QueryRowCtx(ctx, &resp, commentsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentsRows, m.table)
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

func (m *defaultCommentsModel) Insert(ctx context.Context, data *Comments) (sql.Result, error) {
	commentsIdKey := fmt.Sprintf("%s%v", cacheCommentsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, commentsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.VideoId, data.Content)
	}, commentsIdKey)
	return ret, err
}

func (m *defaultCommentsModel) Update(ctx context.Context, data *Comments) error {
	commentsIdKey := fmt.Sprintf("%s%v", cacheCommentsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, commentsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.VideoId, data.Content, data.Id)
	}, commentsIdKey)
	return err
}

func (m *defaultCommentsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheCommentsIdPrefix, primary)
}

func (m *defaultCommentsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCommentsModel) tableName() string {
	return m.table
}
func (m *defaultCommentsModel) FindList(ctx context.Context, id int64) ([]*Comments, error) {
	commentsIdKey := fmt.Sprintf("%s%v", cacheCommentsIdPrefix, id)
	// var resp Comments
	var c []*Comments
	err := m.QueryRowCtx(ctx, &c, commentsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `video_id` = ?", commentsRows, m.table)
		return conn.QueryRowsCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return c, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
