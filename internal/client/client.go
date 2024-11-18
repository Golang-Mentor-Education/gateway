package client

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SendMessage(message string) (string, error) {
	return "Success", nil
}
