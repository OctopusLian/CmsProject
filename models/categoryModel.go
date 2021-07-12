package models

import "github.com/astaxie/beego/orm"

const CATEGORYTABLE = "food_category"

func (category *FoodCategory) CategoryToStringResp() interface{} {
	respDesc := map[string]interface{}{
		"name":  category.CategoryName,
		"id":    category.Id,
		"count": category.Id,
	}

	om := orm.NewOrm()
	var subCategory []*FoodCategory
	om.QueryTable(CATEGORYTABLE).Filter("parent_category_id", category.Id).All(&subCategory)

	var subString []interface{}
	for _, subCate := range subCategory {
		subString = append(subString, subCate)
	}
	respDesc["sub_categories"] = subString
	return respDesc
}
