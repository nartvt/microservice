package response

type AboutView struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Image     string `json:"image,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	Active    bool   `json:"active,omitempty"`
}
