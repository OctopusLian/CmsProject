package models

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (user *User) UserToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":           user.Id,
		"user_id":      user.Id,
		"username":     user.UserName,
		"city":         user.City.CityName,
		"registe_time": user.RegisteTime,
		"mobile":       user.Mobile,
		"is_active":    user.IsActive,
		"balance":      user.Balance,
		"avatar":       user.Avatar,
	}
	return respInfo
}
