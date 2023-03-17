package response

import (
	"api-gateway/app/domain/usercases/about/repo"
	"api-gateway/app/util"
)

func NewAboutResponse(about repo.AboutRepo) AboutView {
	aboutView := AboutView{
		Id:    about.Id,
		Name:  about.Name,
		Image: about.Image,
	}
	if about.CreatedAt != nil {
		aboutView.CreatedAt = util.FormatDateTime(*about.CreatedAt)
	}
	if about.UpdatedAt != nil {
		aboutView.UpdatedAt = util.FormatDateTime(*about.UpdatedAt)
	}
	return aboutView
}

func NewAboutResponses(abouts []repo.AboutRepo) []AboutView {
	if len(abouts) <= 0 {
		return []AboutView{}
	}
	aboutViews := make([]AboutView, len(abouts))
	for i := range abouts {
		aboutViews[i] = NewAboutResponse(abouts[i])
	}
	return aboutViews
}
