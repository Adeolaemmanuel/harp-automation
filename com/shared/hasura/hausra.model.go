package hasura

type Hasura struct {
	BaseURL string `json:"baseUrl"`
	Headers Headers `json:"headers"`
}

type Headers struct {
	X_Hasura_Admin_Secret string `json:"x-hasura-admin-secret"`
	Content_Type string `json:"Content-Type"`
	Authorization string `json:"authorization"`
	X_Hasura_Role string `json:"X-Hasura-Role"`
}
type HasuraResponse struct {
	Data map[string][]map[string]any `json:"data"`
	Error any `json:"error"`
}

type HasuraEventArgRetryConfig struct {
	Num_Retries int `json:"num_retries"`
	Timeout_Seconds int `json:"timeout_seconds"`
	Tolerance_Seconds int `json:"tolerance_seconds"`
	Retry_Interval_Seconds int `json:"retry_interval_seconds"`

}

type HasuraEventArg struct {
	Webhook string `json:"webhook"`
	Schedule_At string `json:"schedule_at"`
	Type string `json:"type"`
	Event_Id string `json:"event_id"`
	Payload map[string]any `json:"payload"`
	Retry_Conf HasuraEventArgRetryConfig `json:"retry_conf"`
	Comment string `json:"comment"`
}

type HasuraEvent struct {
	Type string `json:"type"`
	Arg HasuraEventArg `json:"arg"`
}

type EventResponse struct {
	Message string `json:"message"`
	Event_Id string `json:"event_id"`
}

type Campaign struct {
	Business_id string `json:"business_id"`
	Channel string `json:"channel"`
	Created_At string `json:"created_at"`
	Status string `json:"status"`
	Message string `json:"message"`
	Title string `json:"title"`
	Id string `json:"id"`
	Batches [][]string `json:"batches"`
	Metadata map[string]any `json:"metadata"`
}