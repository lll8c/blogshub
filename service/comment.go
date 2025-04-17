package service

import (
	"bloghub/model"
)

func AddComment(c *model.Comment) error {
	err := model.InsertComment(c)
	if err != nil {
		return err
	}
	//如果是根评论还要更新 rootID 为自身id
	//先更新完后才能获得id
	if c.RootId == 0 {
		c.RootId = c.Id
		err := model.UpdateCommentById(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteComment(id int64) error {
	return model.DeleteCommentById(id)
}

func BatchDeleteComment(ids []int64) error {
	return model.BatchDeleteComment(ids)
}

func UpdateComment(c *model.Comment) error {
	return model.UpdateCommentById(c)
}

func GetComment(id int64) (*model.Comment, error) {
	return model.GetCommentById(id)
}

func GetCommentList(c *model.Comment) ([]*model.Comment, error) {
	return model.GetCommentList(c)
}

func GetCommentListByPage(c *model.Comment, num int, size int) ([]*model.Comment, error) {
	return model.GetCommentListByPage(c, num, size)
}

func GetUserCommentList(c *model.Comment) ([]*model.Comment, error) {
	//先查询所有一级评论
	list, err := model.GetRootCommentList(c)
	if err != nil {
		return nil, err
	}
	//查询每个顶级评论下的所有回复
	for _, v := range list {
		query := &model.Comment{RootId: v.Id}
		children, err := model.GetCommentList(query)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(children); i++ {
			if children[i].RootId == v.Id { //排除根节点
				children = append(children[:i], children[i+1:]...)
			}
		}
		v.Children = children
	}
	return list, nil
}

func CountCommentByFid(fid int64, module string) (int64, error) {
	return model.CountCommentByFid(fid, module)
}
