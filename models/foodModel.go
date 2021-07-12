package models

/**
 * 数据库查询的Food对象，转换为前端请求的json对象类型
 */
func (food *Food) FoodToRespDesc() interface{} {
	respDesc := map[string]interface{}{
		"item_id":       food.Id,
		"name":          food.Name,
		"description":   food.Description,
		"rating":        food.Rating,
		"month_sales":   food.MonthSales,
		"activity":      food.Activity,
		"attributes":    food.Attributes,
		"restaurant_id": food.Restaurant.Id,
		"category_id":   food.Category.Id,
	}
	return respDesc
}
