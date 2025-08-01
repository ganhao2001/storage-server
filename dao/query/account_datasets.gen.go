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

	"webdav/dao/model"
)

func newAccountDataset(db *gorm.DB, opts ...gen.DOOption) accountDataset {
	_accountDataset := accountDataset{}

	_accountDataset.accountDatasetDo.UseDB(db, opts...)
	_accountDataset.accountDatasetDo.UseModel(&model.AccountDataset{})

	tableName := _accountDataset.accountDatasetDo.TableName()
	_accountDataset.ALL = field.NewAsterisk(tableName)
	_accountDataset.ID = field.NewUint(tableName, "id")
	_accountDataset.CreatedAt = field.NewTime(tableName, "created_at")
	_accountDataset.UpdatedAt = field.NewTime(tableName, "updated_at")
	_accountDataset.DeletedAt = field.NewField(tableName, "deleted_at")
	_accountDataset.AccountID = field.NewUint(tableName, "account_id")
	_accountDataset.DatasetID = field.NewUint(tableName, "dataset_id")

	_accountDataset.fillFieldMap()

	return _accountDataset
}

type accountDataset struct {
	accountDatasetDo accountDatasetDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	AccountID field.Uint
	DatasetID field.Uint

	fieldMap map[string]field.Expr
}

func (a accountDataset) Table(newTableName string) *accountDataset {
	a.accountDatasetDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a accountDataset) As(alias string) *accountDataset {
	a.accountDatasetDo.DO = *(a.accountDatasetDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *accountDataset) updateTableName(table string) *accountDataset {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.AccountID = field.NewUint(table, "account_id")
	a.DatasetID = field.NewUint(table, "dataset_id")

	a.fillFieldMap()

	return a
}

func (a *accountDataset) WithContext(ctx context.Context) IAccountDatasetDo {
	return a.accountDatasetDo.WithContext(ctx)
}

func (a accountDataset) TableName() string { return a.accountDatasetDo.TableName() }

func (a accountDataset) Alias() string { return a.accountDatasetDo.Alias() }

func (a accountDataset) Columns(cols ...field.Expr) gen.Columns {
	return a.accountDatasetDo.Columns(cols...)
}

func (a *accountDataset) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *accountDataset) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["account_id"] = a.AccountID
	a.fieldMap["dataset_id"] = a.DatasetID
}

func (a accountDataset) clone(db *gorm.DB) accountDataset {
	a.accountDatasetDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a accountDataset) replaceDB(db *gorm.DB) accountDataset {
	a.accountDatasetDo.ReplaceDB(db)
	return a
}

type accountDatasetDo struct{ gen.DO }

type IAccountDatasetDo interface {
	gen.SubQuery
	Debug() IAccountDatasetDo
	WithContext(ctx context.Context) IAccountDatasetDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAccountDatasetDo
	WriteDB() IAccountDatasetDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAccountDatasetDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAccountDatasetDo
	Not(conds ...gen.Condition) IAccountDatasetDo
	Or(conds ...gen.Condition) IAccountDatasetDo
	Select(conds ...field.Expr) IAccountDatasetDo
	Where(conds ...gen.Condition) IAccountDatasetDo
	Order(conds ...field.Expr) IAccountDatasetDo
	Distinct(cols ...field.Expr) IAccountDatasetDo
	Omit(cols ...field.Expr) IAccountDatasetDo
	Join(table schema.Tabler, on ...field.Expr) IAccountDatasetDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAccountDatasetDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAccountDatasetDo
	Group(cols ...field.Expr) IAccountDatasetDo
	Having(conds ...gen.Condition) IAccountDatasetDo
	Limit(limit int) IAccountDatasetDo
	Offset(offset int) IAccountDatasetDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAccountDatasetDo
	Unscoped() IAccountDatasetDo
	Create(values ...*model.AccountDataset) error
	CreateInBatches(values []*model.AccountDataset, batchSize int) error
	Save(values ...*model.AccountDataset) error
	First() (*model.AccountDataset, error)
	Take() (*model.AccountDataset, error)
	Last() (*model.AccountDataset, error)
	Find() ([]*model.AccountDataset, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AccountDataset, err error)
	FindInBatches(result *[]*model.AccountDataset, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AccountDataset) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAccountDatasetDo
	Assign(attrs ...field.AssignExpr) IAccountDatasetDo
	Joins(fields ...field.RelationField) IAccountDatasetDo
	Preload(fields ...field.RelationField) IAccountDatasetDo
	FirstOrInit() (*model.AccountDataset, error)
	FirstOrCreate() (*model.AccountDataset, error)
	FindByPage(offset int, limit int) (result []*model.AccountDataset, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAccountDatasetDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a accountDatasetDo) Debug() IAccountDatasetDo {
	return a.withDO(a.DO.Debug())
}

func (a accountDatasetDo) WithContext(ctx context.Context) IAccountDatasetDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a accountDatasetDo) ReadDB() IAccountDatasetDo {
	return a.Clauses(dbresolver.Read)
}

func (a accountDatasetDo) WriteDB() IAccountDatasetDo {
	return a.Clauses(dbresolver.Write)
}

func (a accountDatasetDo) Session(config *gorm.Session) IAccountDatasetDo {
	return a.withDO(a.DO.Session(config))
}

func (a accountDatasetDo) Clauses(conds ...clause.Expression) IAccountDatasetDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a accountDatasetDo) Returning(value interface{}, columns ...string) IAccountDatasetDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a accountDatasetDo) Not(conds ...gen.Condition) IAccountDatasetDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a accountDatasetDo) Or(conds ...gen.Condition) IAccountDatasetDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a accountDatasetDo) Select(conds ...field.Expr) IAccountDatasetDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a accountDatasetDo) Where(conds ...gen.Condition) IAccountDatasetDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a accountDatasetDo) Order(conds ...field.Expr) IAccountDatasetDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a accountDatasetDo) Distinct(cols ...field.Expr) IAccountDatasetDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a accountDatasetDo) Omit(cols ...field.Expr) IAccountDatasetDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a accountDatasetDo) Join(table schema.Tabler, on ...field.Expr) IAccountDatasetDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a accountDatasetDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAccountDatasetDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a accountDatasetDo) RightJoin(table schema.Tabler, on ...field.Expr) IAccountDatasetDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a accountDatasetDo) Group(cols ...field.Expr) IAccountDatasetDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a accountDatasetDo) Having(conds ...gen.Condition) IAccountDatasetDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a accountDatasetDo) Limit(limit int) IAccountDatasetDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a accountDatasetDo) Offset(offset int) IAccountDatasetDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a accountDatasetDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAccountDatasetDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a accountDatasetDo) Unscoped() IAccountDatasetDo {
	return a.withDO(a.DO.Unscoped())
}

func (a accountDatasetDo) Create(values ...*model.AccountDataset) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a accountDatasetDo) CreateInBatches(values []*model.AccountDataset, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a accountDatasetDo) Save(values ...*model.AccountDataset) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a accountDatasetDo) First() (*model.AccountDataset, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountDataset), nil
	}
}

func (a accountDatasetDo) Take() (*model.AccountDataset, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountDataset), nil
	}
}

func (a accountDatasetDo) Last() (*model.AccountDataset, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountDataset), nil
	}
}

func (a accountDatasetDo) Find() ([]*model.AccountDataset, error) {
	result, err := a.DO.Find()
	return result.([]*model.AccountDataset), err
}

func (a accountDatasetDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AccountDataset, err error) {
	buf := make([]*model.AccountDataset, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a accountDatasetDo) FindInBatches(result *[]*model.AccountDataset, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a accountDatasetDo) Attrs(attrs ...field.AssignExpr) IAccountDatasetDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a accountDatasetDo) Assign(attrs ...field.AssignExpr) IAccountDatasetDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a accountDatasetDo) Joins(fields ...field.RelationField) IAccountDatasetDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a accountDatasetDo) Preload(fields ...field.RelationField) IAccountDatasetDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a accountDatasetDo) FirstOrInit() (*model.AccountDataset, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountDataset), nil
	}
}

func (a accountDatasetDo) FirstOrCreate() (*model.AccountDataset, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AccountDataset), nil
	}
}

func (a accountDatasetDo) FindByPage(offset int, limit int) (result []*model.AccountDataset, count int64, err error) {
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

func (a accountDatasetDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a accountDatasetDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a accountDatasetDo) Delete(models ...*model.AccountDataset) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *accountDatasetDo) withDO(do gen.Dao) *accountDatasetDo {
	a.DO = *do.(*gen.DO)
	return a
}
