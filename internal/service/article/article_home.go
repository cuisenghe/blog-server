package article

import (
	"blog-server/internal/common/response"
	"blog-server/internal/repository/articleDao"
	"blog-server/internal/repository/tagDao"
	"github.com/gin-gonic/gin"
	"time"
)

type ArticleListData struct {
	Current int    `form:"current"`
	Size    int    `form:"size"`
	Order   string `form:"order"`
	Content string `form:"content"`
	Id      int    `form:"id"`
}
type SimpleArticleList struct {
	Year        string           `json:"year"`
	ArticleList []*SimpleArticle `json:"articleList"`
}
type SimpleArticle struct {
	ID           int       `json:"id"`
	ArticleTitle string    `json:"article_title"`
	ArticleCover string    `json:"article_cover"`
	CreatedAt    time.Time `json:"createdAt"`
}
type RecommendArticleList struct {
	Year        string           `json:"year"`
	ArticleList []*SimpleArticle `json:"articleList"`
}
type ContentArticleList struct {
	Year        string           `json:"year"`
	ArticleList []*SimpleArticle `json:"articleList"`
}
type DetailArticle struct {
	*articleDao.Article
	TagIdList   []int    `json:"tagIdList"`
	TagNameList []string `json:"tagNameList"`
}

func (s *service) GetArticleList(ctx *gin.Context, req *ArticleListData) (*response.PageListResponse, error) {
	resp := &response.PageListResponse{
		Current: req.Current,
		Size:    req.Size,
	}
	if req.Size == 0 {
		return resp, nil
	}
	list, err := articleDao.GetArticleList(GetDB(ctx), req.Current, req.Size)
	if err != nil {
		return resp, err
	}
	count, err := articleDao.GetSumCount(GetDB(ctx))
	if err != nil {
		return resp, err
	}

	resp.List = list
	resp.Total = count
	return resp, nil
}

// BlogTimelineGetArticleList 获取时间线文章
func (s *service) BlogTimelineGetArticleList(ctx *gin.Context, data *ArticleListData) (*response.PageListResponse, error) {
	// 获取文章
	resp := &response.PageListResponse{
		Current: data.Current,
		Size:    data.Size,
	}
	list, err := articleDao.GetArticleList(GetDB(ctx), data.Current, data.Size)
	if err != nil {
		return nil, err
	}
	count, err := articleDao.GetSumCount(GetDB(ctx))
	if err != nil {
		return nil, err
	}
	resp.List = convertTimeLineData(list)
	resp.Total = count
	return resp, nil
}
func convertTimeLineData(list []*articleDao.Article) []*SimpleArticleList {
	// 分组
	simpleList := make([]*SimpleArticleList, 0)
	m := make(map[string]*SimpleArticleList)
	for _, article := range list {
		year := article.CreateAt.Format("2006")
		// 获取对应的simpleArticleList
		if _, ok := m[year]; !ok {
			m[year] = &SimpleArticleList{
				Year:        year,
				ArticleList: make([]*SimpleArticle, 0),
			}
		}
		articleList := m[year].ArticleList
		articleList = append(articleList, &SimpleArticle{
			ID:           article.ID,
			ArticleTitle: article.ArticleTitle,
			ArticleCover: article.ArticleCover,
			CreatedAt:    article.CreateAt,
		})
		m[year].ArticleList = articleList
	}
	// 遍历
	for _, v := range m {
		simpleList = append(simpleList, v)
	}
	return simpleList
}
func (s *service) GetArticleListByTagId(ctx *gin.Context, req *ArticleListData) (*response.PageListResponse, error) {
	// 根据条件查询文章
	resp := &response.PageListResponse{
		Current: req.Current,
		Size:    req.Size,
	}
	condition, err := articleDao.GetArticleListByCondition(GetDB(ctx), req.Current, req.Size, map[string]interface{}{
		"tag_id": req.Id,
	}, nil, nil)
	if err != nil {
		return resp, err
	}
	resp.List = convertTimeLineData(condition)
	return resp, nil
}
func (s *service) GetArticleListByCategoryId(ctx *gin.Context, req *ArticleListData) (*response.PageListResponse, error) {
	resp := &response.PageListResponse{
		Current: req.Current,
		Size:    req.Size,
	}
	condition, err := articleDao.GetArticleListByCondition(GetDB(ctx), req.Current, req.Size, map[string]interface{}{
		"category_id": req.Id,
	}, nil, nil)
	if err != nil {
		return resp, err
	}
	resp.List = convertTimeLineData(condition)
	return resp, nil
}

// 获取推荐
func (s *service) GetRecommendArticleById(ctx *gin.Context, articleId int) (*response.PageListResponse, error) {
	// 获取推荐
	return &response.PageListResponse{}, nil
}

// 根据内容搜索
func (s *service) GetArticleListByContent(ctx *gin.Context, content string) (*response.PageListResponse, error) {
	resp := &response.PageListResponse{
		Current: 1,
		Size:    5,
	}
	articleList, err := articleDao.GetArticleListByContent(GetDB(ctx), content)
	if err != nil {
		return resp, err
	}
	// 获取count
	count, err := articleDao.GetArticleCountByContent(GetDB(ctx), content)
	if err != nil {
		return resp, err
	}
	resp.List = convertSearchData(articleList)
	resp.Total = count
	return resp, nil
}
func convertSearchData(list []*articleDao.Article) *ContentArticleList {
	articles := make([]*SimpleArticle, 0, len(list))
	for _, article := range list {
		articles = append(articles, &SimpleArticle{
			ID:           article.ID,
			ArticleTitle: article.ArticleTitle,
			ArticleCover: article.ArticleCover,
			CreatedAt:    article.CreateAt,
		})
	}
	//
	return &ContentArticleList{
		Year:        "2024",
		ArticleList: articles,
	}
}
func (s *service) GetHotArticle(ctx *gin.Context) (*response.PageListResponse, error) {

	return &response.PageListResponse{}, nil
}
func (s *service) GetArticleById(ctx *gin.Context, articleId int) (*DetailArticle, error) {
	db := GetDB(ctx)
	article, err := articleDao.GetArticleById(db, articleId)
	if err != nil {
		return nil, err
	}
	var resp DetailArticle
	resp.Article = article
	// 获取tag
	ids, err := tagDao.GetTagIdsByArticleId(db, articleId)
	tags, err := tagDao.GetTagsById(db, ids)
	resp.TagIdList = ids
	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.TagName)
	}
	resp.TagNameList = tagNames
	return &resp, nil
}
