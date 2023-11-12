package ratelimit

import (
	"context"
	"fmt"
	"webook/internal/service/sms"
	"webook/internal/service/sms/tencent"
	"webook/pkg/ratelimit"
)

var errLimited = fmt.Errorf("ratelimited")

type RateLimitSMSService struct {
	svc     sms.Service
	limiter ratelimit.Limiter
}

func (s RateLimitSMSService) Send(ctx context.Context, tpl string, args []sms.ArgVal, phones ...string) error {
	limited, err := s.limiter.Limit(ctx, tencent.LimitKey)
	if err != nil {
		return fmt.Errorf("sms ratelimiter limit err: %w", err)
	}
	if limited {
		return errLimited
	}
	err = s.svc.Send(ctx, tpl, args, phones...)

	return err
}

func NewRateLimitSMSService(svc sms.Service, limiter ratelimit.Limiter) sms.Service {
	return &RateLimitSMSService{
		svc:     svc,
		limiter: limiter,
	}
}
