package client

// APIClient is an exported type that contains the credentials used
// by a connecting client to authenticate with heimdall
type APIClient struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
