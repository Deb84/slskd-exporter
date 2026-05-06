package slskd

type Transfers struct {
	Transfers []Transfer
}

type Transfer struct {
	Username    string      `json:"username"`
	Directories []Directory `json:"directories"`
}
