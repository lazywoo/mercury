package service

import (
	"context"
	"github.com/lazywoo/mercury/feed/domain"
)

var _ EventHandler = (*ArticleEventHandler)(nil)

type ArticleEventHandler struct {
}

func (a *ArticleEventHandler) CreateFeedEvent(ctx context.Context, feed domain.FeedEvent) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleEventHandler) FindFeedEvents(ctx context.Context, uid, timestamp, limit int64) ([]domain.FeedEvent, error) {
	//TODO implement me
	panic("implement me")
}
