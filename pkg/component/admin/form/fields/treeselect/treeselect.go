package treeselect

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/utils"
)

type FieldNames struct {
	Label    string `json:"title"`
	Value    string `json:"value"`
	Children string `json:"children"`
}

type TreeData struct {
	Title           string      `json:"title"`
	Value           interface{} `json:"value"`
	Children        []*TreeData `json:"children,omitempty"`
	Disabled        bool        `json:"disabled,omitempty"`
	DisableCheckbox bool        `json:"disableCheckbox,omitempty"`
	Selectable      bool        `json:"selectable,omitempty"`
	Checkable       bool        `json:"checkable,omitempty"`
}

type Component struct {
	ComponentKey string `json:"componentkey"` // 组件标识
	Component    string `json:"component"`    // 组件名称

	Colon         bool        `json:"colon,omitempty"`         // 配合 label 属性使用，表示是否显示 label 后面的冒号
	Extra         string      `json:"extra,omitempty"`         // 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
	HasFeedback   bool        `json:"hasFeedback,omitempty"`   // 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
	Help          string      `json:"help,omitempty"`          // 提示信息，如不设置，则会根据校验规则自动生成
	Hidden        bool        `json:"hidden,omitempty"`        // 是否隐藏字段（依然会收集和校验字段）
	InitialValue  interface{} `json:"initialValue,omitempty"`  // 设置子元素默认值，如果与 Form 的 initialValues 冲突则以 Form 为准
	Label         string      `json:"label,omitempty"`         // label 标签的文本
	LabelAlign    string      `json:"labelAlign,omitempty"`    // 标签文本对齐方式
	LabelCol      interface{} `json:"labelCol,omitempty"`      // label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。你可以通过 Form 的 labelCol 进行统一设置，不会作用于嵌套 Item。当和 Form 同时设置时，以 Item 为准
	Name          string      `json:"name,omitempty"`          // 字段名，支持数组
	NoStyle       bool        `json:"noStyle,omitempty"`       // 为 true 时不带样式，作为纯字段控件使用
	Required      bool        `json:"required,omitempty"`      // 必填样式设置。如不设置，则会根据校验规则自动生成
	Tooltip       string      `json:"tooltip,omitempty"`       // 会在 label 旁增加一个 icon，悬浮后展示配置的信息
	ValuePropName string      `json:"valuePropName,omitempty"` // 子节点的值的属性，如 Switch 的是 'checked'。该属性为 getValueProps 的封装，自定义 getValueProps 后会失效
	WrapperCol    interface{} `json:"wrapperCol,omitempty"`    // 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。你可以通过 Form 的 wrapperCol 进行统一设置，不会作用于嵌套 Item。当和 Form 同时设置时，以 Item 为准

	Api            string          `json:"api,omitempty"` // 获取数据接口
	Ignore         bool            `json:"ignore"`        // 是否忽略保存到数据库，默认为 false
	Rules          []*rule.Rule    `json:"-"`             // 全局校验规则
	CreationRules  []*rule.Rule    `json:"-"`             // 创建页校验规则
	UpdateRules    []*rule.Rule    `json:"-"`             // 编辑页校验规则
	FrontendRules  []*rule.Rule    `json:"frontendRules"` // 前端校验规则，设置字段的校验逻辑
	When           *when.Component `json:"when"`          //
	WhenItem       []*when.Item    `json:"-"`             //
	ShowOnIndex    bool            `json:"-"`             // 在列表页展示
	ShowOnDetail   bool            `json:"-"`             // 在详情页展示
	ShowOnCreation bool            `json:"-"`             // 在创建页面展示
	ShowOnUpdate   bool            `json:"-"`             // 在编辑页面展示
	ShowOnExport   bool            `json:"-"`             // 在导出的Excel上展示
	ShowOnImport   bool            `json:"-"`             // 在导入Excel上展示
	Editable       bool            `json:"-"`             // 表格上是否可编辑
	Column         *table.Column   `json:"-"`             // 表格列
	Callback       interface{}     `json:"-"`             // 回调函数

	AllowClear               bool                   `json:"allowClear,omitempty"`               // 可以点击清除图标删除内容
	AutoClearSearchValue     bool                   `json:"autoClearSearchValue,omitempty"`     // 是否在选中项后清空搜索框，只在 mode 为 multiple 或 tags 时有效
	Bordered                 bool                   `json:"bordered,omitempty"`                 // 是否有边框
	DefaultValue             interface{}            `json:"defaultValue,omitempty"`             // 默认选中的选项
	Disabled                 bool                   `json:"disabled,omitempty"`                 // 整组失效
	PopupClassName           string                 `json:"popupClassName,omitempty"`           // 下拉菜单的 className 属性
	DropdownMatchSelectWidth interface{}            `json:"dropdownMatchSelectWidth,omitempty"` // 下拉菜单和选择器同宽。默认将设置 min-width，当值小于选择框宽度时会被忽略。false 时会关闭虚拟滚动
	DropdownStyle            interface{}            `json:"dropdownStyle,omitempty"`            // 下拉菜单的 style 属性
	FieldNames               *FieldNames            `json:"fieldNames,omitempty"`               // 自定义 options 中 label value children 的字段
	LabelInValue             bool                   `json:"labelInValue,omitempty"`             // 是否把每个选项的 label 包装到 value 中，会把 Select 的 value 类型从 string 变为 { value: string, label: ReactNode } 的格式
	ListHeight               int                    `json:"listHeight,omitempty"`               // 设置弹窗滚动高度 256
	MaxTagCount              int                    `json:"maxTagCount,omitempty"`              // 最多显示多少个 tag，响应式模式会对性能产生损耗
	MaxTagPlaceholder        string                 `json:"maxTagPlaceholder,omitempty"`        // 隐藏 tag 时显示的内容
	MaxTagTextLength         int                    `json:"maxTagTextLength,omitempty"`         // 最大显示的 tag 文本长度
	Multiple                 bool                   `json:"multiple,omitempty"`                 // 支持多选（当设置 treeCheckable 时自动变为 true）
	NotFoundContent          string                 `json:"notFoundContent,omitempty"`          // 当下拉列表为空时显示的内容
	Placeholder              string                 `json:"placeholder,omitempty"`              // 选择框默认文本
	Placement                string                 `json:"placement,omitempty"`                // 选择框弹出的位置 bottomLeft bottomRight topLeft topRight
	SearchValue              string                 `json:"searchValue,omitempty"`              // 控制搜索文本
	ShowArrow                bool                   `json:"showArrow,omitempty"`                // 是否显示下拉小箭头
	ShowSearch               bool                   `json:"showSearch,omitempty"`               // 配置是否可搜索
	Size                     string                 `json:"size,omitempty"`                     // 选择框大小
	Status                   string                 `json:"status,omitempty"`                   // 设置校验状态 'error' | 'warning'
	SuffixIcon               interface{}            `json:"suffixIcon,omitempty"`               // 自定义的选择框后缀图标
	SwitcherIcon             interface{}            `json:"switcherIcon,omitempty"`             // 自定义树节点的展开/折叠图标
	TreeCheckable            bool                   `json:"treeCheckable,omitempty"`            // 显示 Checkbox
	TreeCheckStrictly        bool                   `json:"treeCheckStrictly,omitempty"`        // checkable 状态下节点选择完全受控（父子节点选中状态不再关联），会使得 labelInValue 强制为 true
	TreeData                 []*TreeData            `json:"treeData,omitempty"`                 // treeNodes 数据，如果设置则不需要手动构造 TreeNode 节点（value 在整个树范围内唯一）
	TreeDataSimpleMode       interface{}            `json:"treeDataSimpleMode,omitempty"`       // 使用简单格式的 treeData，具体设置参考可设置的类型 (此时 treeData 应变为这样的数据结构: [{id:1, pId:0, value:'1', title:"test1",...},...]， pId 是父节点的 id)
	TreeDefaultExpandAll     bool                   `json:"treeDefaultExpandAll,omitempty"`     // 默认展开所有树节点
	TreeDefaultExpandedKeys  []interface{}          `json:"treeDefaultExpandedKeys,omitempty"`  // 默认展开的树节点
	TreeExpandAction         interface{}            `json:"treeExpandAction,omitempty"`         // 点击节点 title 时的展开逻辑，可选：false | click | doubleClick
	TreeExpandedKeys         []interface{}          `json:"treeExpandedKeys,omitempty"`         // 设置展开的树节点
	TreeIcon                 bool                   `json:"treeIcon,omitempty"`                 // 是否展示 TreeNode title 前的图标，没有默认样式，如设置为 true，需要自行定义图标相关样式
	TreeLine                 bool                   `json:"treeLine,omitempty"`                 // 是否展示线条样式
	TreeNodeFilterProp       string                 `json:"treeNodeFilterProp,omitempty"`       // 输入项过滤对应的 treeNode 属性
	TreeNodeLabelProp        string                 `json:"treeNodeLabelProp,omitempty"`        // 作为显示的 prop 设置
	Value                    interface{}            `json:"value,omitempty"`                    // 指定当前选中的条目，多选时为一个数组。（value 数组引用未变化时，Select 不会更新）
	Virtual                  bool                   `json:"virtual,omitempty"`                  // 设置 false 时关闭虚拟滚动
	Style                    map[string]interface{} `json:"style,omitempty"`                    // 自定义样式
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "treeSelectField"
	p.Colon = true
	p.LabelAlign = "right"
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = true
	p.ShowOnImport = true
	p.AllowClear = true
	p.Column = (&table.Column{}).Init()
	p.TreeDefaultExpandAll = true
	p.TreeLine = true

	p.SetWidth(200)
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置Key
func (p *Component) SetKey(key string, crypt bool) *Component {
	p.ComponentKey = utils.MakeKey(key, crypt)

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *Component) SetTooltip(tooltip string) *Component {
	p.Tooltip = tooltip

	return p
}

// Field 的长度，我们归纳了常用的 Field 长度以及适合的场景，支持了一些枚举 "xs" , "s" , "m" , "l" , "x"
func (p *Component) SetWidth(width interface{}) *Component {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *Component) SetColon(colon bool) *Component {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *Component) SetExtra(extra string) *Component {
	p.Extra = extra
	return p
}

// 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Component) SetHasFeedback(hasFeedback bool) *Component {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Component) SetHelp(help string) *Component {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *Component) SetNoStyle() *Component {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *Component) SetLabel(label string) *Component {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *Component) SetLabelAlign(align string) *Component {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *Component) SetLabelCol(col interface{}) *Component {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *Component) SetName(name string) *Component {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *Component) SetRequired() *Component {
	p.Required = true
	return p
}

// 获取前端验证规则
func (p *Component) GetFrontendRules(path string) interface{} {
	var (
		frontendRules []*rule.Rule
		rules         []*rule.Rule
		creationRules []*rule.Rule
		updateRules   []*rule.Rule
	)

	uri := strings.Split(path, "/")
	isCreating := (uri[len(uri)-1] == "create") || (uri[len(uri)-1] == "store")
	isEditing := (uri[len(uri)-1] == "edit") || (uri[len(uri)-1] == "update")

	if len(p.Rules) > 0 {
		rules = rule.ConvertToFrontendRules(p.Rules)
	}
	if isCreating && len(p.CreationRules) > 0 {
		creationRules = rule.ConvertToFrontendRules(p.CreationRules)
	}
	if isEditing && len(p.UpdateRules) > 0 {
		updateRules = rule.ConvertToFrontendRules(p.UpdateRules)
	}
	if len(rules) > 0 {
		frontendRules = append(frontendRules, rules...)
	}
	if len(creationRules) > 0 {
		frontendRules = append(frontendRules, creationRules...)
	}
	if len(updateRules) > 0 {
		frontendRules = append(frontendRules, updateRules...)
	}

	p.FrontendRules = frontendRules

	return p
}

// 校验规则，设置字段的校验逻辑
func (p *Component) SetRules(rules []*rule.Rule) *Component {
	for k, v := range rules {
		rules[k] = v.SetName(p.Name)
	}
	p.Rules = rules

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *Component) SetCreationRules(rules []*rule.Rule) *Component {
	for k, v := range rules {
		rules[k] = v.SetName(p.Name)
	}
	p.CreationRules = rules

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *Component) SetUpdateRules(rules []*rule.Rule) *Component {
	for k, v := range rules {
		rules[k] = v.SetName(p.Name)
	}
	p.UpdateRules = rules

	return p
}

// 获取全局验证规则
func (p *Component) GetRules() []*rule.Rule {

	return p.Rules
}

// 获取创建表单验证规则
func (p *Component) GetCreationRules() []*rule.Rule {

	return p.CreationRules
}

// 获取更新表单验证规则
func (p *Component) GetUpdateRules() []*rule.Rule {

	return p.UpdateRules
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *Component) SetValuePropName(valuePropName string) *Component {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *Component) SetWrapperCol(col interface{}) *Component {
	p.WrapperCol = col
	return p
}

// 指定当前选中的条目，多选时为一个数组。（value 数组引用未变化时，Select 不会更新）
func (p *Component) SetValue(value interface{}) *Component {
	p.Value = value
	return p
}

// 设置默认值。
func (p *Component) SetDefault(value interface{}) *Component {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *Component) SetIgnore(ignore bool) *Component {
	p.Ignore = ignore
	return p
}

// 设置When组件数据
func (p *Component) SetWhen(value ...any) *Component {
	w := when.New()
	i := when.NewItem()
	var operator string
	var option any

	if len(value) == 2 {
		operator = "="
		option = value[0]
		callback := value[1].(func() interface{})

		i.Body = callback()
	}

	if len(value) == 3 {
		operator = value[0].(string)
		option = value[1]
		callback := value[2].(func() interface{})

		i.Body = callback()
	}

	getOption := utils.InterfaceToString(option)
	switch operator {
	case "=":
		i.Condition = "<%=String(" + p.Name + ") === '" + getOption + "' %>"
		break
	case ">":
		i.Condition = "<%=String(" + p.Name + ") > '" + getOption + "' %>"
		break
	case "<":
		i.Condition = "<%=String(" + p.Name + ") < '" + getOption + "' %>"
		break
	case "<=":
		i.Condition = "<%=String(" + p.Name + ") <= '" + getOption + "' %>"
		break
	case ">=":
		i.Condition = "<%=String(" + p.Name + ") => '" + getOption + "' %>"
		break
	case "has":
		i.Condition = "<%=(String(" + p.Name + ").indexOf('" + getOption + "') !=-1) %>"
		break
	case "in":
		jsonStr, _ := json.Marshal(option)
		i.Condition = "<%=(" + string(jsonStr) + ".indexOf(" + p.Name + ") !=-1) %>"
		break
	default:
		i.Condition = "<%=String(" + p.Name + ") === '" + getOption + "' %>"
		break
	}

	i.ConditionName = p.Name
	i.ConditionOperator = operator
	i.Option = option
	p.WhenItem = append(p.WhenItem, i)
	p.When = w.SetItems(p.WhenItem)

	return p
}

// 获取When组件数据
func (p *Component) GetWhen() *when.Component {

	return p.When
}

// Specify that the element should be hidden from the index view.
func (p *Component) HideFromIndex(callback bool) *Component {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Component) HideFromDetail(callback bool) *Component {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Component) HideWhenCreating(callback bool) *Component {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Component) HideWhenUpdating(callback bool) *Component {
	p.ShowOnUpdate = !callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Component) HideWhenExporting(callback bool) *Component {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Component) HideWhenImporting(callback bool) *Component {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Component) OnIndexShowing(callback bool) *Component {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Component) OnDetailShowing(callback bool) *Component {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Component) ShowOnCreating(callback bool) *Component {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Component) ShowOnUpdating(callback bool) *Component {
	p.ShowOnUpdate = callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Component) ShowOnExporting(callback bool) *Component {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Component) ShowOnImporting(callback bool) *Component {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *Component) OnlyOnIndex() *Component {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *Component) OnlyOnDetail() *Component {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *Component) OnlyOnForms() *Component {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export file.
func (p *Component) OnlyOnExport() *Component {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import file.
func (p *Component) OnlyOnImport() *Component {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *Component) ExceptOnForms() *Component {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *Component) IsShownOnUpdate() bool {
	return p.ShowOnUpdate
}

// Check showing on index.
func (p *Component) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *Component) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *Component) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *Component) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *Component) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *Component) SetEditable(editable bool) *Component {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *Component) SetColumn(f func(column *table.Column) *table.Column) *Component {
	p.Column = f(p.Column)

	return p
}

// 当前可选项
func (p *Component) GetOptions() []*TreeData {

	return p.TreeData
}

// 设置回调函数
func (p *Component) SetCallback(closure func() interface{}) *Component {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *Component) GetCallback() interface{} {
	return p.Callback
}

// 获取数据接口
func (p *Component) SetApi(api string) *Component {
	p.Api = api

	return p
}

// 可以点击清除图标删除内容
func (p *Component) SetAllowClear(allowClear bool) *Component {
	p.AllowClear = allowClear

	return p
}

// 是否在选中项后清空搜索框，只在 mode 为 multiple 或 tags 时有效
func (p *Component) SetAutoClearSearchValue(autoClearSearchValue bool) *Component {
	p.AutoClearSearchValue = autoClearSearchValue

	return p
}

// 是否有边框
func (p *Component) SetBordered(bordered bool) *Component {
	p.Bordered = bordered

	return p
}

// 下拉菜单的 className 属性
func (p *Component) SetPopupClassName(popupClassName string) *Component {
	p.PopupClassName = popupClassName

	return p
}

// 下拉菜单和选择器同宽。默认将设置 min-width，当值小于选择框宽度时会被忽略。false 时会关闭虚拟滚动
func (p *Component) SetDropdownMatchSelectWidth(dropdownMatchSelectWidth interface{}) *Component {
	p.DropdownMatchSelectWidth = dropdownMatchSelectWidth

	return p
}

// 下拉菜单的 style 属性
func (p *Component) SetDropdownStyle(dropdownStyle interface{}) *Component {
	p.DropdownStyle = dropdownStyle

	return p
}

// 自定义 options 中 label value children 的字段
func (p *Component) SetFieldNames(fieldNames *FieldNames) *Component {
	p.FieldNames = fieldNames

	return p
}

// 是否把每个选项的 label 包装到 value 中，会把 Select 的 value 类型从 string 变为 { value: string, label: ReactNode } 的格式
func (p *Component) SetLabelInValue(labelInValue bool) *Component {
	p.LabelInValue = labelInValue

	return p
}

// 设置弹窗滚动高度 256
func (p *Component) SetListHeight(listHeight int) *Component {
	p.ListHeight = listHeight

	return p
}

// 最多显示多少个 tag，响应式模式会对性能产生损耗
func (p *Component) SetMaxTagCount(maxTagCount int) *Component {
	p.MaxTagCount = maxTagCount

	return p
}

// 隐藏 tag 时显示的内容
func (p *Component) SetMaxTagPlaceholder(maxTagPlaceholder string) *Component {
	p.MaxTagPlaceholder = maxTagPlaceholder

	return p
}

// 最大显示的 tag 文本长度
func (p *Component) SetMaxTagTextLength(maxTagTextLength int) *Component {
	p.MaxTagTextLength = maxTagTextLength

	return p
}

// 支持多选（当设置 treeCheckable 时自动变为 true）
func (p *Component) SetMultiple(multiple bool) *Component {
	p.Multiple = multiple

	return p
}

// 当下拉列表为空时显示的内容
func (p *Component) SetNotFoundContent(notFoundContent string) *Component {
	p.NotFoundContent = notFoundContent

	return p
}

// 选择框默认文本
func (p *Component) SetPlaceholder(placeholder string) *Component {
	p.Placeholder = placeholder

	return p
}

// 选择框弹出的位置 bottomLeft bottomRight topLeft topRight
func (p *Component) SetPlacement(placement string) *Component {
	p.Placement = placement

	return p
}

// 控制搜索文本
func (p *Component) SetSearchValue(searchValue string) *Component {
	p.SearchValue = searchValue

	return p
}

// 是否显示下拉小箭头
func (p *Component) SetShowArrow(showArrow bool) *Component {
	p.ShowArrow = showArrow

	return p
}

// 配置是否可搜索
func (p *Component) SetShowSearch(showSearch bool) *Component {
	p.ShowSearch = showSearch

	return p
}

// 选择框大小
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}

// 设置校验状态 'error' | 'warning'
func (p *Component) SetStatus(status string) *Component {
	p.Status = status

	return p
}

// 自定义的选择框后缀图标
func (p *Component) SetSuffixIcon(suffixIcon interface{}) *Component {
	p.SuffixIcon = suffixIcon

	return p
}

// 自定义树节点的展开/折叠图标
func (p *Component) SetSwitcherIcon(switcherIcon interface{}) *Component {
	p.SwitcherIcon = switcherIcon

	return p
}

// 显示 Checkbox
func (p *Component) SetTreeCheckable(treeCheckable bool) *Component {
	p.TreeCheckable = treeCheckable

	return p
}

// checkable 状态下节点选择完全受控（父子节点选中状态不再关联），会使得 labelInValue 强制为 true
func (p *Component) SetTreeCheckStrictly(treeCheckStrictly bool) *Component {
	p.TreeCheckStrictly = treeCheckStrictly

	return p
}

// treeNodes 数据，如果设置则不需要手动构造 TreeNode 节点（value 在整个树范围内唯一）
func (p *Component) SetTreeData(treeData []*TreeData) *Component {
	p.TreeData = treeData

	return p
}

// treeNodes 数据，如果设置则不需要手动构造 TreeNode 节点（value 在整个树范围内唯一）
func (p *Component) SetData(treeData []*TreeData) *Component {
	p.TreeData = treeData

	return p
}

// 获取 treeNodes 数据
func (p *Component) GetTreeData() []*TreeData {

	return p.TreeData
}

// 获取 treeNodes 数据
func (p *Component) GetData() []*TreeData {

	return p.TreeData
}

// 使用简单格式的 treeData，具体设置参考可设置的类型 (此时 treeData 应变为这样的数据结构: [{id:1, pId:0, value:'1', title:"test1",...},...]， pId 是父节点的 id)
func (p *Component) SetTreeDataSimpleMode(treeDataSimpleMode interface{}) *Component {
	p.TreeDataSimpleMode = treeDataSimpleMode

	return p
}

// 默认展开所有树节点
func (p *Component) SetTreeDefaultExpandAll(treeDefaultExpandAll bool) *Component {
	p.TreeDefaultExpandAll = treeDefaultExpandAll

	return p
}

// 默认展开的树节点
func (p *Component) SetTreeDefaultExpandedKeys(treeDefaultExpandedKeys []interface{}) *Component {
	p.TreeDefaultExpandedKeys = treeDefaultExpandedKeys

	return p
}

// 点击节点 title 时的展开逻辑，可选：false | click | doubleClick
func (p *Component) SetTreeExpandAction(treeExpandAction []interface{}) *Component {
	p.TreeExpandAction = treeExpandAction

	return p
}

// 设置展开的树节点
func (p *Component) SetTreeExpandedKeys(treeExpandedKeys []interface{}) *Component {
	p.TreeExpandedKeys = treeExpandedKeys

	return p
}

// 是否展示 TreeNode title 前的图标，没有默认样式，如设置为 true，需要自行定义图标相关样式
func (p *Component) SetTreeIcon(treeIcon bool) *Component {
	p.TreeIcon = treeIcon

	return p
}

// 是否展示线条样式
func (p *Component) SetTreeLine(treeLine bool) *Component {
	p.TreeLine = treeLine

	return p
}

// 输入项过滤对应的 treeNode 属性
func (p *Component) SetTreeNodeFilterProp(treeNodeFilterProp string) *Component {
	p.TreeNodeFilterProp = treeNodeFilterProp

	return p
}

// 作为显示的 prop 设置
func (p *Component) SetTreeNodeLabelProp(treeNodeLabelProp string) *Component {
	p.TreeNodeLabelProp = treeNodeLabelProp

	return p
}

// 设置 false 时关闭虚拟滚动
func (p *Component) SetVirtual(virtual bool) *Component {
	p.Virtual = virtual

	return p
}

// 自定义样式
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}