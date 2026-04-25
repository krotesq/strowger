package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// represents the json that will be returned
type Body struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type Header struct {
	Key   string
	Value string
}

// represents all information that a our response can hold
type Response struct {
	status  int
	body    Body
	cookies []*http.Cookie
	headers []*Header
}

type Builder struct {
	writer        http.ResponseWriter
	headerWritten bool
	response      *Response
}

func NewBuilder(w http.ResponseWriter, response *Response) *Builder {

	if response == nil {
		response = &Response{}
	}

	return &Builder{
		writer:        w,
		headerWritten: false,
		response:      response,
	}
}

func (builder *Builder) SetStatus(status int) {
	builder.response.status = status
}

func (builder *Builder) SetBody(body Body) {
	builder.response.body = body
}

func (builder *Builder) SetCustomCookie(cookie http.Cookie) {
	builder.response.cookies = append(builder.response.cookies, &cookie)
}

func (builder *Builder) SetCookie(name, value string) {
	builder.response.cookies = append(builder.response.cookies, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}

func (builder *Builder) SetHeader(key, value string) {
	builder.response.headers = append(builder.response.headers, &Header{
		Key: key,
		Value: value,
	})
}

func (builder *Builder) Send() {
	// set default header
	builder.SetHeader("Content-Type", "application/json")

	// write headers to buffer
	for _, header := range(builder.response.headers) {
		builder.writer.Header().Set(header.Key, header.Value)
	}

	// write cookies to buffer
	for _, cookie := range(builder.response.cookies) {
		http.SetCookie(builder.writer, cookie)
	}

	// send headers
	if builder.response.status == 0 {
		builder.response.status = 200
	}

	builder.writer.WriteHeader(builder.response.status)

	// send body
	json.NewEncoder(builder.writer).Encode(builder.response.body)

	log.Printf("Response: %d - %s - %v", builder.response.status, builder.response.body.Message, builder.response.body.Data)
}


func SendSimple(w http.ResponseWriter, status int, message string, data any) {
	builder := NewBuilder(w, &Response{
		status: status,
		body: Body{
			Message: message,
			Data: data,
		},
	})
	builder.Send()
}