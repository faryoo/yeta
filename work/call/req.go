package call

type ReqQuery struct {
	Type      int `json:"type"`
	PageSize  int `json:"pageSize"`
	PageIndex int `json:"pageIndex"`
}

type ReqCallOut struct {
	LineNum    string     `json:"line_num,omitempty"`
	RobotId    string     `json:"robot_id,omitempty"`
	CallColumn []string   `json:"call_column"`
	CallList   [][]string `json:"call_list"`
	VoiceCode  string     `json:"voice_code"`
	RobotSpeed int        `json:"robot_speed"`
}

type ReqCreate struct {
	TaskName       string   `json:"task_name"`
	LineNum        string   `json:"line_num"`
	RobotId        string   `json:"robot_id"`
	RecallCount    int      `json:"recall_count"`
	TimeRecallWait int      `json:"time_recall_wait"`
	TimeRange      []string `json:"time_range"`
	TimeBegin      int64    `json:"time_begin"`
	TimeEnd        int64    `json:"time_end"`
}

type ReqInsert struct {
	TaskId     string     `json:"task_id"`
	CallColumn []string   `json:"call_column"`
	CallList   [][]string `json:"call_list"`
}

type TaskID struct {
	TaskId string `json:"task_id"`
}

type ReqTaskQuery struct {
	TaskID         string `json:"task_id"`
	TimeBegin      int    `json:"time_begin"`
	TimeEnd        int    `json:"time_end"`
	TaskName       string `json:"task_name"`
	TaskStatusList []int  `json:"task_status_list"`
	PageSize       int    `json:"page_size"`
	PageIndex      int    `json:"page_index"`
	SortName       string `json:"sort_name"`
	SortOrder      string `json:"sort_order"`
	ShowRemaining  bool   `json:"show_remaining"`
}

type ReqFailed struct {
	Date string `json:"date"`
}
