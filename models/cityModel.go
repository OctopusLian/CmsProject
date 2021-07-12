package models

/**  20个字段                   5个字段
 * 数据库查询到的model实体转换成前端需要的json格式
 */
func (this *City) CityToRespDesc() interface{} {
	respDesc := map[string]interface{}{
		"pinyin":    this.PinYin,
		"longitude": this.Longitude,
		"latitude":  this.Latitude,
		"area_code": this.AreaCode,
		"abbr":      this.Abbr,
		"name":      this.CityName,
		"id":        this.Id,
	}
	return respDesc
}
