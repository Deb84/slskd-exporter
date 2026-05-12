package slskd

type File struct {
	Id               string  `json:"id"`
	Username         string  `json:"username"`
	Direction        string  `json:"direction"`
	FileName         string  `json:"filename"`
	Size             int64   `json:"size"`
	StartOffset      int     `json:"startOffset"`
	State            string  `json:"state"`
	StateDescription string  `json:"stateDescription"`
	RequestedAt      string  `json:"requestedAt"`
	EnqueuedAt       string  `json:"enqueuedAt"`
	StartedAt        string  `json:"startedAt"`
	EndedAt          string  `json:"endedAt"`
	BytesTransferred int64   `json:"bytesTransferred"`
	AverageSpeed     float64 `json:"averageSpeed"`
	Exception        string  `json:"exception"`
	Attempts         int     `json:"attempts"`
	BytesRemaining   int64   `json:"bytesRemaining"`
	ElapsedTime      string  `json:"elapsedTime"`
	PercentComplete  float64 `json:"percentComplete"`
	RemainingTime    string  `json:"remainingTime"`
}
