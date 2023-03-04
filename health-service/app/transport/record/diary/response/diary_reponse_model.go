package response

type UserDiaryResponse struct {
	Id          int    `json:"id,omitempty"`
	UserId      int    `json:"userId,omitempty"`
	AtTime      int    `json:"atTime,omitempty"`
	Description string `json:"description,omitempty"`
	Calories    int    `json:"calories,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}
