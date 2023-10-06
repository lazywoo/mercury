package cache

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	ErrSetCaptchaTooManyTimes    = errors.New("send captcha too frequently")
	ErrInternal                  = errors.New("internal error")
	ErrCaptchaVerifyTooManyTimes = errors.New("verify captcha too frequently")
	ErrUnknownForCode            = errors.New("unknown error for code")
)

//go:embed lua/set_code.lua
var luaSetCaptcha string

//go:embed lua/verify_code.lua
var luaVerifyCode string

var _ CaptchaCache = (*RedisCaptchaCache)(nil)

type CaptchaCache interface {
	Set(ctx context.Context, biz string, phone string, code string) error
	Verify(ctx context.Context, biz string, phone string, inputCaptcha string) (bool, error)
}

type RedisCaptchaCache struct {
	client redis.Cmdable
}

func (c *RedisCaptchaCache) Set(ctx context.Context, biz string, phone string, code string) error {
	ret, err := c.client.Eval(ctx, luaSetCaptcha, []string{c.key(biz, phone)}, code).Int()
	if err != nil {
		return err
	}
	switch ret {
	case 0:
		return nil
	case -1:
		return ErrSetCaptchaTooManyTimes
	default:
		return ErrInternal
	}
}

func (c *RedisCaptchaCache) Verify(ctx context.Context, biz string, phone string, inputCaptcha string) (bool, error) {
	ret, err := c.client.Eval(ctx, luaVerifyCode, []string{c.key(biz, phone)}, inputCaptcha).Int()
	if err != nil {
		return false, err
	}
	switch ret {
	case 0:
		return true, nil
	case -1:
		// TODO LOG
		return false, ErrCaptchaVerifyTooManyTimes
	case -2:
		return false, nil
	default:
		// TODO LOG
		return false, ErrUnknownForCode
	}
}

func (c *RedisCaptchaCache) key(biz string, phone string) string {
	return fmt.Sprintf("phone_captcha:%s:%s", biz, phone)
}

func NewRedisCaptchaCache(client redis.Cmdable) CaptchaCache {
	return &RedisCaptchaCache{
		client: client,
	}
}
