package response

import (
	"api-gateway/app/domain/usercases/section/repo"
)

func NewSectionResponse(newsfeed repo.NewsfeedSectionRepo) NewsfeedSection {
	return NewsfeedSection{
		Id:   newsfeed.Id,
		Name: newsfeed.Name,
	}
}

func NewSectionResponses(newsfeeds []repo.NewsfeedSectionRepo) []NewsfeedSection {
	if len(newsfeeds) <= 0 {
		return []NewsfeedSection{}
	}
	resp := make([]NewsfeedSection, len(newsfeeds))
	for i := range newsfeeds {
		resp[i] = NewSectionResponse(newsfeeds[i])
	}
	return resp
}
