package call

type CommonError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResQuery struct {
	CommonError
	Result Query `json:"result"`
}

type Query struct {
	Urls       []Urls   `json:"urls"`
	RobotTotal int      `json:"robotTotal"`
	Voices     []Voices `json:"voices"`
	Robots     []Robots `json:"robots"`
	Lines      []Lines  `json:"lines"`
}
type Urls struct {
	Url       string `json:"url"`
	UrlModule string `json:"url_module"`
}
type Voices struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	VoiceCode string `json:"voice_code"`
	VoiceName string `json:"voice_name"`
}

type Robots struct {
	CallColumn  []string    `json:"call_column"`
	CreateName  interface{} `json:"create_name"`
	CreatePhone interface{} `json:"create_phone"`
	Deleted     int         `json:"deleted"`
	RobotId     string      `json:"robot_id"`
	RobotName   string      `json:"robot_name"`
	RobotUuid   string      `json:"robot_uuid"`
	Status      int         `json:"status"`
	TimeCreate  int64       `json:"time_create"`
	TimeUpdate  int64       `json:"time_update"`
	Type        int         `json:"type"`
}
type Lines struct {
	Concurrents        int         `json:"concurrents"`
	Expired            int         `json:"expired"`
	LineNum            string      `json:"line_num"`
	RealExpirationTime int64       `json:"real_expiration_time"`
	Status             int         `json:"status"`
	TimeApply          interface{} `json:"time_apply"`
	TimeExpire         int         `json:"time_expire"`
	TimeWork           []string    `json:"time_work"`
}

type ResCallOut struct {
	CommonError
	Result Out `json:"result"`
}

type Out struct {
	Total       int   `json:"total"`
	TaskDataIds []int `json:"task_data_ids"`
}

type ResCreate struct {
	CommonError
	Result Create `json:"result"`
}

type Create struct {
	TaskId string `json:"task_id"`
}

type ResInsert struct {
	CommonError
	Result Insert `json:"result"`
}

type Insert struct {
	Total       int   `json:"total"`
	TaskDataIds []int `json:"task_data_ids"`
}

type ResTaskQuery struct {
	CommonError
	Result TaskQuery `json:"result"`
}

type TaskQuery struct {
	TotalRows int             `json:"total_rows"`
	Rows      []TaskQueryRows `json:"rows"`
}

type TaskQueryRows struct {
	TaskID                string `json:"task_id"`
	TaskName              string `json:"task_name"`
	Status                int    `json:"status"`
	Deleted               int    `json:"deleted"`
	TimeTaskStart         int64  `json:"time_task_start"`
	TimeTaskFinish        int64  `json:"time_task_finish"`
	CountTotalTask        int    `json:"count_total_task"`
	CountTel              int    `json:"count_tel"`
	CountRecalled         int    `json:"count_recalled"`
	TimeTaskEstimateBegin int64  `json:"time_task_estimate_begin"`
	TimeTaskEstimateEnd   int    `json:"time_task_estimate_end"`
	LineNum               string `json:"line_num"`
	RobotID               string `json:"robot_id"`
	RobotName             string `json:"robot_name"`
	VoiceCode             string `json:"voice_code"`
	VoiceSpeed            int    `json:"voice_speed"`
	CountMaxRecall        int    `json:"count_max_recall"`
	StatusRecall          string `json:"status_recall"`
	TimeRecallWait        int    `json:"time_recall_wait"`
	TimeRange             string `json:"time_range"`
	IntentionPush         string `json:"intention_push"`
	ProcessCount          int    `json:"process_count"`
	ProcessTelCount       int    `json:"process_tel_count"`
	ProcessThroughCount   int    `json:"process_through_count"`
	ProcessThroughRate    int    `json:"process_through_rate"`
	Remaining             int    `json:"remaining"`
}

type ResFailed struct {
	CommonError
	Result Failed `json:"result"`
}

type Failed struct {
	Url []string `json:"url"`
}
