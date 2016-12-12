package gohttp

import "net/http"

// Method represent http method
type Method int8

const (
	// Get represnet http GET method
	Get Method = iota
	// Post represnet http POST method
	Post
	// Put represnet http PUT method
	Put
	// Delete represnet http DELETE method
	Delete
	// Head represent http HEAD method
	Head
	// Option represent http OPTION method
	Option
)

// String will return http method string base on the Method type
func (m Method) String() string {
	switch m {
	case Get:
		return "GET"
	case Post:
		return "POST"
	case Put:
		return "PUT"
	case Delete:
		return "DELETE"
	case Head:
		return "HEAD"
	case Option:
		return "OPTION"
	default:
		return "GET"
	}
}

// Client interface
type Client interface {
	Request(m Method, url string, parameters ...map[string]interface{}) Requester
}

type client struct {
	headers map[string]string
	http.Client
}

// Request will make a request use the given client
func (c *client) Request(m Method, url string, parameters ...map[string]interface{}) Requester {

	req, err := http.NewRequest(m.String(), url, nil)

	var parameter map[string]interface{}
	if len(parameters) > 0 {
		parameter = parameters[0]
	}
	r := request{
		err:        err,
		client:     c,
		request:    req,
		url:        url,
		method:     m,
		parameters: parameter,
	}
	return &r
}
