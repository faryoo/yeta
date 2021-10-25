package call

import (
	"encoding/json"
	"fmt"

	"github.com/faryoo/yeta/util"
	"github.com/faryoo/yeta/work/context"
)

type Call struct {
	*context.Context
}

const (
	callOutURL   = "/openapi/outbound/v1/task/callout"
	queryURL     = "/openapi/config/v1/query"
	createURL    = "/openapi/outbound/v1/task/create"
	insertURL    = "/openapi/outbound/v1/task/insert"
	startURL     = "/openapi/outbound/v1/task/start"
	pauseURL     = "/openapi/outbound/v1/task/pause"
	deleteURL    = "/openapi/outbound/v1/task/delete"
	taskQueryURL = "/openapi/outbound/v1/task/query"
	failedURL    = "/openapi/download/v1/push/failed"
)

func NewCall(context *context.Context) *Call {
	call := new(Call)
	call.Context = context

	return call
}

// CallOut 直接外呼.
func (call *Call) CallOut(reqCall *ReqCallOut) (*ResCallOut, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get accesstoken wrong : %w", err)
	}

	uri := fmt.Sprintf("%s?token=%s", call.URL+callOutURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, fmt.Errorf("post json wrong : %w", err)
	}

	var resdata ResCallOut
	err = json.Unmarshal(response, &resdata)

	if err != nil {
		return nil, fmt.Errorf("json unmarshal wrong : %w", err)
	}

	return &resdata, nil
}

// Query 查询配置.
func (call *Call) Query(reqCall *ReqQuery) (*ResQuery, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", call.URL+queryURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}

	var resdata ResQuery
	err = json.Unmarshal(response, &resdata)

	return &resdata, err
}

// Create 创建外呼任务.
func (call *Call) Create(reqCall *ReqCreate) (*ResCreate, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", call.URL+createURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}

	var resdata ResCreate
	err = json.Unmarshal(response, &resdata)

	return &resdata, err
}

// Insert  提交任务数据.
func (call *Call) Insert(reqCall *ReqInsert) (*ResInsert, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", call.URL+insertURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata ResInsert
	err = json.Unmarshal(response, &resdata)

	return &resdata, err
}

// Start 启动外呼任务.
func (call *Call) Start(reqCall *TaskID) (*CommonError, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", call.URL+startURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata CommonError
	err = json.Unmarshal(response, &resdata)
	return &resdata, err
}

// Pause 暂停外呼任务.
func (call *Call) Pause(reqCall *TaskID) (*CommonError, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", call.URL+pauseURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata CommonError
	err = json.Unmarshal(response, &resdata)
	return &resdata, err
}

// Delete 删除外呼任务.
func (call *Call) Delete(reqCall *TaskID) (*CommonError, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", call.URL+deleteURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata CommonError
	err = json.Unmarshal(response, &resdata)
	return &resdata, err
}

// TaskQuery 查询任务.
func (call *Call) TaskQuery(reqCall *ReqTaskQuery) (*ResTaskQuery, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", call.URL+taskQueryURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata ResTaskQuery
	err = json.Unmarshal(response, &resdata)
	return &resdata, err
}

// Failed 查询推送失败记录.
func (call *Call) Failed(reqCall *ReqFailed) (*ResFailed, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", call.URL+failedURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata ResFailed
	err = json.Unmarshal(response, &resdata)
	return &resdata, err
}
