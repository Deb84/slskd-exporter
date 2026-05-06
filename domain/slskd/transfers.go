package slskd

type Transfers []Transfer

type Transfer struct {
	Username    string      `json:"username"`
	Directories []Directory `json:"directories"`
}
