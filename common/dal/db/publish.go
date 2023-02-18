package db

import (
	"context"
	"mini-douyin/common/dal"

	"gorm.io/gorm"
)

// CreateVideo creates a new video
func CreateVideo(ctx context.Context, video *Video) error {
	err := dal.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(video).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// PublishList returns a list of videos with AuthorID.
func PublishList(ctx context.Context, authorId int64) ([]*Video, error) {
	var pubList []*Video
	err := dal.DB.WithContext(ctx).Model(&Video{}).Where(&Video{AuthorID: int(authorId)}).Find(&pubList).Error
	if err != nil {
		return nil, err
	}
	return pubList, nil
}
