package domain

import "time"

type FeedEventType uint8

const (
	FeedEventTypeUnknown = iota
	FeedEventTypePushEvent
	FeedEventTypePullEvent
)

type FeedEvent struct {
	ID    int64
	Uid   int64
	Type  FeedEventType
	Ctime time.Time
	Ext   ExtendFields
}
