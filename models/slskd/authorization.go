package slskd

type Authorization struct {
	Token     string `json:"token"`
	TokenType string `json:"tokenType"`
	Expires   string `json:"expires"`
	NotBefore string `json:"notBefore"`
}
