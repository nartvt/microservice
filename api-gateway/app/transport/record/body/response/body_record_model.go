package response

type UserBodyRecordResponse struct {
	Id         int     `json:"id,omitempty"`
	UserId     int     `json:"userId,omitempty"`
	Weight     float32 `json:"weight,omitempty"`
	Height     int     `json:"height,omitempty"`
	Percentage float32 `json:"percentage,omitempty"`
	Date       string  `json:"createdAt,omitempty"`
}
