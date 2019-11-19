package soap

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestResponse_PayloadResponse(t *testing.T) {
	type fields struct {
		Request         *Request
		RawResponse     *http.Response
		payloadResponse []byte
		receivedAt      time.Time
	}
	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	fault := DummyFault{}

	response := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" >
					   <soapenv:Header/>
					   <soapenv:Body>
						  <Response>
							 <string>Hello World!</string>
						  </Response>
					   </soapenv:Body>
					</soapenv:Envelope>`

	testRequest := &Request{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            time.Time{},
	}

	testFields := fields{
		Request:         testRequest,
		RawResponse:     &http.Response{},
		payloadResponse: []byte(response),
		receivedAt:      time.Time{},
	}

	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name:   "Test Retrieve Payload Response",
			fields: testFields,
			want:   []byte(response),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Request:         tt.fields.Request,
				RawResponse:     tt.fields.RawResponse,
				payloadResponse: tt.fields.payloadResponse,
				receivedAt:      tt.fields.receivedAt,
			}
			if got := r.PayloadResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PayloadResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_PayloadResult(t *testing.T) {
	type fields struct {
		Request         *Request
		RawResponse     *http.Response
		payloadResponse []byte
		receivedAt      time.Time
	}

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	fault := DummyFault{}

	response := DummyResponse{}

	textResponse := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" >
					   <soapenv:Header/>
					   <soapenv:Body>
						  <Response>
							 <string>Hello World!</string>
						  </Response>
					   </soapenv:Body>
					</soapenv:Envelope>`

	testRequest := &Request{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            time.Time{},
	}

	testFields := fields{
		Request:         testRequest,
		RawResponse:     &http.Response{},
		payloadResponse: []byte(textResponse),
		receivedAt:      time.Time{},
	}

	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name:   "Test Retrieve PayloadResponse",
			fields: testFields,
			want:   &response,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Request:         tt.fields.Request,
				RawResponse:     tt.fields.RawResponse,
				payloadResponse: tt.fields.payloadResponse,
				receivedAt:      tt.fields.receivedAt,
			}
			if got := r.PayloadResult(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PayloadResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_PayloadResultError(t *testing.T) {
	type fields struct {
		Request         *Request
		RawResponse     *http.Response
		payloadResponse []byte
		receivedAt      time.Time
	}

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	fault := DummyFault{}

	response := DummyResponse{}

	textResponse := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" >
					   <soapenv:Header/>
					   <soapenv:Body>
						  <Response>
							 <string>Hello World!</string>
						  </Response>
					   </soapenv:Body>
					</soapenv:Envelope>`

	testRequest := &Request{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            time.Time{},
	}

	testFields := fields{
		Request:         testRequest,
		RawResponse:     &http.Response{},
		payloadResponse: []byte(textResponse),
		receivedAt:      time.Time{},
	}

	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name:   "Test Retrieve PayloadResultError",
			fields: testFields,
			want:   &fault,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Request:         tt.fields.Request,
				RawResponse:     tt.fields.RawResponse,
				payloadResponse: tt.fields.payloadResponse,
				receivedAt:      tt.fields.receivedAt,
			}
			if got := r.PayloadResultError(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PayloadResultError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_ReceivedAt(t *testing.T) {
	type fields struct {
		Request         *Request
		RawResponse     *http.Response
		payloadResponse []byte
		receivedAt      time.Time
	}

	const longDate = "Jan 2, 2006 at 3:04pm (MST)"
	testTime, _ := time.Parse(longDate, "Nov 19, 2019 at 11:00pm (PST)")

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	fault := DummyFault{}

	response := DummyResponse{}

	textResponse := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" >
					   <soapenv:Header/>
					   <soapenv:Body>
						  <Response>
							 <string>Hello World!</string>
						  </Response>
					   </soapenv:Body>
					</soapenv:Envelope>`

	testRequest := &Request{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            testTime,
	}

	testFields := fields{
		Request:         testRequest,
		RawResponse:     &http.Response{},
		payloadResponse: []byte(textResponse),
		receivedAt:      testTime,
	}

	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name:   "Test Retrieve ReceivedAtTime",
			fields: testFields,
			want:   testTime,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Request:         tt.fields.Request,
				RawResponse:     tt.fields.RawResponse,
				payloadResponse: tt.fields.payloadResponse,
				receivedAt:      tt.fields.receivedAt,
			}
			if got := r.ReceivedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReceivedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_Status(t *testing.T) {
	type fields struct {
		Request         *Request
		RawResponse     *http.Response
		payloadResponse []byte
		receivedAt      time.Time
	}

	const longDate = "Jan 2, 2006 at 3:04pm (MST)"
	testTime, _ := time.Parse(longDate, "Nov 19, 2019 at 11:00pm (PST)")

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	fault := DummyFault{}

	response := DummyResponse{}

	textResponse := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" >
					   <soapenv:Header/>
					   <soapenv:Body>
						  <Response>
							 <string>Hello World!</string>
						  </Response>
					   </soapenv:Body>
					</soapenv:Envelope>`

	testRequest := &Request{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            testTime,
	}

	testFields := fields{
		Request:         testRequest,
		RawResponse:     &http.Response{},
		payloadResponse: []byte(textResponse),
		receivedAt:      testTime,
	}
	testFields.RawResponse.Status = "200"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Test Retrieve Status",
			fields: testFields,
			want:   "200",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Request:         tt.fields.Request,
				RawResponse:     tt.fields.RawResponse,
				payloadResponse: tt.fields.payloadResponse,
				receivedAt:      tt.fields.receivedAt,
			}
			if got := r.Status(); got != tt.want {
				t.Errorf("Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_StatusCode(t *testing.T) {
	type fields struct {
		Request         *Request
		RawResponse     *http.Response
		payloadResponse []byte
		receivedAt      time.Time
	}

	const longDate = "Jan 2, 2006 at 3:04pm (MST)"
	testTime, _ := time.Parse(longDate, "Nov 19, 2019 at 11:00pm (PST)")

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	fault := DummyFault{}

	response := DummyResponse{}

	textResponse := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" >
					   <soapenv:Header/>
					   <soapenv:Body>
						  <Response>
							 <string>Hello World!</string>
						  </Response>
					   </soapenv:Body>
					</soapenv:Envelope>`

	testRequest := &Request{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            testTime,
	}

	// 1. Test 200
	testFields := fields{
		Request:         testRequest,
		RawResponse:     &http.Response{},
		payloadResponse: []byte(textResponse),
		receivedAt:      testTime,
	}
	testFields.RawResponse.StatusCode = http.StatusOK

	// 2. Test 202
	test2Fields := fields{
		Request:         testRequest,
		RawResponse:     &http.Response{},
		payloadResponse: []byte(textResponse),
		receivedAt:      testTime,
	}
	test2Fields.RawResponse.StatusCode = http.StatusAccepted

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Test Retrieve StatusCode 200",
			fields: testFields,
			want:   http.StatusOK,
		},
		{
			name:   "Test Retrieve StatusCode 202",
			fields: test2Fields,
			want:   http.StatusAccepted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Request:         tt.fields.Request,
				RawResponse:     tt.fields.RawResponse,
				payloadResponse: tt.fields.payloadResponse,
				receivedAt:      tt.fields.receivedAt,
			}
			if got := r.StatusCode(); got != tt.want {
				t.Errorf("StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
