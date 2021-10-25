package yeta

import (
	"github.com/faryoo/yeta/credential"
	"github.com/faryoo/yeta/work/call"
	"github.com/faryoo/yeta/work/config"
	"github.com/faryoo/yeta/work/context"
)

// Work yeta工作实例.
type Work struct {
	ctx *context.Context
}

// NewWork init work.
func NewWork(cfg *config.Config) *Work {
	defaultAkHandle := credential.NewWorkAccessToken(cfg.URL, cfg.AppKey, cfg.AppSecret, credential.CacheKeyWorkPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}

	return &Work{ctx: ctx}
}

// SetAccessTokenHandle 自定义access_token获取方式.
func (wk *Work) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	wk.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context.
func (wk *Work) GetContext() *context.Context {
	return wk.ctx
}

// GetAccessToken 获取access_token.
func (wk *Work) GetAccessToken() (string, error) {
	return wk.ctx.GetAccessToken() //nolint:wrapcheck
}

func (wk *Work) GetCall() *call.Call {
	return call.NewCall(wk.ctx)
}
