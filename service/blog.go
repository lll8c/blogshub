package service

import (
	"bloghub/model"
)

func AddBlog(b *model.Blog) error {
	return model.InsertBlog(b)
}

func DeleteBlog(id int64) error {
	return model.DeleteBlogById(id)
}

func BatchDeleteBlog(ids []int64) error {
	return model.BatchDeleteBlogByIds(ids)
}

func UpdateBlog(b *model.Blog) error {
	return model.UpdateBlogById(b)
}

func GetBlog(id int64) (*model.Blog, error) {
	return model.GetBlogByID(id)
}

func GetAllBlog(query *model.Blog) (list []*model.Blog, err error) {
	return model.GetBlogList(query)
}

func GetBlogByPage(blog *model.Blog, pageNum, pageSize int) ([]*model.Blog, error) {
	return model.GetBlogByPage(blog, pageNum, pageSize)
}

func UpdateReadCount(id int64) error {
	return model.UpdateBlogReadCount(id)
}
