package objectValues

type responseWithData struct {
	StatusCode int         `json:"status_code"`
	Title      string      `json:"title"`
	Message    string      `json:"message"`
	Data       interface{} `json:",omitempty"`
}

func NewResponseWithData(StatusCode int, Title string, Message string, Data interface{}) responseWithData {

	return responseWithData{
		StatusCode: StatusCode,
		Title:      Title,
		Message:    Message,
		Data:       Data,
	}
}

//	if Data == nil {
//	values := responseWithData{StatusCode: StatusCode,
//		Title:      Title,
//		Message:    Message,}
//	respMarshal,_:= json.Marshal(values)

//}
