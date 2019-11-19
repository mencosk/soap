package soap

import (
	"net/http"
)

// Version # of soap
const Version = "1.0.0"

// New method creates a new soap client.
func New() *Client {
	return NewClient(&http.Client{})
}
