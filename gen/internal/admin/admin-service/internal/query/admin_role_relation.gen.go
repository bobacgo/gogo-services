// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/gogoclouds/gogo-services/gen/internal/admin/admin-service/internal/model"
)

func newAdminRoleRelation(db *gorm.DB, opts ...gen.DOOption) adminRoleRelation {
	_adminRoleRelation := adminRoleRelation{}

	_adminRoleRelation.adminRoleRelationDo.UseDB(db, opts...)
	_adminRoleRelation.adminRoleRelationDo.UseModel(&model.AdminRoleRelation{})

	tableName := _adminRoleRelation.adminRoleRelationDo.TableName()
	_adminRoleRelation.ALL = field.NewAsterisk(tableName)
	_adminRoleRelation.ID = field.NewInt64(tableName, "id")
	_adminRoleRelation.AdminID = field.NewInt64(tableName, "admin_id")
	_adminRoleRelation.RoleID = field.NewInt64(tableName, "role_id")

	_adminRoleRelation.fillFieldMap()

	return _adminRoleRelation
}

// adminRoleRelation 后台用户和角色关系表
type adminRoleRelation struct {
	adminRoleRelationDo adminRoleRelationDo

	ALL     field.Asterisk
	ID      field.Int64
	AdminID field.Int64
	RoleID  field.Int64

	fieldMap map[string]field.Expr
}

func (a adminRoleRelation) Table(newTableName string) *adminRoleRelation {
	a.adminRoleRelationDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminRoleRelation) As(alias string) *adminRoleRelation {
	a.adminRoleRelationDo.DO = *(a.adminRoleRelationDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminRoleRelation) updateTableName(table string) *adminRoleRelation {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.AdminID = field.NewInt64(table, "admin_id")
	a.RoleID = field.NewInt64(table, "role_id")

	a.fillFieldMap()

	return a
}

func (a *adminRoleRelation) WithContext(ctx context.Context) IAdminRoleRelationDo {
	return a.adminRoleRelationDo.WithContext(ctx)
}

func (a adminRoleRelation) TableName() string { return a.adminRoleRelationDo.TableName() }

func (a adminRoleRelation) Alias() string { return a.adminRoleRelationDo.Alias() }

func (a adminRoleRelation) Columns(cols ...field.Expr) gen.Columns {
	return a.adminRoleRelationDo.Columns(cols...)
}

func (a *adminRoleRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminRoleRelation) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 3)
	a.fieldMap["id"] = a.ID
	a.fieldMap["admin_id"] = a.AdminID
	a.fieldMap["role_id"] = a.RoleID
}

func (a adminRoleRelation) clone(db *gorm.DB) adminRoleRelation {
	a.adminRoleRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminRoleRelation) replaceDB(db *gorm.DB) adminRoleRelation {
	a.adminRoleRelationDo.ReplaceDB(db)
	return a
}

type adminRoleRelationDo struct{ gen.DO }

type IAdminRoleRelationDo interface {
	gen.SubQuery
	Debug() IAdminRoleRelationDo
	WithContext(ctx context.Context) IAdminRoleRelationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminRoleRelationDo
	WriteDB() IAdminRoleRelationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminRoleRelationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminRoleRelationDo
	Not(conds ...gen.Condition) IAdminRoleRelationDo
	Or(conds ...gen.Condition) IAdminRoleRelationDo
	Select(conds ...field.Expr) IAdminRoleRelationDo
	Where(conds ...gen.Condition) IAdminRoleRelationDo
	Order(conds ...field.Expr) IAdminRoleRelationDo
	Distinct(cols ...field.Expr) IAdminRoleRelationDo
	Omit(cols ...field.Expr) IAdminRoleRelationDo
	Join(table schema.Tabler, on ...field.Expr) IAdminRoleRelationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminRoleRelationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminRoleRelationDo
	Group(cols ...field.Expr) IAdminRoleRelationDo
	Having(conds ...gen.Condition) IAdminRoleRelationDo
	Limit(limit int) IAdminRoleRelationDo
	Offset(offset int) IAdminRoleRelationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminRoleRelationDo
	Unscoped() IAdminRoleRelationDo
	Create(values ...*model.AdminRoleRelation) error
	CreateInBatches(values []*model.AdminRoleRelation, batchSize int) error
	Save(values ...*model.AdminRoleRelation) error
	First() (*model.AdminRoleRelation, error)
	Take() (*model.AdminRoleRelation, error)
	Last() (*model.AdminRoleRelation, error)
	Find() ([]*model.AdminRoleRelation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminRoleRelation, err error)
	FindInBatches(result *[]*model.AdminRoleRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AdminRoleRelation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminRoleRelationDo
	Assign(attrs ...field.AssignExpr) IAdminRoleRelationDo
	Joins(fields ...field.RelationField) IAdminRoleRelationDo
	Preload(fields ...field.RelationField) IAdminRoleRelationDo
	FirstOrInit() (*model.AdminRoleRelation, error)
	FirstOrCreate() (*model.AdminRoleRelation, error)
	FindByPage(offset int, limit int) (result []*model.AdminRoleRelation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminRoleRelationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminRoleRelationDo) Debug() IAdminRoleRelationDo {
	return a.withDO(a.DO.Debug())
}

func (a adminRoleRelationDo) WithContext(ctx context.Context) IAdminRoleRelationDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminRoleRelationDo) ReadDB() IAdminRoleRelationDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminRoleRelationDo) WriteDB() IAdminRoleRelationDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminRoleRelationDo) Session(config *gorm.Session) IAdminRoleRelationDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminRoleRelationDo) Clauses(conds ...clause.Expression) IAdminRoleRelationDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminRoleRelationDo) Returning(value interface{}, columns ...string) IAdminRoleRelationDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminRoleRelationDo) Not(conds ...gen.Condition) IAdminRoleRelationDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminRoleRelationDo) Or(conds ...gen.Condition) IAdminRoleRelationDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminRoleRelationDo) Select(conds ...field.Expr) IAdminRoleRelationDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminRoleRelationDo) Where(conds ...gen.Condition) IAdminRoleRelationDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminRoleRelationDo) Order(conds ...field.Expr) IAdminRoleRelationDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminRoleRelationDo) Distinct(cols ...field.Expr) IAdminRoleRelationDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminRoleRelationDo) Omit(cols ...field.Expr) IAdminRoleRelationDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminRoleRelationDo) Join(table schema.Tabler, on ...field.Expr) IAdminRoleRelationDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminRoleRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminRoleRelationDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminRoleRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminRoleRelationDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminRoleRelationDo) Group(cols ...field.Expr) IAdminRoleRelationDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminRoleRelationDo) Having(conds ...gen.Condition) IAdminRoleRelationDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminRoleRelationDo) Limit(limit int) IAdminRoleRelationDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminRoleRelationDo) Offset(offset int) IAdminRoleRelationDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminRoleRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminRoleRelationDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminRoleRelationDo) Unscoped() IAdminRoleRelationDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminRoleRelationDo) Create(values ...*model.AdminRoleRelation) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminRoleRelationDo) CreateInBatches(values []*model.AdminRoleRelation, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminRoleRelationDo) Save(values ...*model.AdminRoleRelation) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminRoleRelationDo) First() (*model.AdminRoleRelation, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRoleRelation), nil
	}
}

func (a adminRoleRelationDo) Take() (*model.AdminRoleRelation, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRoleRelation), nil
	}
}

func (a adminRoleRelationDo) Last() (*model.AdminRoleRelation, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRoleRelation), nil
	}
}

func (a adminRoleRelationDo) Find() ([]*model.AdminRoleRelation, error) {
	result, err := a.DO.Find()
	return result.([]*model.AdminRoleRelation), err
}

func (a adminRoleRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminRoleRelation, err error) {
	buf := make([]*model.AdminRoleRelation, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminRoleRelationDo) FindInBatches(result *[]*model.AdminRoleRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminRoleRelationDo) Attrs(attrs ...field.AssignExpr) IAdminRoleRelationDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminRoleRelationDo) Assign(attrs ...field.AssignExpr) IAdminRoleRelationDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminRoleRelationDo) Joins(fields ...field.RelationField) IAdminRoleRelationDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminRoleRelationDo) Preload(fields ...field.RelationField) IAdminRoleRelationDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminRoleRelationDo) FirstOrInit() (*model.AdminRoleRelation, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRoleRelation), nil
	}
}

func (a adminRoleRelationDo) FirstOrCreate() (*model.AdminRoleRelation, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminRoleRelation), nil
	}
}

func (a adminRoleRelationDo) FindByPage(offset int, limit int) (result []*model.AdminRoleRelation, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminRoleRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminRoleRelationDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminRoleRelationDo) Delete(models ...*model.AdminRoleRelation) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminRoleRelationDo) withDO(do gen.Dao) *adminRoleRelationDo {
	a.DO = *do.(*gen.DO)
	return a
}
