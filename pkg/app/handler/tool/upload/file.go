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

type File struct {
	toolupload.Template
}

// 初始化
func (p *File) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 限制文件大小
	p.LimitSize = 1024 * 1024 * 1024 * 2

	// 限制文件类型
	p.LimitType = []string{
		"image/png",
		"image/gif",
		"image/jpeg",
		"video/mp4",
		"video/mpeg",
		"application/x-xls",
		"application/x-ppt",
		"application/msword",
		"application/zip",
		"application/pdf",
		"application/vnd.ms-excel",
		"application/vnd.ms-powerpoint",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/vnd.openxmlformats-officedocument.presentationml.presentation",
	}

	// 设置文件上传路径
	p.SavePath = "./website/storage/files/" + time.Now().Format("20060102") + "/"

	return p
}

// 上传前回调
func (p *File) BeforeHandle(ctx *builder.Context, fileSystem *storage.FileSystem) (*storage.FileSystem, *storage.FileInfo, error) {
	fileHash, err := fileSystem.GetFileHash()
	if err != nil {
		return fileSystem, nil, err
	}

	getFileInfo, _ := (&model.File{}).GetInfoByHash(fileHash)
	if err != nil {
		return fileSystem, nil, err
	}
	if getFileInfo.Id != 0 {
		fileInfo := &storage.FileInfo{
			Name: getFileInfo.Name,
			Size: getFileInfo.Size,
			Ext:  getFileInfo.Ext,
			Path: getFileInfo.Path,
			Url:  getFileInfo.Url,
			Hash: getFileInfo.Hash,
		}

		return fileSystem, fileInfo, err
	}

	return fileSystem, nil, err
}

// 上传完成后回调
func (p *File) AfterHandle(ctx *builder.Context, result *storage.FileInfo) interface{} {
	driver := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Driver").String()

	// 重写url
	if driver == storage.LocalDriver {
		result.Url = (&model.File{}).GetPath(result.Url)
	}

	// 插入数据库
	id, err := (&model.File{}).InsertGetId(&model.File{
		Name:   result.Name,
		Size:   result.Size,
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
		"name":        result.Name,
		"path":        result.Path,
		"size":        result.Size,
		"url":         result.Url,
	}))
}
