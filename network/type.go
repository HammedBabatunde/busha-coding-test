package network

type NetworkResponse struct {
	Body       *[]byte
	StatusCode int
	Error      error
}
