package router

import (
	"blog-server/datasbase/mysql"
	"blog-server/datasbase/redis"
	"blog-server/internal/api/article"
	"blog-server/internal/api/blogConfig"
	"blog-server/internal/api/category"
	"blog-server/internal/api/chat"
	"blog-server/internal/api/comment"
	"blog-server/internal/api/like"
	"blog-server/internal/api/pageHeader"
	"blog-server/internal/api/statistic"
	"blog-server/internal/api/tag"
	"blog-server/internal/api/user"
	"github.com/gin-gonic/gin"
)

// InitRouter router 分组，先写用户模块
func InitRouter() *gin.Engine {
	// 获取router engine
	r := gin.Default()
	initMiddleware(r)
	initHandler(r)
	InitInnerRouter(r)
	return r
}
func initMiddleware(r *gin.Engine) {
	db := mysql.InitMySQL()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	rdb := redis.InitRedis()
	r.Use(func(c *gin.Context) {
		c.Set("rdb", rdb)
		c.Next()
	})
}
func initHandler(r *gin.Engine) {
	// 用户模块
	userGroup := r.Group("/user")
	{
		handler := user.NewHandler()
		// 用户注册
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
		userGroup.GET("/getUserInfoById/:id", handler.GetUserInfoById)
		userGroup.PUT("/updateOwnUserInfo", handler.UpdateOwnUserInfo)
		userGroup.PUT("/updatePassword", handler.UpdatePassword)
		//userGroup.PUT("/updateRole/:")
	}
	// 文章模块
	articleGroup := r.Group("/article")
	{
		handler := article.NewHandler()
		// 前台
		articleGroup.GET("/blogHomeGetArticleList/:current/:size", handler.GetArticleList)
		articleGroup.GET("/blogTimelineGetArticleList/:current/:size", handler.BlogTimelineGetArticleList)
		articleGroup.POST("/getArticleListByTagId", handler.GetArticleListByTagId)
		articleGroup.GET("/getRecommendArticleById/:id", handler.GetRecommendArticleById)
		articleGroup.GET("/getArticleListByContent/:content", handler.GetArticleListByContent)
		articleGroup.GET("/getHotArticle", handler.GetHotArticle)
		articleGroup.GET("/getArticleById/:id", handler.GetArticleById)
		// 后台
		articleGroup.POST("/add", handler.AddArticle)
		articleGroup.POST("/update", handler.UpdateArticle)
		articleGroup.DELETE("/delete/:id/:status", handler.DeleteArticle)
		articleGroup.PUT("/revert/:id", handler.RevertArticle)
		articleGroup.POST("/titleExist", handler.TitleExist)
		articleGroup.PUT("/isPublic/:id/:status", handler.IsPublic)
		articleGroup.POST("/updateTop/:id/:is_top", handler.UpdateTop)
		articleGroup.POST("/getArticleList", handler.AdminGetArticleList)
	}
	configGroup := r.Group("/config")
	{
		handler := blogConfig.NewHandler()
		configGroup.GET("", handler.GetConfig)
		configGroup.PUT("/addView", handler.AddView)
	}
	chatGroup := r.Group("/chat")
	{
		handler := chat.NewHandler()
		chatGroup.POST("/getChatList", handler.GetChatList)
	}
	pageHeadGroup := r.Group("/pageHeader")
	{
		handler := pageHeader.NewHandler()
		pageHeadGroup.GET("/getAll", handler.GetAll)
	}
	statisticGroup := r.Group("/statistic")
	{
		handler := statistic.NewHandler()
		statisticGroup.GET("", handler.GetStatistic)
	}
	tagGroup := r.Group("/tag")
	{
		handler := tag.NewHandler()
		tagGroup.GET("/getTagDictionary", handler.GetTagDictionary)
		tagGroup.POST("/getTagList", handler.GetTagList)
		tagGroup.POST("/add", handler.AddTag)
	}
	likeGroup := r.Group("/like")
	{
		handler := like.NewHandler()
		likeGroup.POST("/getIsLikeByIdOrIpAndType", handler.GetIsLikeByIdAndType)
		likeGroup.POST("/addLike", handler.AddLike)
		likeGroup.POST("cancelLike", handler.CancelLike)
	}
	commentGroup := r.Group("/comment")
	{
		handler := comment.NewHandler()
		commentGroup.POST("getCommentTotal", handler.GetCommentTotal)
	}
	categoryGroup := r.Group("/category")
	{
		handler := category.NewHandler()
		categoryGroup.GET("/getCategoryDictionary", handler.GetCategoryDictionary)
		categoryGroup.POST("/getCategoryList", handler.GetCategoryList)
		categoryGroup.POST("/add", handler.AddCategory)
	}
}
