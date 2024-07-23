package service

import (
	"context"
	"github.com/lazywoo/mercury/feed/domain"
)

type FeedService interface {
	CreateFeedEvent(ctx context.Context, feed domain.FeedEvent) error
	GetFeedEventList(ctx context.Context, uid, timestamp, limit int64) ([]domain.FeedEvent, error)
}

type EventHandler interface {
	CreateFeedEvent(ctx context.Context, feed domain.FeedEvent) error
	FindFeedEvents(ctx context.Context, uid, timestamp, limit int64) ([]domain.FeedEvent, error)
}
