package soap

import (
	"net"
	"net/http"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestClient_NewRequest(t *testing.T) {
	type fields struct {
		httpClient *http.Client
	}

	httpClient := &http.Client{
		Transport: createTransport(nil),
	}

	tests := []struct {
		name   string
		fields fields
		want   *Request
	}{
		{
			name: "Create NewRequest()",
			fields: fields{
				httpClient: httpClient,
			},
			want: &Request{
				client: &Client{
					httpClient: httpClient,
				},
				Header: http.Header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			if got := c.NewRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_R(t *testing.T) {
	type fields struct {
		httpClient *http.Client
	}

	httpClient := &http.Client{
		Transport: createTransport(nil),
	}

	tests := []struct {
		name   string
		fields fields
		want   *Request
	}{
		{
			name: "Create New R()",
			fields: fields{
				httpClient: httpClient,
			},
			want: &Request{
				client: &Client{
					httpClient: httpClient,
				},
				Header: http.Header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			if got := c.R(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SetTimeOut(t *testing.T) {
	type fields struct {
		httpClient *http.Client
	}
	type args struct {
		timeout time.Duration
	}

	httpClient := &http.Client{
		Transport: createTransport(nil),
	}

	client := &Client{
		httpClient: httpClient,
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Client
	}{
		{
			name: "Crete SetTimeOut()",
			fields: fields{
				httpClient: httpClient,
			},
			args: args{timeout: 10 * time.Second},
			want: client,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpClient: tt.fields.httpClient,
			}
			if got := c.SetTimeOut(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTimeOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		hc *http.Client
	}

	httpClient := &http.Client{
		Transport: createTransport(nil),
	}

	arguments := args{
		hc: httpClient,
	}

	client := &Client{
		httpClient: httpClient,
	}

	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "Create NewClient()",
			args: arguments,
			want: client,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.hc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createTransport(t *testing.T) {
	type args struct {
		httpTransport *http.Transport
	}

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	httpTransport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}
	arguments := args{
		httpTransport: httpTransport,
	}

	tests := []struct {
		name string
		args args
		want *http.Transport
	}{
		{
			name: "Create Transport()",
			args: arguments,
			want: httpTransport,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTransport(tt.args.httpTransport); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTransport() = %v, want %v", got, tt.want)
			}
		})
	}
}
