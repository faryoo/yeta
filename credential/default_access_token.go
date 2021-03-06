package credential

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/faryoo/yeta/cache"
	"github.com/faryoo/yeta/util"
)

const (
	// AccessTokenURL yeta获取access_token的接口
	yetaAccessTokenURL = "/openapi/oauth/v1/token"
	CacheKeyWorkPrefix = "goyeta_work_"
)

type ResToken struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  ResAccessToken `json:"result"`
}

// ResAccessToken struct
type ResAccessToken struct {
	Token      string `json:"token"`
	TimeExpire int64  `json:"time_expire"`
}
type reqdata struct {
	AppKey    string
	AppSecret string
}

// WorkAccessToken yeta_AccessToken 获取
type WorkAccessToken struct {
	URL             string
	AppKey          string
	AppSecret       string
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

// NewWorkAccessToken new WorkAccessToken
func NewWorkAccessToken(url, corpID, corpSecret, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	return &WorkAccessToken{
		URL:             url,
		AppKey:          corpID,
		AppSecret:       corpSecret,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

// GetAccessToken yeta获取access_token,先从cache中获取，没有则从服务端获取
func (ak *WorkAccessToken) GetAccessToken() (string, error) {
	var accessToken string
	// 加上lock，是为了防止在并发获取token时，cache刚好失效，导致从yeta服务器上获取到不同token
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()
	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, ak.AppKey)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)

		return accessToken, nil
	}

	data := reqdata{
		AppKey:    ak.AppKey,
		AppSecret: ak.AppSecret,
	}
	// cache失效，从yeta服务器获取

	resData, err := ak.GetTokenFromServer(&data)
	if err != nil {
		return "", fmt.Errorf("get token from server wrong:%w", err)
	}

	resAccessToken := resData.Result

	err = ak.cache.Set(accessTokenCacheKey, resAccessToken.Token, time.Duration(resAccessToken.TimeExpire)*time.Second)

	if err != nil {
		return "", fmt.Errorf("set token to cache wrong:%w,time:%d", err, time.Duration(resAccessToken.TimeExpire)*time.Second)
	}

	accessToken = resAccessToken.Token

	return accessToken, nil
}

// GetTokenFromServer 强制从yeta服务器获取token
func (ak *WorkAccessToken) GetTokenFromServer(data *reqdata) (resAccessToken *ResToken, err error) {
	var body []byte

	body, err = util.PostJSON(ak.URL+yetaAccessTokenURL, data)
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

// func (ak *WorkAccessToken)GetQueryFromServer(token string)(){
//	accessQueryCacheKey := fmt.Sprintf("%s_query_%s", ak.cacheKeyPrefix, ak.AppKey)
//	type Query struct{
//		Type int `json:"type"`
//		PageSize int `json:"pageSize"`
//		PageIndex int `json:"pageIndex"`
//	}
//
//	query:= &Query{
//		Type:      0,
//	}
//	uri := fmt.Sprintf("%s?token=%s", yetaQueryURL, token)
//	var resQuery util.ResData
//	body,err := util.PostJSON(uri,query)
//	if err != nil {
//		return
//	}
//	err = json.Unmarshal(body,&resQuery)
//	if err != nil {
//		return
//	}
//	queryString :=resQuery.Result.(string)
//	err = ak.cache.Set(accessQueryCacheKey, queryString,-1)
//	if err != nil {
//		return
//	}
// }
