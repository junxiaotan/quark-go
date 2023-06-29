package resources

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
)

type File struct {
	adminresource.Template
}

// 初始化
func (p *File) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "文件"

	// 模型
	p.Model = &model.File{}

	// 分页
	p.PerPage = 10

	return p
}

// 字段
func (p *File) Fields(ctx *builder.Context) []interface{} {
	field := &adminresource.Field{}

	return []interface{}{
		field.ID("id", "ID"),
		field.Text("name", "名称"),
		field.Text("size", "大小").
			SetColumn(func(column *table.Column) *table.Column {
				return column.SetSorter(true)
			}),
		field.Text("ext", "扩展名"),
		field.Datetime("created_at", "上传时间", func() interface{} {
			if p.Field["created_at"] == nil {
				return p.Field["created_at"]
			}

			return p.Field["created_at"].(time.Time).Format("2006-01-02 15:04:05")
		}),
	}
}

// 搜索
func (p *File) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{
		(&searches.Input{}).Init("name", "名称"),
		(&searches.DateTimeRange{}).Init("created_at", "上传时间"),
	}
}

// 行为
func (p *File) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		(&actions.Delete{}).Init("批量删除"),
		(&actions.Delete{}).Init("删除"),
	}
}
