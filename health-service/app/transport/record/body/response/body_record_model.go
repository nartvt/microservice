package response

type UserBodyRecordResponse struct {
	Id         int     `json:"id,omit"`
	UserId     int     `json:"userId"`
	Weight     float32 `json:"weight"`
	Height     int     `json:"height"`
	Percentage float32 `json:"percentage"`
	Date       string  `json:"createdAt"`
}
