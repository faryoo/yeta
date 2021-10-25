// Package config 企业微信config配置
package config

import "github.com/faryoo/yeta/cache"

// Config for 企业微信.
type Config struct {
	URL       string `json:"url"`
	AppKey    string `json:"app_key"`    // corp_id
	AppSecret string `json:"app_secret"` // corp_secret,如果需要获取会话存档实例，当前参数请填写聊天内容存档的Secret，可以在企业微信管理端--管理工具--聊天内容存档查看
	Cache     cache.Cache
}
