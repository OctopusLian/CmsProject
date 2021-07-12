package models

/**
 * 从数据库中查询出来的实体转变成前端所需要的json格式
 */
func (this *Shop) ShopToRespDesc() interface{} {
	respDesc := map[string]interface{}{
		"id":               this.Id,
		"name":             this.Name,
		"address":          this.Address,
		"phone":            this.Phone,
		"status":           this.Status,
		"recent_order_num": this.RecentOrderNum,
		"rating_count":     this.RatingCount,
		"rating":           this.Rating,
	}
	return respDesc
}
