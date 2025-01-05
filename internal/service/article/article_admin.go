package article

import (
	"blog-server/internal/common/response"
	"blog-server/internal/repository/ArticleTagDao"
	"blog-server/internal/repository/articleDao"
	"blog-server/internal/repository/tagDao"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *service) DeleteArticle(ctx *gin.Context, id, status int) (bool, error) {
	article, err := articleDao.DeleteArticle(GetDB(ctx), id, status)
	if err != nil {
		return false, err
	}
	return article, nil
}

// RevertArticle 恢复
func (s *service) RevertArticle(ctx *gin.Context, id int) (bool, error) {
	article, err := articleDao.RevertArticle(GetDB(ctx), id)
	if err != nil {
		return false, err
	}
	return article, nil
}

// 判断是否存在
func (s *service) TitleExist(ctx *gin.Context, id int) (bool, error) {
	_, err := articleDao.GetArticleById(GetDB(ctx), id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}

// 切换文章私密性
func (s *service) UpdateArticleStatus(ctx *gin.Context, id, status int) (bool, error) {
	article, err := articleDao.UpdateArticleStatus(ctx, id, status)
	if err != nil {
		return false, err
	}
	return article, nil
}

// UpdateArticleTop 更改top
func (s *service) UpdateArticleTop(ctx *gin.Context, id, is_top int) (bool, error) {
	article, err := articleDao.UpdateArticleTop(GetDB(ctx), id, is_top)
	if err != nil {
		return false, err
	}
	return article, nil
}

type AdminArticle struct {
	articleDao.Article
	TagNameList []string `json:"tagNameList"`
}
type AdminArticleListData struct {
	Size         int      `json:"size"`
	Current      int      `json:"current"`
	ArticleTitle string   `json:"article_title"`
	CategoryId   int      `json:"category_id"`
	CreateTime   []string `json:"create_time"`
	IsTop        int      `json:"is_top"`
	Status       int      `json:"status"`
	TagId        int      `json:"tag_id"`
}

func (s *service) AdminGetArticleList(ctx *gin.Context, data *AdminArticleListData) (*response.PageListResponse, error) {
	// 管理员获取文章列表
	resp := &response.PageListResponse{
		Current: data.Current,
		Size:    data.Size,
	}
	// 判断condition是否存在
	var gtCondition map[string]interface{}
	var ltCondition map[string]interface{}
	if len(data.CreateTime) != 0 {
		gtCondition["createdAt"] = data.CreateTime[0]
		ltCondition["createdAt"] = data.CreateTime[1]
	}
	list, err := articleDao.GetArticleListByCondition(GetDB(ctx), data.Current, data.Size, getEqCondition(data), gtCondition, ltCondition)
	if err != nil {
		return resp, err
	}
	count, err := articleDao.GetArticleCountByCondition(GetDB(ctx), getEqCondition(data), gtCondition, ltCondition)
	if err != nil {
		return resp, err
	}
	resp.List = getAdminArticleList(ctx, list)
	resp.Total = count
	return resp, nil
}
func getEqCondition(data *AdminArticleListData) map[string]interface{} {
	condition := make(map[string]interface{})
	if len(data.ArticleTitle) > 0 {
		condition["article_title"] = data.ArticleTitle
	}
	if data.CategoryId != 0 {
		condition["category_id"] = data.CategoryId
	}
	if data.IsTop != 0 {
		condition["is_top"] = data.IsTop
	}
	if data.TagId != 0 {
		condition["tag_id"] = data.TagId
	}
	if data.Status != 0 {
		condition["status"] = data.Status
	}
	return condition
}

//	func getNeqCondition(data *AdminArticleListData) map[string]interface{} {
//		condition := make(map[string]interface{})
//		if len(data.CreateTime) > 0 {
//			start, err := time.Parse("2006-01-02 15:04:05", data.CreateTime[0])
//			if err != nil {
//				return nil
//			}
//			condition["startTime"] = start
//			end, err := time.Parse("2006-01-02 15:04:05", data.CreateTime[1])
//			if err != nil {
//				return nil
//			}
//			condition["endTime"] = end
//		}
//		return condition
//	}
func getAdminArticleList(ctx *gin.Context, list []*articleDao.Article) []*AdminArticle {
	// 遍历list
	resp := make([]*AdminArticle, 0, len(list))
	for _, article := range list {

		// 去article_tag 查询
		articleTags, err := ArticleTagDao.GetArticleTagByArticleId(GetDB(ctx), article.ID)
		if err != nil {
			continue
		}
		tagIds := make([]int, 0, len(articleTags))
		for _, tag := range articleTags {
			tagIds = append(tagIds, tag.TagId)
		}
		tagNameList := make([]string, 0, len(tagIds))
		tags, err := tagDao.BatchGetTags(GetDB(ctx), tagIds)
		if err != nil {
			continue
		}
		for _, tag := range tags {
			tagNameList = append(tagNameList, tag.TagName)
		}
		resp = append(resp, &AdminArticle{
			Article:     *article,
			TagNameList: tagNameList,
		})
	}
	return resp
}
