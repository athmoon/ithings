package mysql

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/domain/schema"
	"github.com/i-Things/things/shared/store"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductSchemaModel = (*customProductSchemaModel)(nil)

type (
	ProductSchemaFilter struct {
		ProductID   string   //产品id  必填
		Type        int64    //物模型类型 1:property属性 2:event事件 3:action行为
		Tag         int64    //过滤条件: 物模型标签 1:自定义 2:可选 3:必选
		Identifiers []string //过滤标识符列表
	}
	PropertyDef struct {
		Define schema.Define       `json:"define"` //数据定义
		Mode   schema.PropertyMode `json:"mode"`   //读写类型: 1:r(只读) 2:rw(可读可写)
	}
	EventDef struct {
		Type   schema.EventType `json:"type"`   //事件类型: 1:信息:info  2:告警alert  3:故障:fault
		Params schema.Params    `json:"params"` //事件参数
	}
	ActionDef struct {
		Input  schema.Params `json:"input"`  //调用参数
		Output schema.Params `json:"output"` //返回参数
	}
	ProductSchemaModel interface {
		productSchemaModel
		DeleteWithFilter(ctx context.Context, filter ProductSchemaFilter) error
		FindByFilter(ctx context.Context, filter ProductSchemaFilter, page def.PageInfo) ([]*ProductSchema, error)
		GetCountByFilter(ctx context.Context, filter ProductSchemaFilter) (size int64, err error)
	}

	customProductSchemaModel struct {
		*defaultProductSchemaModel
	}
)

func (p *ProductSchemaFilter) FmtSql(sql sq.SelectBuilder) sq.SelectBuilder {
	if p.ProductID != "" {
		sql = sql.Where("productID=?", p.ProductID)
	}
	if p.Type != 0 {
		sql = sql.Where("type=?", p.Type)
	}
	if p.Tag != 0 {
		sql = sql.Where("tag=?", p.Tag)
	}
	if len(p.Identifiers) != 0 {
		sql = sql.Where(fmt.Sprintf("identifier in (%v)", store.ArrayToSql(p.Identifiers)))
	}
	return sql
}

// NewProductSchemaModel returns a model for the database table.
func NewProductSchemaModel(conn sqlx.SqlConn) *customProductSchemaModel {
	return &customProductSchemaModel{
		defaultProductSchemaModel: newProductSchemaModel(conn),
	}
}
func (p *customProductSchemaModel) DeleteWithFilter(ctx context.Context, filter ProductSchemaFilter) error {
	query := fmt.Sprintf("delete from %s where `productID` = ?", p.table)
	_, err := p.conn.ExecCtx(ctx, query, filter.ProductID)
	return err
}

func (p customProductSchemaModel) FindByFilter(
	ctx context.Context, filter ProductSchemaFilter, page def.PageInfo) (
	[]*ProductSchema, error) {
	var resp []*ProductSchema
	sql := sq.Select(productSchemaRows).From(p.table).Limit(uint64(page.GetLimit())).Offset(uint64(page.GetOffset()))
	sql = filter.FmtSql(sql)
	query, arg, err := sql.ToSql()
	if err != nil {
		return nil, err
	}
	err = p.conn.QueryRowsCtx(ctx, &resp, query, arg...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (p customProductSchemaModel) GetCountByFilter(ctx context.Context, filter ProductSchemaFilter) (size int64, err error) {
	sql := sq.Select("count(1)").From(p.table)
	sql = filter.FmtSql(sql)
	query, arg, err := sql.ToSql()
	if err != nil {
		return 0, err
	}
	err = p.conn.QueryRowCtx(ctx, &size, query, arg...)

	switch err {
	case nil:
		return size, nil
	default:
		return 0, err
	}
}
