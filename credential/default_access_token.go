package credential

import (
	"encoding/json"
	"fmt"
	"golang/util"
	"golang/util/yeta/cache"
	errord "golang/util/yeta/error"
	"sync"
	"time"
)

const (
	// AccessTokenURL 企业微信获取access_token的接口
	workAccessTokenURL = "https://www.xfyeta.com/openapi/oauth/v1/token?corpid=%s&corpsecret=%s"
	// CacheKeyWorkPrefix 企业微信cache key前缀
	CacheKeyWorkPrefix = "goyeta_work_"
)

// ResAccessToken struct
type ResAccessToken struct {
	Token      string `json:"token"`
	TimeExpire int64  `json:"time_expire"`
}
type reqdata struct {
	AppKey    string
	AppSecret string
}

// WorkAccessToken 企业微信AccessToken 获取
type WorkAccessToken struct {
	AppKey          string
	AppSecret       string
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

// NewWorkAccessToken new WorkAccessToken
func NewWorkAccessToken(corpID, corpSecret, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	return &WorkAccessToken{
		AppKey:          corpID,
		AppSecret:       corpSecret,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

// GetAccessToken 企业微信获取access_token,先从cache中获取，没有则从服务端获取
func (ak *WorkAccessToken) GetAccessToken() (accessToken string, err error) {
	// 加上lock，是为了防止在并发获取token时，cache刚好失效，导致从微信服务器上获取到不同token
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()
	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, ak.AppKey)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	data := reqdata{
		AppKey:    ak.AppKey,
		AppSecret: ak.AppSecret,
	}
	// cache失效，从微信服务器获取
	var resData *errord.ResData
	resData, err = GetTokenFromServer(&data)
	if err != nil {
		return
	}
	j, _ := json.Marshal(resData.Result)
	resAccessToken := new(ResAccessToken)
	_ = json.Unmarshal(j, resAccessToken)

	expires := resAccessToken.TimeExpire - 1500
	err = ak.cache.Set(accessTokenCacheKey, resAccessToken.Token, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}
	accessToken = resAccessToken.Token
	return
}

// GetTokenFromServer 强制从微信服务器获取token
func GetTokenFromServer(data *reqdata) (resAccessToken *errord.ResData, err error) {
	var body []byte

	body, err = util.HTTPPost(workAccessTokenURL, data)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resAccessToken)
	if err != nil {
		return
	}
	if resAccessToken.Code != 0 {
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resAccessToken.Code, resAccessToken.Message)
		return
	}
	return
}
