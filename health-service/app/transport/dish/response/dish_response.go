package response

import (
	"health-service/app/domain/usercases/dish/repo"
	"health-service/app/util"
)

func NewDishResponse(dish repo.DishRepo) DishView {
	dishView := DishView{
		Id:    dish.Id,
		Name:  dish.Name,
		Image: dish.Image,
	}
	if dish.CreatedAt != nil {
		dishView.CreatedAt = util.FormatDateTime(*dish.CreatedAt)
	}
	if dish.UpdatedAt != nil {
		dishView.UpdatedAt = util.FormatDateTime(*dish.UpdatedAt)
	}
	return dishView
}

func NewDishResponses(dishes []repo.DishRepo) []DishView {
	if len(dishes) <= 0 {
		return []DishView{}
	}
	dishViews := make([]DishView, len(dishes))
	for i := range dishes {
		dishViews[i] = NewDishResponse(dishes[i])
	}
	return dishViews
}
