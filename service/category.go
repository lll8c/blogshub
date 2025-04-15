package service

import (
	"bloghub/model"
)

func AddCategory(d *model.Category) error {
	return model.InsertCategory(d)
}

func DeleteCategoryById(id int64) error {
	return model.DeleteCategoryById(id)
}

func BatchDeleteCategory(ids []int64) error {
	return model.BatchDeleteCategoryByIds(ids)
}

func UpdateCategory(d *model.Category) error {
	return model.UpdateCategory(d)
}
