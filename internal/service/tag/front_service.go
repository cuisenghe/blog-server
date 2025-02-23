package tag

import (
	"blog-server/internal/common/response"
	"blog-server/internal/repository/tagDao"
	"github.com/gin-gonic/gin"
)

type ListData struct {
	Current int    `json:"current"`
	Size    int    `json:"size"`
	TagName string `json:"tag_name"`
}
type SimpleTagResp struct {
	ID      int    `json:"id"`
	TagName string `json:"tag_name"`
}

func (s *service) GetTagList(ctx *gin.Context, data *ListData) (*response.PageListResponse, error) {
	tags, err := tagDao.GetTagListByCondition(GetDB(ctx), data.TagName)
	if err != nil {
		return &response.PageListResponse{
			Total:   0,
			List:    nil,
			Current: data.Current,
			Size:    data.Size,
		}, err
	}
	resp := make([]*SimpleTagResp, 0, len(tags))
	for _, tag := range tags {
		resp = append(resp, &SimpleTagResp{
			ID:      tag.ID,
			TagName: tag.TagName,
		})
	}
	count, err := tagDao.GetCountByCondition(GetDB(ctx), data.TagName)
	if err != nil {
		return &response.PageListResponse{
			Total:   0,
			List:    nil,
			Current: data.Current,
			Size:    data.Size,
		}, err
	}
	return &response.PageListResponse{
		Total:   count,
		List:    resp,
		Current: data.Current,
		Size:    data.Size,
	}, nil
}
func (s *service) AddTag(ctx *gin.Context, tagName string) (int, error) {
	tag, err := tagDao.AddTag(GetDB(ctx), tagName)
	if err != nil {
		return 0, err
	}
	return tag.ID, err
}
func (s *service) GetTagDict(ctx *gin.Context) ([]*SimpleTagResp, error) {
	tags, err := tagDao.GetTagListByCondition(GetDB(ctx), "")
	if err != nil {
		return nil, err
	}
	resp := make([]*SimpleTagResp, 0, len(tags))
	for _, tag := range tags {
		resp = append(resp, &SimpleTagResp{
			ID:      tag.ID,
			TagName: tag.TagName,
		})
	}
	return resp, nil
}
