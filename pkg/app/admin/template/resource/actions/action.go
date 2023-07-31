package actions

import (
	"reflect"
	"strings"

	"github.com/gobeam/stringy"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Action struct {
	Name                  string      `json:"name"`
	Reload                string      `json:"reload"`
	ApiParams             []string    `json:"apiParams"`
	Api                   string      `json:"api"`
	ActionType            string      `json:"actionType"`
	SubmitForm            string      `json:"submitForm"`
	Icon                  string      `json:"icon"`
	Type                  string      `json:"type"`
	Size                  string      `json:"size"`
	WithLoading           bool        `json:"withLoading"`
	Fields                interface{} `json:"fields"`
	ConfirmTitle          string      `json:"confirmTitle"`
	ConfirmText           string      `json:"confirmText"`
	ConfirmType           string      `json:"confirmType"`
	OnlyOnIndex           bool        `json:"onlyOnIndex"`
	OnlyOnForm            bool        `json:"onlyOnForm"`
	OnlyOnDetail          bool        `json:"onlyOnDetail"`
	ShowOnIndex           bool        `json:"showOnIndex"`
	ShowOnIndexTableRow   bool        `json:"showOnIndexTableRow"`
	ShowOnIndexTableAlert bool        `json:"showOnIndexTableAlert"`
	ShowOnForm            bool        `json:"showOnForm"`
	ShowOnFormExtra       bool        `json:"showOnFormExtra"`
	ShowOnDetail          bool        `json:"showOnDetail"`
	ShowOnDetailExtra     bool        `json:"showOnDetailExtra"`
}

// 初始化
func (p *Action) Init(ctx *builder.Context) interface{} {
	return p
}

// 初始化模板
func (p *Action) TemplateInit(ctx *builder.Context) interface{} {
	p.ActionType = "ajax"

	return p
}

// 行为key
func (p *Action) GetUriKey(action interface{}) string {
	uriKey := reflect.TypeOf(action).String()
	uriKeys := strings.Split(uriKey, ".")
	uriKey = stringy.New(uriKeys[1]).KebabCase("?", "").ToLower()

	return uriKey
}

// 获取名称
func (p *Action) GetName() string {
	return p.Name
}

// 执行成功后刷新的组件
func (p *Action) GetReload() string {
	return p.Reload
}

// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
func (p *Action) GetApiParams() []string {
	return p.ApiParams
}

// 执行行为的接口
func (p *Action) GetApi() string {
	return p.Api
}

// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
func (p *Action) GetActionType() string {
	return p.ActionType
}

// 当 action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
func (p *Action) GetSubmitForm() string {
	return p.SubmitForm
}

// 设置按钮类型，primary | ghost | dashed | link | text | default
func (p *Action) GetType() string {
	return p.Type
}

// 设置按钮大小,large | middle | small | default
func (p *Action) GetSize() string {
	return p.Size
}

// 是否具有loading，当action 的作用类型为ajax,submit时有效
func (p *Action) GetWithLoading() bool {
	return p.WithLoading
}

// 设置按钮的图标组件
func (p *Action) GetIcon() string {
	return p.Icon
}

// 行为表单字段
func (p *Action) GetFields() interface{} {
	return p.Fields
}

// 确认标题
func (p *Action) GetConfirmTitle() string {
	return p.ConfirmTitle
}

// 确认文字
func (p *Action) GetConfirmText() string {
	return p.ConfirmText
}

// 确认类型
func (p *Action) GetConfirmType() string {
	return p.ConfirmType
}

// 设置名称
func (p *Action) SetName(name string) *Action {
	p.Name = name

	return p
}

// 设置执行成功后刷新的组件
func (p *Action) SetReload(componentKey string) *Action {
	p.Reload = componentKey

	return p
}

// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
func (p *Action) SetApiParams(apiParams []string) *Action {
	p.ApiParams = apiParams

	return p
}

// 执行行为的接口
func (p *Action) SetApi(api string) *Action {
	p.Api = api

	return p
}

// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
func (p *Action) SetActionType(actionType string) *Action {
	p.ActionType = actionType

	return p
}

// 当 action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
func (p *Action) SetSubmitForm(submitForm string) *Action {
	p.SubmitForm = submitForm

	return p
}

// 设置按钮类型，primary | ghost | dashed | link | text | default
func (p *Action) SetType(buttonType string) *Action {
	p.Type = buttonType

	return p
}

// 设置按钮大小,large | middle | small | default
func (p *Action) SetSize(size string) *Action {
	p.Size = size

	return p
}

// 是否具有loading，当action 的作用类型为ajax,submit时有效
func (p *Action) SetWithLoading(loading bool) *Action {
	p.WithLoading = loading

	return p
}

// 设置按钮的图标组件
func (p *Action) SetIcon(icon string) *Action {
	p.Icon = icon

	return p
}

// 行为表单字段
func (p *Action) SetFields(fields interface{}) *Action {
	p.Fields = fields

	return p
}

// 确认标题
func (p *Action) SetConfirmTitle(confirmTitle string) *Action {
	p.ConfirmTitle = confirmTitle

	return p
}

// 确认文字
func (p *Action) SetConfirmText(confirmText string) *Action {
	p.ConfirmText = confirmText

	return p
}

// 确认类型
func (p *Action) SetConfirmType(confirmType string) *Action {
	p.ConfirmType = confirmType

	return p
}

// 设置行为前的确认操作
func (p *Action) WithConfirm(title string, text string, confirmType string) *Action {

	p.ConfirmTitle = title
	p.ConfirmText = text
	p.ConfirmType = confirmType

	return p
}

// 只在列表页展示
func (p *Action) SetOnlyOnIndex(value bool) *Action {
	p.OnlyOnIndex = value
	p.ShowOnIndex = value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value

	return p
}

// 除了列表页外展示
func (p *Action) SetExceptOnIndex() *Action {
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnIndexTableAlert = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true
	p.ShowOnIndex = false

	return p
}

// 只在表单页展示
func (p *Action) SetOnlyOnForm(value bool) *Action {
	p.ShowOnForm = value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value

	return p
}

// 除了表单页外展示
func (p *Action) SetExceptOnForm() *Action {
	p.ShowOnIndexTableAlert = true
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = false
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true

	return p
}

// 只在表单页右上角自定义区域展示
func (p *Action) SetOnlyOnFormExtra(value bool) *Action {
	p.ShowOnForm = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnFormExtra = value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value

	return p
}

// 除了表单页右上角自定义区域外展示
func (p *Action) SetExceptOnFormExtra() *Action {
	p.ShowOnIndexTableAlert = true
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = false
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true

	return p
}

// 只在详情页展示
func (p *Action) SetOnlyOnDetail(value bool) *Action {
	p.OnlyOnDetail = value
	p.ShowOnDetail = value
	p.ShowOnIndex = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetailExtra = !value

	return p
}

// 除了详情页外展示
func (p *Action) SetExceptOnDetail() *Action {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnIndexTableRow = true
	p.ShowOnIndexTableAlert = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetailExtra = true

	return p
}

// 只在详情页右上角自定义区域展示
func (p *Action) SetOnlyOnDetailExtra(value bool) *Action {
	p.ShowOnForm = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = value

	return p
}

// 除了详情页右上角自定义区域外展示
func (p *Action) SetExceptOnDetailExtra() *Action {
	p.ShowOnIndexTableAlert = true
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = false

	return p
}

// 在表格行内展示
func (p *Action) SetOnlyOnIndexTableRow(value bool) *Action {
	p.ShowOnIndexTableRow = value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value

	return p
}

// 除了表格行内外展示
func (p *Action) SetExceptOnIndexTableRow() *Action {
	p.ShowOnIndexTableRow = false
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableAlert = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true

	return p
}

// 在表格多选弹出层展示
func (p *Action) SetOnlyOnIndexTableAlert(value bool) *Action {
	p.ShowOnIndexTableAlert = value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value

	return p
}

// 除了表格多选弹出层外展示
func (p *Action) SetExceptOnIndexTableAlert() *Action {
	p.ShowOnIndexTableAlert = false
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true

	return p
}

// 在列表页展示
func (p *Action) SetShowOnIndex() *Action {
	p.ShowOnIndex = true

	return p
}

// 在表单页展示
func (p *Action) SetShowOnForm() *Action {
	p.ShowOnForm = true

	return p
}

// 在表单页右上角自定义区域展示
func (p *Action) SetShowOnFormExtra() *Action {
	p.ShowOnFormExtra = true

	return p
}

// 在详情页展示
func (p *Action) SetShowOnDetail() *Action {
	p.ShowOnDetail = true

	return p
}

// 在详情页右上角自定义区域展示
func (p *Action) SetShowOnDetailExtra() *Action {
	p.ShowOnDetailExtra = true

	return p
}

// 在表格行内展示
func (p *Action) SetShowOnIndexTableRow() *Action {
	p.ShowOnIndexTableRow = true

	return p
}

// 在多选弹出层展示
func (p *Action) SetShowOnIndexTableAlert() *Action {
	p.ShowOnIndexTableAlert = true

	return p
}

// 判断是否在列表页展示
func (p *Action) ShownOnIndex() bool {
	if p.OnlyOnIndex {
		return true
	}

	if p.OnlyOnDetail {
		return false
	}

	if p.OnlyOnForm {
		return false
	}

	return p.ShowOnIndex
}

// 判断是否在表单页展示
func (p *Action) ShownOnForm() bool {
	if p.OnlyOnForm {
		return true
	}

	if p.OnlyOnDetail {
		return false
	}

	if p.OnlyOnIndex {
		return false
	}

	return p.ShowOnForm
}

// 判断是否在详情页展示
func (p *Action) ShownOnDetail() bool {
	if p.OnlyOnDetail {
		return true
	}

	if p.OnlyOnIndex {
		return false
	}

	if p.OnlyOnForm {
		return false
	}

	return p.ShowOnDetail
}

// 判断是否在表格行内展示
func (p *Action) ShownOnIndexTableRow() bool {
	return p.ShowOnIndexTableRow
}

// 判断是否在多选弹出层展示
func (p *Action) ShownOnIndexTableAlert() bool {
	return p.ShowOnIndexTableAlert
}

// 判断是否在表单页右上角自定义区域展示
func (p *Action) ShownOnFormExtra() bool {
	return p.ShowOnFormExtra
}

// 判断是否在详情页右上角自定义区域展示
func (p *Action) ShownOnDetailExtra() bool {
	return p.ShowOnDetailExtra
}