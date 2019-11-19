package soap

import (
	"net/http"
	"time"
)

// Response struct holds response values of executed request.
type Response struct {
	Request         *Request
	RawResponse     *http.Response
	payloadResponse []byte
	receivedAt      time.Time
}

// Body method returns HTTP response as []byte array for the executed request.
func (r *Response) PayloadResponse() []byte {
	if r.RawResponse == nil {
		return []byte{}
	}
	return r.payloadResponse
}

// Status method returns the HTTP status string for the executed request.
//	Example: 200 OK
func (r *Response) Status() string {
	if r.RawResponse == nil {
		return ""
	}
	return r.RawResponse.Status
}

// StatusCode method returns the HTTP status code for the executed request.
//	Example: 200
func (r *Response) StatusCode() int {
	if r.RawResponse == nil {
		return 0
	}
	return r.RawResponse.StatusCode
}

// Result method returns the response value as an object if it has one
func (r *Response) PayloadResult() interface{} {
	return r.Request.PayloadResponse
}

// Result method returns the response value as an object if it has one
func (r *Response) PayloadResultError() interface{} {
	return r.Request.PayloadFault
}

// ReceivedAt method returns when response got recevied from server for the request.
func (r *Response) ReceivedAt() time.Time {
	return r.receivedAt
}
