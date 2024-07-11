package response

type Response struct {
	Code    int    `json:"Code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	//Header  string      `json:header"`
	Data interface{} `json:"data,omitempty"`
}
