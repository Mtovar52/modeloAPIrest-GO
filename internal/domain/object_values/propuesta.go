package objectValues

import (
	"encoding/json"
	"net/http"
)

type interfData interface {
	RespData(OK bool, ruta string, Data interface{}) string

	SetStatusCode(statusCode int)

	SetTitle(title string)

	SetMessage(message string)

	SetData(data interface{})
}

func NewData() interfData {
	return &responseData{}
}

type responseData struct {
	StatusCode int         `json:"status_code"`
	Title      string      `json:"title"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

// implement interfData
func (d *responseData) RespData(OK bool, ruta string, Data interface{}) string {

	AllData := d.messages(OK, ruta, Data)
	if Data == nil {
		respMarshal, _ := json.Marshal(AllData)
		return string(respMarshal)
	}

	respMarshal, _ := json.Marshal(AllData)
	return string(respMarshal)
}

func (d *responseData) messages(OK bool, ruta string, Data interface{}) responseData {

	if OK == false {
		switch ruta {
		case "unauthorized":
			return responseData{
				Title:      "Unauthorized",
				Message:    "Unauthorized",
				StatusCode: http.StatusUnauthorized,
			}

		case "post":
			return responseData{
				Title:      "Bad Request ",
				Message:    "Bad Request",
				StatusCode: http.StatusBadRequest,
			}

		case "delete":
			return responseData{
				Title:      "NotModified",
				Message:    "NotModified",
				StatusCode: http.StatusNotModified,
			}

		case "put":
			return responseData{
				Title:      "NotModified",
				Message:    "NotModified",
				StatusCode: http.StatusNotModified,
			}

		case "get":
		}
	}

	if OK == true {
		switch ruta {
		case "post":
			return responseData{
				Title:      "Creado",
				StatusCode: http.StatusCreated,
			}
		case "delete":
			return responseData{
				Title:      "Modified",
				Message:    "Modified",
				StatusCode: http.StatusOK,
			}
		case "put":
			return responseData{
				Title:      "Modified",
				Message:    "Modified",
				StatusCode: http.StatusOK,
			}
		case "get":
			return responseData{
				Title:      "Modified",
				Message:    "Modified",
				StatusCode: http.StatusOK,
			}
		}
	}

	return responseData{
		Title:      "Data No suministrada",
		Message:    "",
		StatusCode: 402,
	}
}

// SETTER

func (d *responseData) SetStatusCode(statusCode int) {
	d.StatusCode = statusCode
}

func (d *responseData) SetTitle(title string) {
	d.Title = title
}

func (d *responseData) SetMessage(message string) {
	d.Message = message
}

func (d *responseData) SetData(data interface{}) {
	d.Data = data
}
