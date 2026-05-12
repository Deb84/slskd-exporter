package slskd

type Transfer struct {
	Username    string       `json:"username"`
	Directories []*Directory `json:"directories"`
}

type Transfers []*Transfer

type ExtractedTransfers struct {
	Downloads *Transfers
	Uploads   *Transfers
}
