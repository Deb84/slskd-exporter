package slskd

type Directory struct {
	Directory string  `json:"directories"`
	FileCount int     `json:"fileCount"`
	Files     []*File `json:"files"`
}
