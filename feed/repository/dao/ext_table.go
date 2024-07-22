package dao

import (
	"context"
	"gorm.io/gorm"
)

// ext_table usecase

type FeedEvent struct {
	Id   int64
	Type string
	Ext  any `gorm:"-"`
}

type ArticleEvent struct {
	Id        int64
	Fid       int64
	ArticleId int64
}

type FeedEventDAO struct {
	db *gorm.DB
}

func (dao *FeedEventDAO) Find(ctx context.Context) ([]FeedEvent, error) {
	var feedEvents []FeedEvent
	err := dao.db.WithContext(ctx).Where("").Find(&feedEvents).Error
	if err != nil {
		return nil, err
	}

	feedIds := make([]int64, 0, len(feedEvents))
	feedEventMap := make(map[int64]*FeedEvent)
	for i := range feedEvents {
		feedIds = append(feedIds, feedEvents[i].Id)
		feedEventMap[feedEvents[i].Id] = &feedEvents[i]
	}

	var atclFeedEvents []ArticleEvent
	err = dao.db.Where("fid in ?", feedIds).Find(&atclFeedEvents).Error
	if err != nil {
		return nil, err
	}

	// 组装
	for i := range atclFeedEvents {
		feedEventMap[atclFeedEvents[i].Fid].Ext = atclFeedEvents[i]
	}

	return feedEvents, nil
}
