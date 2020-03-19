package tools

import orm "go-admin/database"

type SysColumns struct {
	ColumnId      int64  `gorm:"column:configId;primary_key" json:"configId"`
	TableId       int64  `gorm:"column:configId" json:"configId"`
	ColumnName    string `gorm:"column:configId" json:"configId"`
	ColumnComment string `gorm:"column:column_comment" json:"columnComment"`
	ColumnType    string `gorm:"column:column_type" json:"columnType"`
	GoType        string `gorm:"column:go_type" json:"goType"`
	GoField       string `gorm:"column:go_field" json:"goField"`
	IsPk          string `gorm:"column:is_pk" json:"isPk"`
	IsIncrement   string `gorm:"column:is_increment" json:"isIncrement"`
	IsRequired    string `gorm:"column:is_required" json:"isRequired"`
	IsInsert      string `gorm:"column:is_insert" json:"isInsert"`
	IsEdit        string `gorm:"column:is_edit" json:"isEdit"`
	IsList        string `gorm:"column:is_list" json:"isList"`
	IsQuery       string `gorm:"column:is_query" json:"isQuery"`
	QueryType     string `gorm:"column:query_type" json:"queryType"`
	HtmlType      string `gorm:"column:html_type" json:"htmlType"`
	DictType      string `gorm:"column:dict_type" json:"dictType"`
	Sort          string `gorm:"column:sort" json:"sort"`
	List          string `gorm:"column:list" json:"list"`
	Pk            bool   `gorm:"column:pk" json:"pk"`
	Required      bool   `gorm:"column:required" json:"required"`
	SuperColumn   bool   `gorm:"column:super_column" json:"superColumn"`
	UsableColumn  bool   `gorm:"column:usable_column" json:"usableColumn"`
	Increment     bool   `gorm:"column:increment" json:"increment"`
	Insert        bool   `gorm:"column:insert" json:"insert"`
	Edit          bool   `gorm:"column:edit" json:"edit"`
	Query         bool   `gorm:"column:query" json:"query"`
	Remark        string `gorm:"column:remark" json:"remark"`
	CreateBy      string `gorm:"column:create_by" json:"createBy"`
	CreateTime    string `gorm:"column:create_time" json:"createTime"`
	UpdateBy      string `gorm:"column:update_By" json:"updateBy"`
	UpdateTime    string `gorm:"column:update_time" json:"updateTime"`
}

func (e *SysColumns) GetList() ([]SysColumns, error) {
	var doc []SysColumns

	table := orm.Eloquent.Select("*").Table("sys_columns")

	table = table.Where("table_id = ?", e.TableId)


	if err := table.Find(&doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}