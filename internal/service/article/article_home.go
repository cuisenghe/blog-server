package article

import (
	"blog-server/internal/repository/articleDao"
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
type ArticleListResp struct {
	Current int                   `json:"current"`
	Size    int                   `json:"size"`
	List    []*articleDao.Article `json:"list"`
	Total   int64                 `json:"total"`
}
type SimpleArticleListResp struct {
	Current int                  `json:"current"`
	Size    int                  `json:"size"`
	List    []*SimpleArticleList `json:"list"`
	Total   int64                `json:"total"`
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
type RecommendArticleListResp struct {
	Current int                   `json:"current"`
	Size    int                   `json:"size"`
	List    *RecommendArticleList `json:"list"`
	Total   int64                 `json:"total"`
}
type RecommendArticleList struct {
	Year        string           `json:"year"`
	ArticleList []*SimpleArticle `json:"articleList"`
}
type ContentArticleListResp struct {
	Current int                 `json:"current"`
	Size    int                 `json:"size"`
	List    *ContentArticleList `json:"list"`
	Total   int64               `json:"total"`
}
type ContentArticleList struct {
	Year        string           `json:"year"`
	ArticleList []*SimpleArticle `json:"articleList"`
}

func (s *service) GetArticleList(ctx *gin.Context, req *ArticleListData) (*ArticleListResp, error) {
	if req.Size == 0 {
		return &ArticleListResp{
			Current: 0,
			Size:    0,
			List:    nil,
			Total:   0,
		}, nil
	}
	list, err := articleDao.GetArticleList(ctx, req.Current, req.Size)
	if err != nil {
		return nil, err
	}
	count, err := articleDao.GetSumCount(ctx)
	if err != nil {
		return nil, err
	}

	return &ArticleListResp{
		Current: req.Current,
		Size:    req.Size,
		List:    list,
		Total:   count,
	}, nil
}

// BlogTimelineGetArticleList 获取时间线文章
func (s *service) BlogTimelineGetArticleList(ctx *gin.Context, data *ArticleListData) (*SimpleArticleListResp, error) {
	// 获取文章
	list, err := articleDao.GetArticleList(ctx, data.Current, data.Size)
	if err != nil {
		return nil, err
	}
	count, err := articleDao.GetSumCount(ctx)
	if err != nil {
		return nil, err
	}
	resp := convertTimeLineData(list)
	resp.Size = data.Size
	resp.Total = count
	resp.Current = data.Current
	return resp, nil
}
func convertTimeLineData(list []*articleDao.Article) *SimpleArticleListResp {
	// 分组
	var resp SimpleArticleListResp
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
	resp.List = simpleList
	// return
	return &resp

}
func (s *service) GetArticleListByTagId(ctx *gin.Context, req *ArticleListData) (*SimpleArticleListResp, error) {
	// 根据条件查询文章
	condition, err := articleDao.GetArticleListByCondition(ctx, req.Current, req.Size, map[string]interface{}{
		"tag_id": req.Id,
	})
	if err != nil {
		return nil, err
	}
	return convertTimeLineData(condition), nil
}
func (s *service) GetArticleListByCategoryId(ctx *gin.Context, req *ArticleListData) (*SimpleArticleListResp, error) {
	condition, err := articleDao.GetArticleListByCondition(ctx, req.Current, req.Size, map[string]interface{}{
		"category_id": req.Id,
	})
	if err != nil {
		return nil, err
	}
	return convertTimeLineData(condition), nil
}

// 获取推荐
func (s *service) GetRecommendArticleById(ctx *gin.Context, articleId int) (*RecommendArticleListResp, error) {
	// 获取推荐
	return &RecommendArticleListResp{}, nil
}

// 根据内容搜索
func (s *service) GetArticleListByContent(ctx *gin.Context, content string) (*ContentArticleListResp, error) {
	articleList, err := articleDao.GetArticleListByContent(ctx, content)
	if err != nil {
		return &ContentArticleListResp{
			Current: 0,
			Size:    0,
			List:    nil,
			Total:   0,
		}, err
	}
	// 获取count
	count, err := articleDao.GetArticleCountByContent(ctx, content)
	return &ContentArticleListResp{
		Current: 0,
		Size:    0,
		List:    convertSearchData(articleList),
		Total:   count,
	}, nil
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
func (s *service) GetHotArticle(ctx *gin.Context) (*SimpleArticleListResp, error) {

	return &SimpleArticleListResp{}, nil
}
