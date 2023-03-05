package network

type NetworkResponse struct {
	Body       interface{}
	StatusCode int
	Error      error
}
