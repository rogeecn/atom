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

	"atom/database/models"
)

func newSysDictionary(db *gorm.DB, opts ...gen.DOOption) sysDictionary {
	_sysDictionary := sysDictionary{}

	_sysDictionary.sysDictionaryDo.UseDB(db, opts...)
	_sysDictionary.sysDictionaryDo.UseModel(&models.SysDictionary{})

	tableName := _sysDictionary.sysDictionaryDo.TableName()
	_sysDictionary.ALL = field.NewAsterisk(tableName)
	_sysDictionary.ID = field.NewUint64(tableName, "id")
	_sysDictionary.CreatedAt = field.NewTime(tableName, "created_at")
	_sysDictionary.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysDictionary.DeletedAt = field.NewField(tableName, "deleted_at")
	_sysDictionary.Name = field.NewString(tableName, "name")
	_sysDictionary.Alias_ = field.NewString(tableName, "alias")
	_sysDictionary.Status = field.NewBool(tableName, "status")
	_sysDictionary.Description = field.NewString(tableName, "description")

	_sysDictionary.fillFieldMap()

	return _sysDictionary
}

type sysDictionary struct {
	sysDictionaryDo sysDictionaryDo

	ALL         field.Asterisk
	ID          field.Uint64
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Name        field.String // 字典名（中）
	Alias_      field.String // 字典名（英）
	Status      field.Bool   // 状态
	Description field.String // 描述

	fieldMap map[string]field.Expr
}

func (s sysDictionary) Table(newTableName string) *sysDictionary {
	s.sysDictionaryDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysDictionary) As(alias string) *sysDictionary {
	s.sysDictionaryDo.DO = *(s.sysDictionaryDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysDictionary) updateTableName(table string) *sysDictionary {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint64(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.Name = field.NewString(table, "name")
	s.Alias_ = field.NewString(table, "alias")
	s.Status = field.NewBool(table, "status")
	s.Description = field.NewString(table, "description")

	s.fillFieldMap()

	return s
}

func (s *sysDictionary) WithContext(ctx context.Context) ISysDictionaryDo {
	return s.sysDictionaryDo.WithContext(ctx)
}

func (s sysDictionary) TableName() string { return s.sysDictionaryDo.TableName() }

func (s sysDictionary) Alias() string { return s.sysDictionaryDo.Alias() }

func (s *sysDictionary) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysDictionary) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["name"] = s.Name
	s.fieldMap["alias"] = s.Alias_
	s.fieldMap["status"] = s.Status
	s.fieldMap["description"] = s.Description
}

func (s sysDictionary) clone(db *gorm.DB) sysDictionary {
	s.sysDictionaryDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysDictionary) replaceDB(db *gorm.DB) sysDictionary {
	s.sysDictionaryDo.ReplaceDB(db)
	return s
}

type sysDictionaryDo struct{ gen.DO }

type ISysDictionaryDo interface {
	gen.SubQuery
	Debug() ISysDictionaryDo
	WithContext(ctx context.Context) ISysDictionaryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysDictionaryDo
	WriteDB() ISysDictionaryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysDictionaryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysDictionaryDo
	Not(conds ...gen.Condition) ISysDictionaryDo
	Or(conds ...gen.Condition) ISysDictionaryDo
	Select(conds ...field.Expr) ISysDictionaryDo
	Where(conds ...gen.Condition) ISysDictionaryDo
	Order(conds ...field.Expr) ISysDictionaryDo
	Distinct(cols ...field.Expr) ISysDictionaryDo
	Omit(cols ...field.Expr) ISysDictionaryDo
	Join(table schema.Tabler, on ...field.Expr) ISysDictionaryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysDictionaryDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysDictionaryDo
	Group(cols ...field.Expr) ISysDictionaryDo
	Having(conds ...gen.Condition) ISysDictionaryDo
	Limit(limit int) ISysDictionaryDo
	Offset(offset int) ISysDictionaryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysDictionaryDo
	Unscoped() ISysDictionaryDo
	Create(values ...*models.SysDictionary) error
	CreateInBatches(values []*models.SysDictionary, batchSize int) error
	Save(values ...*models.SysDictionary) error
	First() (*models.SysDictionary, error)
	Take() (*models.SysDictionary, error)
	Last() (*models.SysDictionary, error)
	Find() ([]*models.SysDictionary, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.SysDictionary, err error)
	FindInBatches(result *[]*models.SysDictionary, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.SysDictionary) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysDictionaryDo
	Assign(attrs ...field.AssignExpr) ISysDictionaryDo
	Joins(fields ...field.RelationField) ISysDictionaryDo
	Preload(fields ...field.RelationField) ISysDictionaryDo
	FirstOrInit() (*models.SysDictionary, error)
	FirstOrCreate() (*models.SysDictionary, error)
	FindByPage(offset int, limit int) (result []*models.SysDictionary, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysDictionaryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysDictionaryDo) Debug() ISysDictionaryDo {
	return s.withDO(s.DO.Debug())
}

func (s sysDictionaryDo) WithContext(ctx context.Context) ISysDictionaryDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysDictionaryDo) ReadDB() ISysDictionaryDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysDictionaryDo) WriteDB() ISysDictionaryDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysDictionaryDo) Session(config *gorm.Session) ISysDictionaryDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysDictionaryDo) Clauses(conds ...clause.Expression) ISysDictionaryDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysDictionaryDo) Returning(value interface{}, columns ...string) ISysDictionaryDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysDictionaryDo) Not(conds ...gen.Condition) ISysDictionaryDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysDictionaryDo) Or(conds ...gen.Condition) ISysDictionaryDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysDictionaryDo) Select(conds ...field.Expr) ISysDictionaryDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysDictionaryDo) Where(conds ...gen.Condition) ISysDictionaryDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysDictionaryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysDictionaryDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysDictionaryDo) Order(conds ...field.Expr) ISysDictionaryDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysDictionaryDo) Distinct(cols ...field.Expr) ISysDictionaryDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysDictionaryDo) Omit(cols ...field.Expr) ISysDictionaryDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysDictionaryDo) Join(table schema.Tabler, on ...field.Expr) ISysDictionaryDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysDictionaryDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysDictionaryDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysDictionaryDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysDictionaryDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysDictionaryDo) Group(cols ...field.Expr) ISysDictionaryDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysDictionaryDo) Having(conds ...gen.Condition) ISysDictionaryDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysDictionaryDo) Limit(limit int) ISysDictionaryDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysDictionaryDo) Offset(offset int) ISysDictionaryDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysDictionaryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysDictionaryDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysDictionaryDo) Unscoped() ISysDictionaryDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysDictionaryDo) Create(values ...*models.SysDictionary) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysDictionaryDo) CreateInBatches(values []*models.SysDictionary, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysDictionaryDo) Save(values ...*models.SysDictionary) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysDictionaryDo) First() (*models.SysDictionary, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDictionary), nil
	}
}

func (s sysDictionaryDo) Take() (*models.SysDictionary, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDictionary), nil
	}
}

func (s sysDictionaryDo) Last() (*models.SysDictionary, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDictionary), nil
	}
}

func (s sysDictionaryDo) Find() ([]*models.SysDictionary, error) {
	result, err := s.DO.Find()
	return result.([]*models.SysDictionary), err
}

func (s sysDictionaryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.SysDictionary, err error) {
	buf := make([]*models.SysDictionary, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysDictionaryDo) FindInBatches(result *[]*models.SysDictionary, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysDictionaryDo) Attrs(attrs ...field.AssignExpr) ISysDictionaryDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysDictionaryDo) Assign(attrs ...field.AssignExpr) ISysDictionaryDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysDictionaryDo) Joins(fields ...field.RelationField) ISysDictionaryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysDictionaryDo) Preload(fields ...field.RelationField) ISysDictionaryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysDictionaryDo) FirstOrInit() (*models.SysDictionary, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDictionary), nil
	}
}

func (s sysDictionaryDo) FirstOrCreate() (*models.SysDictionary, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDictionary), nil
	}
}

func (s sysDictionaryDo) FindByPage(offset int, limit int) (result []*models.SysDictionary, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysDictionaryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysDictionaryDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysDictionaryDo) Delete(models ...*models.SysDictionary) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysDictionaryDo) withDO(do gen.Dao) *sysDictionaryDo {
	s.DO = *do.(*gen.DO)
	return s
}
