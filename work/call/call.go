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

// reqMenu 设置菜单请求数据

const (
	callOutURL   = "https://www.xfyeta.com/openapi/outbound/v1/task/callout"
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

// Out 设置按钮
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
	return &resdata, nil
}

// Create Out 设置按钮
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

	return &resdata, nil
}

// Insert  设置按钮
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
	return &resdata, nil
}

// Start Insert  设置按钮
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
	return &resdata, nil
}

// Pause 设置按钮
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
	return &resdata, nil
}

// Delete Pause 设置按钮
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
	return &resdata, nil
}

// TaskQuery Pause 设置按钮
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
	return &resdata, nil
}

// TaskQuery Pause 设置按钮
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
	return &resdata, nil
}
