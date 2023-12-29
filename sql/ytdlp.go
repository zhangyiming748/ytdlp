package sql

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Ytdlp struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"status"`
	URL       string `gorm:"url,type=string"`
	Status    string `gorm:"status"`
	ErrorMsg  string `gorm:"error_msg"`
	ErrorCode string `gorm:"error_code"`
	Request   string `gorm:"request"`
	CreatedAt time.Time
}

func (y *Ytdlp) FindOneByURL() *gorm.DB {
	return GetEngine().First(&y, "url = ?", y.URL)
}
func (y *Ytdlp) SetOne() *gorm.DB {
	return GetEngine().Create(&y)
}
func (y *Ytdlp) UpdateStatusById() *gorm.DB {
	return GetEngine().Model(Ytdlp{}).Where("id = ?", y.ID).Update("status", y.Status)
}
func (y *Ytdlp) FindDupByName() (Success, Failure, Skip []Ytdlp) {
	var ys []Ytdlp
	GetEngine().Table("ytdlps").Where("name = ?", y.Name).Find(&ys)
	fmt.Println(ys)
	for _, v := range ys {
		{
			switch v.Status {
			case "下载失败":
				Failure = append(Failure, v)
			case "跳过下载":
				Skip = append(Skip, v)
			case "下载成功":
				Success = append(Success, v)
			}
		}
	}
	return Success, Failure, Skip
}
