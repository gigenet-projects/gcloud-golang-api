package api

type Client struct {
	ApiKey string
	secret string
	server string
}

func NewClient(server, key, secret string) *Client {
	return &Client{
		ApiKey: key,
		secret: secret,
		server: server}
}
