package core

// Ref: http://nginx.org/en/docs/http/ngx_http_core_module.html

import (
	access "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_access"
	proxy "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_proxy"
)

type HttpContext struct {
	ClientPropsHttp
	access.HttpAccessContext

	ErrorPageContext

	// Proxy
	Proxy []*proxy.HttpProxy `json:"proxy"`
}

type ServerContext struct {
	ServerNames []string `json:"server_names"`
	Listens     []string `json:"listen"`

	ClientPropsServer
	access.HttpAccessContext

	// Properties
	CoreProps ClientPropsServer `json:"core_props"`
	Proxy     proxy.HttpProxy   `json:"proxy"`

	// ErrorPageContext
	ErrorPageContext

	// Paths
	Location []*LocationContext `json:"location"`
}

type ClientPropsServer struct {
	AbsoluteRedirect        string `json:"absolute_redirect"`
	Aio                     string `json:"aio"`
	AioWrite                string `json:"aio_write"`
	ChunkedTransferEncoding string `json:"chunked_transfer_encoding"`
	ClientBodyBufferSize    string `json:"client_body_buffer_size"`
	KeepaliveRequests       string `json:"keepalive_requests"`
	ClientMaxBodySize       string `json:"client_max_body_size"`
}

type LocationContext struct {
	CoreProps ClientPropsLocation `json:"core_props"`
	access.HttpAccessContext

	// Paths
	Location []string `json:"location"`

	// Error page
	ErrorPageContext
}

// define error_page
// Syntax: error_page code ... [=[response]] uri;
type ErrorPageContext struct {
	// code can be a number or "*"
	Codes []int `json:"code"`

	// Return code if any
	Response string `json:"response"`

	// uri can be text, variable, or a URI
	URI string `json:"uri"`
}

type ClientPropsHttp struct {
	AbsoluteRedirect        string `json:"absolute_redirect"`
	Aio                     string `json:"aio"`
	AioWrite                string `json:"aio_write"`
	ChunkedTransferEncoding string `json:"chunked_transfer_encoding"`
	ClientBodyBufferSize    string `json:"client_body_buffer_size"`
	KeepaliveRequests       string `json:"keepalive_requests"`
	ClientMaxBodySize       string `json:"client_max_body_size"`
}

type ClientPropsLocation struct {
	AbsoluteRedirect        string `json:"absolute_redirect"`
	Aio                     string `json:"aio"`
	AioWrite                string `json:"aio_write"`
	ChunkedTransferEncoding string `json:"chunked_transfer_encoding"`
	ClientBodyBufferSize    string `json:"client_body_buffer_size"`
	KeepaliveRequests       string `json:"keepalive_requests"`
	ClientMaxBodySize       string `json:"client_max_body_size"`
}
