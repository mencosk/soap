package soap

import (
	"encoding/xml"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestRequest_Call(t *testing.T) {
	type fields struct {
		Url             string
		Header          http.Header
		PayloadRequest  interface{}
		PayloadResponse interface{}
		PayloadFault    interface{}
		RawRequest      *http.Request
		client          *Client
		Time            time.Time
	}

	// Create dummy service
	go createDummyWebService()

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	response := DummyResponse{}
	fault := DummyFault{}

	responseText := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" >
					   <soapenv:Header/>
					   <soapenv:Body>
						  <Response>
							 <string>Hello World!</string>
						  </Response>
					   </soapenv:Body>
					</soapenv:Envelope>`

	flds := fields{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            time.Time{},
	}

	req := &Request{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            time.Time{},
	}

	resp := &Response{
		Request:         req,
		RawResponse:     nil,
		payloadResponse: []byte(responseText),
		receivedAt:      time.Time{},
	}

	tests := []struct {
		name    string
		fields  fields
		want    *Response
		wantErr bool
	}{
		{
			name:    "Create Request and validate Response",
			fields:  flds,
			want:    resp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &Request{
				Url:             tt.fields.Url,
				Header:          tt.fields.Header,
				PayloadRequest:  tt.fields.PayloadRequest,
				PayloadResponse: tt.fields.PayloadResponse,
				PayloadFault:    tt.fields.PayloadFault,
				RawRequest:      tt.fields.RawRequest,
				client:          tt.fields.client,
				Time:            tt.fields.Time,
			}
			got, err := r.Call()
			if (err != nil) != tt.wantErr {
				t.Errorf("Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.PayloadResult(), tt.want.PayloadResult()) {
				t.Errorf("Call() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetHeader(t *testing.T) {
	type fields struct {
		Url             string
		Header          http.Header
		PayloadRequest  interface{}
		PayloadResponse interface{}
		PayloadFault    interface{}
		RawRequest      *http.Request
		client          *Client
		Time            time.Time
	}
	type args struct {
		header string
		value  string
	}

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	response := DummyResponse{}
	fault := DummyFault{}

	testFields := fields{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            time.Time{},
	}

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

	testArgs := args{
		header: "Content-Type",
		value:  "text/xml; charset=utf-8",
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test Set Header",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Url:             tt.fields.Url,
				Header:          tt.fields.Header,
				PayloadRequest:  tt.fields.PayloadRequest,
				PayloadResponse: tt.fields.PayloadResponse,
				PayloadFault:    tt.fields.PayloadFault,
				RawRequest:      tt.fields.RawRequest,
				client:          tt.fields.client,
				Time:            tt.fields.Time,
			}
			if got := r.SetHeader(tt.args.header, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetHeaders(t *testing.T) {
	type fields struct {
		Url             string
		Header          http.Header
		PayloadRequest  interface{}
		PayloadResponse interface{}
		PayloadFault    interface{}
		RawRequest      *http.Request
		client          *Client
		Time            time.Time
	}
	type args struct {
		headers map[string]string
	}

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	response := DummyResponse{}
	fault := DummyFault{}

	testFields := fields{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            time.Time{},
	}

	testHeader := make(map[string]string)
	testHeader["Content-Type"] = "text/xml; charset=utf-8"

	testArgs := args{
		headers: testHeader,
	}

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

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test Set Headers",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Url:             tt.fields.Url,
				Header:          tt.fields.Header,
				PayloadRequest:  tt.fields.PayloadRequest,
				PayloadResponse: tt.fields.PayloadResponse,
				PayloadFault:    tt.fields.PayloadFault,
				RawRequest:      tt.fields.RawRequest,
				client:          tt.fields.client,
				Time:            tt.fields.Time,
			}
			if got := r.SetHeaders(tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetPayloadFault(t *testing.T) {
	type fields struct {
		Url             string
		Header          http.Header
		PayloadRequest  interface{}
		PayloadResponse interface{}
		PayloadFault    interface{}
		RawRequest      *http.Request
		client          *Client
		Time            time.Time
	}
	type args struct {
		payloadFault interface{}
	}

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	response := DummyResponse{}
	fault := DummyFault{}

	testFields := fields{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            time.Time{},
	}

	testArgs := args{
		payloadFault: &fault,
	}

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

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test Set Fault response",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Url:             tt.fields.Url,
				Header:          tt.fields.Header,
				PayloadRequest:  tt.fields.PayloadRequest,
				PayloadResponse: tt.fields.PayloadResponse,
				PayloadFault:    tt.fields.PayloadFault,
				RawRequest:      tt.fields.RawRequest,
				client:          tt.fields.client,
				Time:            tt.fields.Time,
			}
			if got := r.SetPayloadFault(tt.args.payloadFault); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPayloadFault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetPayloadRequest(t *testing.T) {
	type fields struct {
		Url             string
		Header          http.Header
		PayloadRequest  interface{}
		PayloadResponse interface{}
		PayloadFault    interface{}
		RawRequest      *http.Request
		client          *Client
		Time            time.Time
	}
	type args struct {
		payloadRequest interface{}
	}

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	response := DummyResponse{}
	fault := DummyFault{}

	testFields := fields{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            time.Time{},
	}

	testArgs := args{
		payloadRequest: &request,
	}

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

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{

		{
			name:   "Test set Request",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Url:             tt.fields.Url,
				Header:          tt.fields.Header,
				PayloadRequest:  tt.fields.PayloadRequest,
				PayloadResponse: tt.fields.PayloadResponse,
				PayloadFault:    tt.fields.PayloadFault,
				RawRequest:      tt.fields.RawRequest,
				client:          tt.fields.client,
				Time:            tt.fields.Time,
			}
			if got := r.SetPayloadRequest(tt.args.payloadRequest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPayloadRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetPayloadResponse(t *testing.T) {
	type fields struct {
		Url             string
		Header          http.Header
		PayloadRequest  interface{}
		PayloadResponse interface{}
		PayloadFault    interface{}
		RawRequest      *http.Request
		client          *Client
		Time            time.Time
	}
	type args struct {
		payloadResponse interface{}
	}

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	response := DummyResponse{}
	fault := DummyFault{}

	testFields := fields{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            time.Time{},
	}

	testArgs := args{
		payloadResponse: &response,
	}

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

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test set Response",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Url:             tt.fields.Url,
				Header:          tt.fields.Header,
				PayloadRequest:  tt.fields.PayloadRequest,
				PayloadResponse: tt.fields.PayloadResponse,
				PayloadFault:    tt.fields.PayloadFault,
				RawRequest:      tt.fields.RawRequest,
				client:          tt.fields.client,
				Time:            tt.fields.Time,
			}
			if got := r.SetPayloadResponse(tt.args.payloadResponse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPayloadResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetUrl(t *testing.T) {
	type fields struct {
		Url             string
		Header          http.Header
		PayloadRequest  interface{}
		PayloadResponse interface{}
		PayloadFault    interface{}
		RawRequest      *http.Request
		client          *Client
		Time            time.Time
	}
	type args struct {
		url string
	}

	headers := http.Header{}
	headers.Set("Content-Type", "text/xml; charset=utf-8")

	client := New()
	request := DummyRequest{}
	request.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	request.Body.Request.String = "Hello World!"

	response := DummyResponse{}
	fault := DummyFault{}

	testFields := fields{
		Url:             "http://127.0.0.1:3000/test",
		Header:          headers,
		PayloadRequest:  &request,
		PayloadResponse: &response,
		PayloadFault:    &fault,
		RawRequest:      nil,
		client:          client,
		Time:            time.Time{},
	}

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

	testArgs := args{
		url: "http://127.0.0.1:3000/test",
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Test Set URL",
			fields: testFields,
			args:   testArgs,
			want:   testRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Url:             tt.fields.Url,
				Header:          tt.fields.Header,
				PayloadRequest:  tt.fields.PayloadRequest,
				PayloadResponse: tt.fields.PayloadResponse,
				PayloadFault:    tt.fields.PayloadFault,
				RawRequest:      tt.fields.RawRequest,
				client:          tt.fields.client,
				Time:            tt.fields.Time,
			}
			if got := r.SetUrl(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPointer(t *testing.T) {
	type args struct {
		v interface{}
	}

	testArgs := args{
		v: &DummyRequest{},
	}

	testType := DummyRequest{}
	test2Args := args{
		v: testType,
	}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Test getPointer function with Ptr",
			args: testArgs,
			want: &DummyRequest{},
		},
		{
			name: "Test getPointer function without Ptr",
			args: test2Args,
			want: &testType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPointer(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createDummyWebService() {
	response := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" >
			   <soapenv:Header/>
			   <soapenv:Body>
				  <Response>
					 <string>Hello World!</string>
				  </Response>
			   </soapenv:Body>
			</soapenv:Envelope>`

	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(response)
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.Write([]byte(response))
	})
	http.ListenAndServe(":3000", mux)
}

type DummyRequest struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text    string `xml:",chardata"`
		Request struct {
			Text   string `xml:",chardata"`
			String string `xml:"string"`
		} `xml:"Request"`
	} `xml:"Body"`
}

type DummyResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text     string `xml:",chardata"`
		Response struct {
			Text   string `xml:",chardata"`
			String string `xml:"string"`
		} `xml:"Response"`
	} `xml:"Body"`
}

type DummyFault struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text  string `xml:",chardata"`
		Fault struct {
			Text        string `xml:",chardata"`
			Faultcode   string `xml:"faultcode"`
			Faultstring string `xml:"faultstring"`
			Detail      string `xml:"detail"`
		} `xml:"Fault"`
	} `xml:"Body"`
}
