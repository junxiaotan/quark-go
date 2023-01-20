package model

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/pkg/dal/db"
)

// 字段
type Picture struct {
	Id                int       `json:"id" gorm:"autoIncrement"`
	ObjType           string    `json:"obj_type" gorm:"size:255"`
	ObjId             int       `json:"obj_id" gorm:"size:11;default:0"`
	PictureCategoryId int       `json:"picture_category_id" gorm:"size:11;default:0"`
	Sort              int       `json:"sort" gorm:"size:11;default:0"`
	Name              string    `json:"name" gorm:"size:255;not null"`
	Size              int64     `json:"size" gorm:"size:20;default:0"`
	Width             int       `json:"width" gorm:"size:11;default:0"`
	Height            int       `json:"height" gorm:"size:11;default:0"`
	Ext               string    `json:"ext" gorm:"size:255"`
	Path              string    `json:"path" gorm:"size:255;not null"`
	Url               string    `json:"url" gorm:"size:255;not null"`
	Hash              string    `json:"hash" gorm:"size:255;not null"`
	Status            int       `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// 插入数据并返回ID
func (model *Picture) InsertGetId(picture *Picture) (id int, Error error) {
	err := db.Client.Create(&picture).Error
	if err != nil {
		return id, err
	}

	return picture.Id, nil
}

// 根据hash查询文件信息
func (model *Picture) GetInfoByHash(hash string) (picture *Picture, Error error) {
	err := db.Client.Where("status = ?", 1).Where("hash = ?", hash).First(&picture).Error

	return picture, err
}

// 获取图片路径
func (model *Picture) GetPath(id interface{}) string {
	http, path := "", ""
	webSiteDomain := (&Config{}).GetValue("WEB_SITE_DOMAIN")
	WebConfig := (&Config{}).GetValue("SSL_OPEN")
	if webSiteDomain != "" {
		if WebConfig == "1" {
			http = "https://"
		} else {
			http = "http://"
		}
	}

	if getId, ok := id.(string); ok {
		if strings.Contains(getId, "//") && !strings.Contains(getId, "{") {
			return getId
		}
		if strings.Contains(getId, "./") && !strings.Contains(getId, "{") {
			return http + webSiteDomain + strings.Replace(getId, "./website", "./", -1)
		}
		if strings.Contains(getId, "/") && !strings.Contains(getId, "{") {
			return http + webSiteDomain + getId
		}

		// json字符串
		if strings.Contains(getId, "{") {
			var jsonData interface{}
			json.Unmarshal([]byte(getId), &jsonData)
			if mapData, ok := jsonData.(map[string]interface{}); ok {
				path = mapData["url"].(string)
			}

			// 如果为数组，返回第一个key的path
			if arrayData, ok := jsonData.([]map[string]interface{}); ok {
				path = arrayData[0]["url"].(string)
			}
		}
		if strings.Contains(path, "//") {
			return path
		}
		if strings.Contains(path, "./") {
			path = strings.Replace(path, "./website", "./", -1)
		}
		if path != "" {
			// 如果设置域名，则加上域名前缀
			return http + webSiteDomain + path
		}
	}

	picture := &Picture{}
	db.Client.Where("id", id).Where("status", 1).First(&picture)
	if picture.Id != 0 {
		path = picture.Path
		if strings.Contains(path, "//") {
			return path
		}
		if strings.Contains(path, "./") {
			path = strings.Replace(path, "./website", "./", -1)
		}
	}
	if path != "" {
		// 如果设置域名，则加上域名前缀
		return http + webSiteDomain + path
	}

	return http + webSiteDomain + "/admin/default.png"
}
