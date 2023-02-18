package db

import (
	"context"
	"mini-douyin/common/dal"

	"gorm.io/gorm"
	"mini-douyin/common/errno"
)

// Comment 属于 Video, VideoID是外键(belongs to)
// Comment 也属于 User,UserID是外键(belongs to)
type Comment struct {
	gorm.Model
	Video   Video  `gorm:"foreignkey:VideoID"`
	VideoID int    `gorm:"index:idx_videoid;not null"`
	User    User   `gorm:"foreignkey:UserID"`
	UserID  int    `gorm:"index:idx_userid;not null"`
	Content string `gorm:"type:varchar(255);not null"`
}

func (Comment) TableName() string {
	return "comment"
}

// NewComment creates a new Comment
func NewComment(ctx context.Context, comment *Comment) error {
	err := dal.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		// 1. 新增评论数据
		err := tx.Create(comment).Error
		if err != nil {
			return err
		}

		// 2.改变 video 表中的 comment count
		res := tx.Model(&Video{}).Where("ID = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDataBase
		}

		return nil
	})
	return err
}

// DelComment deletes a comment from the database.
func DelComment(ctx context.Context, commentID int64, vid int64) error {
	err := dal.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		comment := new(Comment)
		if err := tx.First(&comment, commentID).Error; err != nil {
			return err
		}

		// 1. 删除评论数据
		// 因为 Comment中包含了gorm.Model所以拥有软删除能力
		// 而tx.Unscoped().Delete()将永久删除记录
		err := tx.Unscoped().Delete(&comment).Error
		if err != nil {
			return err
		}

		// 2.改变 video 表中的 comment count
		res := tx.Model(&Video{}).Where("ID = ?", vid).Update("comment_count", gorm.Expr("comment_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDataBase
		}

		return nil
	})
	return err
}

// GetVideoComments returns a list of video comments.
func GetVideoComments(ctx context.Context, vid int64) ([]*Comment, error) {
	var comments []*Comment
	err := dal.DB.WithContext(ctx).Model(&Comment{}).Where(&Comment{VideoID: int(vid)}).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
