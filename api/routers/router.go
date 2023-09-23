package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"pss/api/controller"
	"pss/api/controller/article"
	"pss/api/controller/pet"
	"pss/api/controller/sample"
	"pss/docs"
	"pss/pkg/export"
	"pss/pkg/qrcode"
	"pss/pkg/upload"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", controller.GetAuth)
	r.POST("/upload", controller.UploadImage)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	sample1 := r.Group("/sample")
	article1 := r.Group("/article")
	pet1 := r.Group("/pet")
	user1 := r.Group("/user")
	video1 := r.Group("/video")
	market1 := r.Group("/market")
	mall1 := r.Group("/mall")
	friend1 := r.Group("/friend")

	sample1.Use()
	{
		sample1.GET("/db/insert", sample.DbInsert)
		sample1.GET("/db/update", sample.DbInsert)
		sample1.GET("/db/select", sample.DbInsert)
		sample1.GET("/db/delete", sample.DbInsert)
		sample1.GET("mongo/insert", sample.MongoInsert)
		sample1.GET("mongo/get", sample.MongoGet)
		sample1.GET("mongo/update", sample.MongoUpdate)
		sample1.GET("mongo/delete", sample.MongoDelete)
		sample1.GET("redis/5", sample.DbInsert)
		sample1.GET("redis/6", sample.DbInsert)
		sample1.GET("redis/7", sample.DbInsert)
		sample1.GET("search/8", sample.DbInsert)
		sample1.GET("search/9", sample.DbInsert)
		sample1.GET("sso/qiniu/upload", sample.QiniuUpload)
	}

	article1.Use()
	{
		article1.GET("/tags", article.GetTags)
		article1.POST("/tags", article.AddTag)
		article1.PUT("/tags/:id", article.EditTag)
		article1.DELETE("/tags/:id", article.DeleteTag)
		r.POST("/tags/export", article.ExportTag)
		r.POST("/tags/import", article.ImportTag)
		article1.GET("/articles", article.GetArticles)
		article1.GET("/articles/:id", article.GetArticle)
		article1.POST("/articles", article.AddArticle)
		article1.PUT("/articles/:id", article.EditArticle)
		article1.DELETE("/articles/:id", article.DeleteArticle)
		article1.POST("/articles/poster/generate", article.GenerateArticlePoster)
	}

	pet1.Use()
	{
		pet1.GET("/pets", pet.GetRecPets)
		pet1.GET("/rec/pets", article.GetTags)
		pet1.GET("/getTotal", pet.GetTotal)
	}

	user1.Use()
	{

	}

	video1.Use()
	{

	}

	market1.Use()
	{

	}

	mall1.Use()
	{

	}

	friend1.Use()
	{

	}

	return r
}
