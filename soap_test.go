package soap

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Client
	}{
		{
			name: "Create New client",
			want: NewClient(&http.Client{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
