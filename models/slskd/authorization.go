package slskd

type Authorization struct {
	Token     string `json:"token"`
	TokenType string `json:"tokenType"`
	Expires   int64  `json:"expires"`
	NotBefore int64  `json:"notBefore"`
}
