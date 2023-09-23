package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pss/api/controller"
	"pss/api/controller/user"
	"pss/api/controller/v1"
	"pss/api/controller/v2"
	"pss/api/middleware/jwt"
	"pss/docs"
	"pss/pkg/export"
	"pss/pkg/qrcode"
	"pss/pkg/upload"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//login
	r.GET("/auth", controller.GetAuth)

	//上传图片
	r.POST("/upload", controller.UploadImage)

	article := r.Group("/article")
	pet := r.Group("/pet")
	user1 := r.Group("/user")

	article.Use(jwt.JWT())
	{
		//获取标签列表
		article.GET("/tags", v1.GetTags)
		//新建标签
		article.POST("/tags", v1.AddTag)
		//更新指定标签
		article.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		article.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)
		//获取文章列表
		article.GET("/articles", v1.GetArticles)
		//获取指定文章
		article.GET("/articles/:id", v1.GetArticle)
		//新建文章
		article.POST("/articles", v1.AddArticle)
		//更新指定文章
		article.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		article.DELETE("/articles/:id", v1.DeleteArticle)
		//生成二维码
		article.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	pet.Use()
	{
		//获取宠物
		pet.GET("/pets", v2.GetRecPets)
		//获取推荐宠物
		pet.GET("/rec/pets", v1.GetTags)

	}

	user1.Use()
	{
		user1.Group("login", user.Login)
	}

	return r
}
