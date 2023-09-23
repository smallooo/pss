package models

import (
	"pss/pkg/setting"
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

//	func ExistArticleByID(id int) bool {
//		var article Article
//		db.Select("id").Where("id = ?", id).First(&article)
//
//		if article.ID > 0 {
//			return true
//		}
//
//		return false
//	}
//
// ExistArticleByID checks if an article exists based on ID
func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := setting.MysqlClient.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetArticleTotal(maps interface{}) (count int) {
	setting.MysqlClient.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	setting.MysqlClient.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

//func GetArticle(id int) (article Article) {
//	db.Where("id = ?", id).First(&article)
//	db.Model(&article).Related(&article.Tag)
//
//	return
//}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := setting.MysqlClient.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

func EditArticle(id int, data interface{}) bool {
	setting.MysqlClient.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

// AddArticle add a single article
func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}
	if err := setting.MysqlClient.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

// DeleteArticle delete a single article
func DeleteArticle(id int) error {
	if err := setting.MysqlClient.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func CleanAllArticle() bool {
	setting.MysqlClient.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})

	return true
}
