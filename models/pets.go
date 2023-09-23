package models

import (
	"github.com/jinzhu/gorm"
	"pss/pkg/setting"
)

type Pet struct {
	Model
	PetId      string `json:"pet_id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
	Owner      string `json:"owner"`
	Category   string `json:"category"` // 种类  猫 狗 ......
	Pic        string `json:"pic"`
}

// GetTags gets a list of tags based on paging and constraints
func GetPets(pageNum int, pageSize int, maps interface{}) ([]Pet, error) {
	var (
		pets []Pet
		err  error
	)

	if pageSize > 0 && pageNum > 0 {
		err = setting.MysqlClient.Find(&pets).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = setting.MysqlClient.Find(&pets).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return pets, nil
}

func GetPetTotal(maps interface{}) (count int) {
	setting.MysqlClient.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func GetPetByCategory(maps interface{}) (count int) {
	setting.MysqlClient.Model(&Pet{}).Where(maps)
	return
}
