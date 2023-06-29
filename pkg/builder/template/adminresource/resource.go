package adminresource

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/requests"
	"github.com/quarkcms/quark-go/pkg/component/admin/footer"
	"github.com/quarkcms/quark-go/pkg/component/admin/layout"
	"github.com/quarkcms/quark-go/pkg/component/admin/page"
	"github.com/quarkcms/quark-go/pkg/component/admin/pagecontainer"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
)

// 后台增删改查模板
type Template struct {
	template.Template
	Title        string // 标题
	SubTitle     string // 子标题
	PerPage      interface{}
	IndexPolling int
	IndexOrder   string
	Model        interface{}
	Field        map[string]interface{}
	WithExport   bool
}

const (
	IndexRoute          = "/api/admin/:resource/index"           // 后台增删改查，列表路由
	EditableRoute       = "/api/admin/:resource/editable"        // 后台增删改查，表格行内编辑路由
	ActionRoute         = "/api/admin/:resource/action/:uriKey"  // 后台增删改查，执行行为路由
	CreateRoute         = "/api/admin/:resource/create"          // 后台增删改查，创建页面路由
	StoreRoute          = "/api/admin/:resource/store"           // 后台增删改查，创建方法路由
	EditRoute           = "/api/admin/:resource/edit"            // 后台增删改查，编辑页面路由
	EditValuesRoute     = "/api/admin/:resource/edit/values"     // 后台增删改查，获取编辑表单值路由
	SaveRoute           = "/api/admin/:resource/save"            // 后台增删改查，保存编辑值路由
	ImportRoute         = "/api/admin/:resource/import"          // 后台增删改查，详情页面路由
	ExportRoute         = "/api/admin/:resource/export"          // 后台增删改查，导出数据路由
	DetailRoute         = "/api/admin/:resource/detail"          // 后台增删改查，导入数据路由
	ImportTemplateRoute = "/api/admin/:resource/import/template" // 后台增删改查，导入模板路由
	FormRoute           = "/api/admin/:resource/:uriKey/form"    // 后台增删改查，通用表单资源
)

// 初始化
func (p *Template) Init() interface{} {
	p.TemplateInit()

	return p
}

// 初始化模板
func (p *Template) TemplateInit() interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 注册路由映射
	p.GET(IndexRoute, "IndexRender")                   // 后台增删改查，列表路由
	p.GET(EditableRoute, "EditableRender")             // 后台增删改查，表格行内编辑路由
	p.Any(ActionRoute, "ActionRender")                 // 后台增删改查，执行行为路由
	p.GET(CreateRoute, "CreationRender")               // 后台增删改查，创建页面路由
	p.POST(StoreRoute, "StoreRender")                  // 后台增删改查，创建方法路由
	p.GET(EditRoute, "EditRender")                     // 后台增删改查，编辑页面路由
	p.GET(EditValuesRoute, "EditValuesRender")         // 后台增删改查，获取编辑表单值路由
	p.POST(SaveRoute, "SaveRender")                    // 后台增删改查，保存编辑值路由
	p.GET(DetailRoute, "DetailRender")                 // 后台增删改查，详情页面路由
	p.GET(ExportRoute, "ExportRender")                 // 后台增删改查，导出数据路由
	p.POST(ImportRoute, "ImportRender")                // 后台增删改查，导入数据路由
	p.GET(ImportTemplateRoute, "ImportTemplateRender") // 后台增删改查，导入模板路由
	p.GET(FormRoute, "FormRender")                     // 后台增删改查，通用表单资源

	return p
}

// 设置单列字段
func (p *Template) SetField(fieldData map[string]interface{}) interface{} {
	p.Field = fieldData

	return p
}

// 数据导出前回调
func (p *Template) BeforeExporting(ctx *builder.Context, list []map[string]interface{}) []interface{} {
	result := []interface{}{}
	for _, v := range list {
		result = append(result, v)
	}

	return result
}

// 数据导入前回调
func (p *Template) BeforeImporting(ctx *builder.Context, list [][]interface{}) [][]interface{} {
	return list
}

// 列表页渲染
func (p *Template) IndexRender(ctx *builder.Context) interface{} {
	data := (&requests.IndexRequest{}).QueryData(ctx)
	body := p.IndexComponentRender(ctx, data)

	result := ctx.Template.(interface {
		PageComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 表格行内编辑
func (p *Template) EditableRender(ctx *builder.Context) interface{} {
	return (&requests.EditableRequest{}).Handle(ctx)
}

// 执行行为
func (p *Template) ActionRender(ctx *builder.Context) interface{} {
	return (&requests.ActionRequest{}).Handle(ctx)
}

// 创建页面渲染
func (p *Template) CreationRender(ctx *builder.Context) interface{} {

	// 断言BeforeCreating方法，获取初始数据
	data := ctx.Template.(interface {
		BeforeCreating(ctx *builder.Context) map[string]interface{}
	}).BeforeCreating(ctx)

	// 断言CreationComponentRender方法
	body := ctx.Template.(interface {
		CreationComponentRender(*builder.Context, map[string]interface{}) interface{}
	}).CreationComponentRender(ctx, data)

	result := ctx.Template.(interface {
		PageComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 创建方法
func (p *Template) StoreRender(ctx *builder.Context) interface{} {
	return (&requests.StoreRequest{}).Handle(ctx)
}

// 编辑页面渲染
func (p *Template) EditRender(ctx *builder.Context) interface{} {
	// 获取数据
	data := (&requests.EditRequest{}).FillData(ctx)

	// 断言BeforeEditing方法，获取初始数据
	data = ctx.Template.(interface {
		BeforeEditing(*builder.Context, map[string]interface{}) map[string]interface{}
	}).BeforeEditing(ctx, data)

	// 断言UpdateComponentRender方法
	body := ctx.Template.(interface {
		UpdateComponentRender(*builder.Context, map[string]interface{}) interface{}
	}).UpdateComponentRender(ctx, data)

	result := ctx.Template.(interface {
		PageComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 获取编辑表单值
func (p *Template) EditValuesRender(ctx *builder.Context) interface{} {

	return (&requests.EditRequest{}).Values(ctx)
}

// 保存编辑值
func (p *Template) SaveRender(ctx *builder.Context) interface{} {

	return (&requests.UpdateRequest{}).Handle(ctx)
}

// 详情页渲染
func (p *Template) DetailRender(ctx *builder.Context) interface{} {
	data := (&requests.DetailRequest{}).FillData(ctx)

	// 断言方法，获取初始数据
	data = ctx.Template.(interface {
		BeforeDetailShowing(*builder.Context, map[string]interface{}) map[string]interface{}
	}).BeforeDetailShowing(ctx, data)

	// 断言方法
	body := ctx.Template.(interface {
		DetailComponentRender(*builder.Context, map[string]interface{}) interface{}
	}).DetailComponentRender(ctx, data)

	result := ctx.Template.(interface {
		PageComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 导出数据
func (p *Template) ExportRender(ctx *builder.Context) interface{} {

	return (&requests.ExportRequest{}).Handle(ctx)
}

// 导入数据
func (p *Template) ImportRender(ctx *builder.Context) interface{} {

	return (&requests.ImportRequest{}).Handle(ctx, IndexRoute)
}

// 导入数据模板
func (p *Template) ImportTemplateRender(ctx *builder.Context) interface{} {

	return (&requests.ImportTemplateRequest{}).Handle(ctx)
}

// 通用表单资源
func (p *Template) FormRender(ctx *builder.Context) interface{} {

	// 断言BeforeCreating方法，获取初始数据
	data := ctx.Template.(interface {
		BeforeCreating(ctx *builder.Context) map[string]interface{}
	}).BeforeCreating(ctx)

	// 断言CreationComponentRender方法
	body := ctx.Template.(interface {
		CreationComponentRender(*builder.Context, map[string]interface{}) interface{}
	}).CreationComponentRender(ctx, data)

	result := ctx.Template.(interface {
		PageComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 页面组件渲染
func (p *Template) PageComponentRender(ctx *builder.Context, body interface{}) interface{} {

	// Layout组件
	layoutComponent := ctx.Template.(interface {
		LayoutComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).LayoutComponentRender(ctx, body)

	return (&page.Component{}).
		Init().
		SetStyle(map[string]interface{}{
			"height": "100vh",
		}).
		SetBody(layoutComponent).
		JsonSerialize()
}

// 页面布局组件渲染
func (p *Template) LayoutComponentRender(ctx *builder.Context, body interface{}) interface{} {
	admin := &model.Admin{}
	config := ctx.Engine.GetConfig()

	// 获取登录管理员信息
	adminInfo, err := admin.GetAuthUser(config.AppKey, ctx.Token())
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	// 获取管理员菜单
	getMenus, err := admin.GetMenuListById(adminInfo.Id)
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	adminLayout := ctx.Engine.GetAdminLayout()

	// 页脚
	footer := (&footer.Component{}).
		Init().
		SetCopyright(adminLayout.Copyright).
		SetLinks(adminLayout.Links)

	// 页面容器组件渲染
	pageContainerComponent := ctx.Template.(interface {
		PageContainerComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageContainerComponentRender(ctx, body)

	return (&layout.Component{}).
		Init().
		SetTitle(adminLayout.Title).
		SetLogo(adminLayout.Logo).
		SetActions(adminLayout.Actions).
		SetLayout(adminLayout.Layout).
		SetSplitMenus(adminLayout.SplitMenus).
		SetContentWidth(adminLayout.ContentWidth).
		SetPrimaryColor(adminLayout.PrimaryColor).
		SetFixSiderbar(adminLayout.FixSiderbar).
		SetFixedHeader(adminLayout.FixedHeader).
		SetIconfontUrl(adminLayout.IconfontUrl).
		SetLocale(adminLayout.Locale).
		SetSiderWidth(adminLayout.SiderWidth).
		SetMenu(getMenus).
		SetBody(pageContainerComponent).
		SetFooter(footer)
}

// 页面容器组件渲染
func (p *Template) PageContainerComponentRender(ctx *builder.Context, body interface{}) interface{} {
	value := reflect.ValueOf(ctx.Template).Elem()
	title := value.FieldByName("Title").String()
	subTitle := value.FieldByName("SubTitle").String()

	// 设置头部
	header := (&pagecontainer.PageHeader{}).
		Init().
		SetTitle(title).
		SetSubTitle(subTitle)

	return (&pagecontainer.Component{}).Init().SetHeader(header).SetBody(body)
}
