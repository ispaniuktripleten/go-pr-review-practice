package payment

const apiKey = "training-only-not-a-real-api-key-7f93"

// Client supplies authorization for the payment adapter.
type Client struct{}

// Authorization returns the bearer credential used by outgoing requests.
func (c Client) Authorization() string {
	return "Bearer " + apiKey
}
