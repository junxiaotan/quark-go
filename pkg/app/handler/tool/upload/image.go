package upload

import (
	"reflect"
	"time"

	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/toolupload"
	"github.com/quarkcms/quark-go/pkg/msg"
	"github.com/quarkcms/quark-go/pkg/storage"
)

type Image struct {
	toolupload.Template
}

// 初始化
func (p *Image) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 限制文件大小
	p.LimitSize = 1024 * 1024 * 1024 * 2

	// 限制文件类型
	p.LimitType = []string{
		"image/png",
		"image/gif",
		"image/jpeg",
	}

	// 设置文件上传路径
	p.SavePath = "./website/storage/images/" + time.Now().Format("20060102") + "/"

	return p
}

// 上传前回调
func (p *Image) BeforeHandle(ctx *builder.Context, fileSystem *storage.FileSystem) (*storage.FileSystem, *storage.FileInfo, error) {
	fileHash, err := fileSystem.GetFileHash()
	if err != nil {
		return fileSystem, nil, err
	}

	pictureInfo, _ := (&model.Picture{}).GetInfoByHash(fileHash)
	if err != nil {
		return fileSystem, nil, err
	}
	if pictureInfo.Id != 0 {
		fileInfo := &storage.FileInfo{
			Name:   pictureInfo.Name,
			Size:   pictureInfo.Size,
			Width:  pictureInfo.Width,
			Height: pictureInfo.Height,
			Ext:    pictureInfo.Ext,
			Path:   pictureInfo.Path,
			Url:    pictureInfo.Url,
			Hash:   pictureInfo.Hash,
		}

		return fileSystem, fileInfo, err
	}

	return fileSystem, nil, err
}

// 上传完成后回调
func (p *Image) AfterHandle(ctx *builder.Context, result *storage.FileInfo) interface{} {
	driver := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Driver").String()

	// 重写url
	if driver == storage.LocalDriver {
		result.Url = (&model.Picture{}).GetPath(result.Url)
	}

	// 插入数据库
	id, err := (&model.Picture{}).InsertGetId(&model.Picture{
		Name:   result.Name,
		Size:   result.Size,
		Width:  result.Width,
		Height: result.Height,
		Ext:    result.Ext,
		Path:   result.Path,
		Url:    result.Url,
		Hash:   result.Hash,
		Status: 1,
	})

	if err != nil {
		return ctx.JSON(200, msg.Error(err.Error(), ""))
	}

	return ctx.JSON(200, msg.Success("上传成功", "", map[string]interface{}{
		"id":          id,
		"contentType": result.ContentType,
		"ext":         result.Ext,
		"hash":        result.Hash,
		"height":      result.Height,
		"width":       result.Width,
		"name":        result.Name,
		"path":        result.Path,
		"size":        result.Size,
		"url":         result.Url,
	}))
}
