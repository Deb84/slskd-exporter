package slskd

type File struct {
	Id               string  `json:"id"`
	Username         string  `json:"username"`
	Direction        string  `json:"direction"`
	FileName         string  `json:"filename"`
	Size             int     `json:"size"`
	StartOffset      int     `json:"startOffset"`
	State            string  `json:"state"`
	StateDescription string  `json:"stateDescription"`
	RequestedAt      string  `json:"requestedAt"`
	EnqueuedAt       string  `json:"enqueuedAt"`
	StartedAt        string  `json:"startedAt"`
	EndedAt          string  `json:"endedAt"`
	BytesTransferred int     `json:"bytesTransferred"`
	AverageSpeed     float64 `json:"averageSpeed"`
	Attempts         int     `json:"attempts"`
	BytesRemaining   int     `json:"bytesRemaining"`
	ElapsedTime      string  `json:"elaspedTime"`
	PercentComplete  float64 `json:"percentComplete"`
	RemainingTime    string  `json:"remainingTime"`
}
