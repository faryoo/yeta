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
	callOutURL   = "https://www.xfyeta.com/openapi/outbound/v1/task/callout"
	queryURL     = "https://www.xfyeta.com/openapi/config/v1/query"
	createURL    = "https://www.xfyeta.com/openapi/outbound/v1/task/create"
	insertURL    = "https://www.xfyeta.com/openapi/outbound/v1/task/insert"
	startURL     = "https://www.xfyeta.com/openapi/outbound/v1/task/start"
	pauseURL     = "https://www.xfyeta.com/openapi/outbound/v1/task/pause"
	deleteURL    = "https://www.xfyeta.com/openapi/outbound/v1/task/delete"
	taskQueryURL = "https://www.xfyeta.com/openapi/outbound/v1/task/query"
	failedURL    = "https://www.xfyeta.com/openapi/download/v1/push/failed"
)

func NewCall(context *context.Context) *Call {
	call := new(Call)
	call.Context = context
	return call
}

// CallOut 直接外呼
func (call *Call) CallOut(reqCall *ReqCallOut) (*ResCallOut, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", callOutURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata ResCallOut
	err = json.Unmarshal(response, &resdata)
	return &resdata, err
}

// Query 查询配置
func (call *Call) Query(reqCall *ReqQuery) (*ResQuery, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", queryURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata ResQuery
	err = json.Unmarshal(response, &resdata)

	return &resdata, err
}

// Create 创建外呼任务
func (call *Call) Create(reqCall *ReqCreate) (*ResCreate, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", createURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata ResCreate
	err = json.Unmarshal(response, &resdata)

	return &resdata, err
}

// Insert  提交任务数据
func (call *Call) Insert(reqCall *ReqInsert) (*ResInsert, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", insertURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata ResInsert
	err = json.Unmarshal(response, &resdata)
	return &resdata, err
}

// Start 启动外呼任务
func (call *Call) Start(reqCall *TaskID) (*CommonError, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", startURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata CommonError
	err = json.Unmarshal(response, &resdata)
	return &resdata, err
}

// Pause 暂停外呼任务
func (call *Call) Pause(reqCall *TaskID) (*CommonError, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", pauseURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata CommonError
	err = json.Unmarshal(response, &resdata)
	return &resdata, err
}

// Delete 删除外呼任务
func (call *Call) Delete(reqCall *TaskID) (*CommonError, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", deleteURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata CommonError
	err = json.Unmarshal(response, &resdata)
	return &resdata, err
}

// TaskQuery 查询任务
func (call *Call) TaskQuery(reqCall *ReqTaskQuery) (*ResTaskQuery, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", taskQueryURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata ResTaskQuery
	err = json.Unmarshal(response, &resdata)
	return &resdata, err
}

// Failed 查询推送失败记录
func (call *Call) Failed(reqCall *ReqFailed) (*ResFailed, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", failedURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata ResFailed
	err = json.Unmarshal(response, &resdata)
	return &resdata, err
}
