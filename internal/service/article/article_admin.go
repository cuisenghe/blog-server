package article

import (
	"blog-server/internal/repository/articleDao"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *service) DeleteArticle(ctx *gin.Context, id, status int) (bool, error) {
	article, err := articleDao.DeleteArticle(ctx, id, status)
	if err != nil {
		return false, err
	}
	return article, nil
}

// RevertArticle 恢复
func (s *service) RevertArticle(ctx *gin.Context, id int) (bool, error) {
	article, err := articleDao.RevertArticle(ctx, id)
	if err != nil {
		return false, err
	}
	return article, nil
}

// 判断是否存在
func (s *service) TitleExist(ctx *gin.Context, id int) (bool, error) {
	_, err := articleDao.GetArticleById(ctx, id)
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
	article, err := articleDao.UpdateArticleTop(ctx, id, is_top)
	if err != nil {
		return false, err
	}
	return article, nil
}
func (s *service) AdminGetArticleList(ctx *gin.Context, data *ArticleListData) (*ArticleListResp, error) {
	// 管理员获取文章列表
	list, err := articleDao.GetArticleList(ctx, data.Current, data.Size)
	if err != nil {
		return nil, err
	}
	count, err := articleDao.GetSumCount(ctx)
	if err != nil {
		return nil, err
	}
	return &ArticleListResp{
		Current: data.Current,
		Size:    data.Size,
		List:    list,
		Total:   count,
	}, nil
}
