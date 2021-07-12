package models

/**
 * 查询得到的userOrder实体转换为resp的json格式
 */
func (this *UserOrder) UserOrder2Resp() interface{} {
	respDesc := map[string]interface{}{
		"id":                   this.Id,
		"total_amount":         this.SumMoney,
		"user_id":              this.User.UserName,//用户名
		"status":               this.OrderStatus.StatusDesc,//订单状态
		"restaurant_id":        this.Shop.Id, //商铺id
		"restaurant_image_url": this.Shop.ImagePath,//商铺图片
		"restaurant_name":      this.Shop.Name,//商铺名称
		"formatted_created_at": this.Time,
		"status_code":          0,
		"address_id":           this.Address.Id,//订单地址id
	}

	statusDesc := map[string]interface{}{
		"color":     "f60",
		"sub_title": "15分钟内支付",
		"title":     this.OrderStatus.StatusDesc,
	}

	respDesc["status_bar"] = statusDesc
	return respDesc
}
