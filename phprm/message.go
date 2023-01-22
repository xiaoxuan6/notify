package phprm

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		MessageIdList []string `json:"messageIdList"`
	} `json:"data"`
}
