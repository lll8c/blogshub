package service

import (
	"bloghub/model"
	"encoding/json"
	"sort"
	"strings"
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

// GetBlog 博客详情
func GetBlog(id int64) (*model.Blog, error) {
	blog, err := model.GetBlogByID(id)
	if err != nil {
		return nil, err
	}
	//查询并设置作者信息
	author := &model.User{}
	author, err = model.GetUserById(blog.UserId)
	if err != nil {
		return nil, err
	}
	//获取作者博客数量、点赞数、收藏数
	blogs, err := model.GetUserBlog(author.Id)
	if err != nil {
		return nil, err
	}
	var userLikeCount int64 = 0
	var userCollectCount int64 = 0
	for _, b := range blogs {
		blogLikeCount, err := model.CountLikesByFidAndModule(b.Id, model.BLOG)
		if err != nil {
			return nil, err
		}
		userLikeCount += blogLikeCount
		blogCollectCount, err := model.CountCollectByFidAndModule(b.Id, model.BLOG)
		if err != nil {
			return nil, err
		}
		userCollectCount += blogCollectCount
	}
	author.LikesCount = userLikeCount
	author.CollectCount = userCollectCount
	blog.User = author

	//设置当前博客的点赞、收藏数
	likeCount, err := model.CountLikesByFidAndModule(blog.Id, model.BLOG)
	if err != nil {
		return nil, err
	}
	collectCount, err := model.CountCollectByFidAndModule(blog.Id, model.BLOG)
	if err != nil {
		return nil, err
	}
	blog.LikesCount = likeCount
	blog.CollectCount = collectCount
	return blog, nil
}

func GetAllBlog(query *model.Blog) (list []*model.Blog, err error) {
	return model.GetBlogList(query)
}

func GetBlogByPage(blog *model.Blog, pageNum, pageSize int) ([]*model.Blog, error) {
	list, err := model.GetBlogByPage(blog, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		likeCount, err := model.CountLikesByFidAndModule(v.Id, model.BLOG)
		if err != nil {
			return nil, err
		}
		blog.LikesCount = likeCount
	}
	return list, nil
}

func UpdateReadCount(id int64) error {
	return model.UpdateBlogReadCount(id)
}

func GetUserBlog(blog *model.Blog, page int, size int) ([]*model.Blog, error) {
	return GetBlogByPage(blog, page, size)
}

func GetUserLikeBlog(blog *model.Blog, page int, size int) ([]*model.Blog, error) {
	return model.GetUserLikeBlogs(blog, page, size)
}

func GetUserCollectBlog(blog *model.Blog, page int, size int) ([]*model.Blog, error) {
	return model.GetUserCollectBlogs(blog, page, size)
}

func GetUserCommentBlog(blog *model.Blog, page int, size int) ([]*model.Blog, error) {
	return model.GetUserCommentBlogs(blog, page, size)
}

// GetTopBlogs todo 热榜算法
func GetTopBlogs() ([]*model.Blog, error) {
	list, err := GetAllBlog(&model.Blog{})
	if err != nil {
		return nil, err
	}
	//使用sort 按 readCount 进行排序
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].ReadCount > list[j].ReadCount
	})
	return list, err
}

func GetRecommendBlog(id int64) ([]*model.Blog, error) {
	blog, err := model.GetBlogByID(id)
	if err != nil {
		return nil, err
	}
	//查询所有博客
	blogList, err := model.GetBlogList(&model.Blog{})
	if err != nil {
		return nil, err
	}
	//根据推荐算法从中筛选五篇
	resList, err := getBlogFromList(blog, blogList)
	return resList, err
}

// 推荐算法
func getBlogFromList(blog *model.Blog, list []*model.Blog) ([]*model.Blog, error) {
	tags := []string{}
	json.Unmarshal([]byte(blog.Tags), &tags)
	var resList []*model.Blog
	for _, v := range list {
		for _, tag := range tags {
			if strings.Contains(v.Tags, tag) {
				if v.Id == blog.Id {
					continue
				}
				resList = append(resList, v)
				if len(resList) > 5 {
					return resList, nil
				}
				break
			}
		}
	}
	return resList, nil
}
