package call

import (
	"encoding/json"
	"fmt"
	"golang/util"
	"golang/util/yeta/context"
	error2 "golang/util/yeta/error"
)

type Call struct {
	*context.Context
}

// reqMenu 设置菜单请求数据
type ReqCall struct {
	LineNum    string     `json:"line_num,omitempty"`
	RobotId    string     `json:"robot_id,omitempty"`
	CallColumn []string   `json:"call_column"`
	CallList   [][]string `json:"call_list"`
	VoiceCode  string     `json:"voice_code"`
	RobotSpeed int        `json:"robot_speed"`
}

const (
	callOutURL = "https://www.xfyeta.com/openapi/outbound/v1/task/callout"
)

func NewCall(context *context.Context) *Call {
	call := new(Call)
	call.Context = context
	return call
}

// CallOut 设置按钮
func (call *Call) CallOut(reqCall *ReqCall) (*error2.ResData, error) {
	accessToken, err := call.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?token=%s", callOutURL, accessToken)

	response, err := util.PostJSON(uri, reqCall)
	if err != nil {
		return nil, err
	}
	var resdata error2.ResData
	err = json.Unmarshal(response, &resdata)
	return &resdata, nil
}
