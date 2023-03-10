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

func newSysDictionaryDetail(db *gorm.DB, opts ...gen.DOOption) sysDictionaryDetail {
	_sysDictionaryDetail := sysDictionaryDetail{}

	_sysDictionaryDetail.sysDictionaryDetailDo.UseDB(db, opts...)
	_sysDictionaryDetail.sysDictionaryDetailDo.UseModel(&models.SysDictionaryDetail{})

	tableName := _sysDictionaryDetail.sysDictionaryDetailDo.TableName()
	_sysDictionaryDetail.ALL = field.NewAsterisk(tableName)
	_sysDictionaryDetail.ID = field.NewUint64(tableName, "id")
	_sysDictionaryDetail.CreatedAt = field.NewTime(tableName, "created_at")
	_sysDictionaryDetail.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysDictionaryDetail.DeletedAt = field.NewField(tableName, "deleted_at")
	_sysDictionaryDetail.SysDictionaryID = field.NewInt64(tableName, "sys_dictionary_id")
	_sysDictionaryDetail.Label = field.NewString(tableName, "label")
	_sysDictionaryDetail.Value = field.NewString(tableName, "value")
	_sysDictionaryDetail.Status = field.NewBool(tableName, "status")
	_sysDictionaryDetail.Weight = field.NewInt64(tableName, "weight")

	_sysDictionaryDetail.fillFieldMap()

	return _sysDictionaryDetail
}

type sysDictionaryDetail struct {
	sysDictionaryDetailDo sysDictionaryDetailDo

	ALL             field.Asterisk
	ID              field.Uint64
	CreatedAt       field.Time
	UpdatedAt       field.Time
	DeletedAt       field.Field
	SysDictionaryID field.Int64  // 关联标记
	Label           field.String // 展示值
	Value           field.String // 字典值
	Status          field.Bool   // 启用状态
	Weight          field.Int64  // 排序权重

	fieldMap map[string]field.Expr
}

func (s sysDictionaryDetail) Table(newTableName string) *sysDictionaryDetail {
	s.sysDictionaryDetailDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysDictionaryDetail) As(alias string) *sysDictionaryDetail {
	s.sysDictionaryDetailDo.DO = *(s.sysDictionaryDetailDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysDictionaryDetail) updateTableName(table string) *sysDictionaryDetail {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint64(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.SysDictionaryID = field.NewInt64(table, "sys_dictionary_id")
	s.Label = field.NewString(table, "label")
	s.Value = field.NewString(table, "value")
	s.Status = field.NewBool(table, "status")
	s.Weight = field.NewInt64(table, "weight")

	s.fillFieldMap()

	return s
}

func (s *sysDictionaryDetail) WithContext(ctx context.Context) ISysDictionaryDetailDo {
	return s.sysDictionaryDetailDo.WithContext(ctx)
}

func (s sysDictionaryDetail) TableName() string { return s.sysDictionaryDetailDo.TableName() }

func (s sysDictionaryDetail) Alias() string { return s.sysDictionaryDetailDo.Alias() }

func (s *sysDictionaryDetail) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysDictionaryDetail) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["sys_dictionary_id"] = s.SysDictionaryID
	s.fieldMap["label"] = s.Label
	s.fieldMap["value"] = s.Value
	s.fieldMap["status"] = s.Status
	s.fieldMap["weight"] = s.Weight
}

func (s sysDictionaryDetail) clone(db *gorm.DB) sysDictionaryDetail {
	s.sysDictionaryDetailDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysDictionaryDetail) replaceDB(db *gorm.DB) sysDictionaryDetail {
	s.sysDictionaryDetailDo.ReplaceDB(db)
	return s
}

type sysDictionaryDetailDo struct{ gen.DO }

type ISysDictionaryDetailDo interface {
	gen.SubQuery
	Debug() ISysDictionaryDetailDo
	WithContext(ctx context.Context) ISysDictionaryDetailDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysDictionaryDetailDo
	WriteDB() ISysDictionaryDetailDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysDictionaryDetailDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysDictionaryDetailDo
	Not(conds ...gen.Condition) ISysDictionaryDetailDo
	Or(conds ...gen.Condition) ISysDictionaryDetailDo
	Select(conds ...field.Expr) ISysDictionaryDetailDo
	Where(conds ...gen.Condition) ISysDictionaryDetailDo
	Order(conds ...field.Expr) ISysDictionaryDetailDo
	Distinct(cols ...field.Expr) ISysDictionaryDetailDo
	Omit(cols ...field.Expr) ISysDictionaryDetailDo
	Join(table schema.Tabler, on ...field.Expr) ISysDictionaryDetailDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysDictionaryDetailDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysDictionaryDetailDo
	Group(cols ...field.Expr) ISysDictionaryDetailDo
	Having(conds ...gen.Condition) ISysDictionaryDetailDo
	Limit(limit int) ISysDictionaryDetailDo
	Offset(offset int) ISysDictionaryDetailDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysDictionaryDetailDo
	Unscoped() ISysDictionaryDetailDo
	Create(values ...*models.SysDictionaryDetail) error
	CreateInBatches(values []*models.SysDictionaryDetail, batchSize int) error
	Save(values ...*models.SysDictionaryDetail) error
	First() (*models.SysDictionaryDetail, error)
	Take() (*models.SysDictionaryDetail, error)
	Last() (*models.SysDictionaryDetail, error)
	Find() ([]*models.SysDictionaryDetail, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.SysDictionaryDetail, err error)
	FindInBatches(result *[]*models.SysDictionaryDetail, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.SysDictionaryDetail) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysDictionaryDetailDo
	Assign(attrs ...field.AssignExpr) ISysDictionaryDetailDo
	Joins(fields ...field.RelationField) ISysDictionaryDetailDo
	Preload(fields ...field.RelationField) ISysDictionaryDetailDo
	FirstOrInit() (*models.SysDictionaryDetail, error)
	FirstOrCreate() (*models.SysDictionaryDetail, error)
	FindByPage(offset int, limit int) (result []*models.SysDictionaryDetail, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysDictionaryDetailDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysDictionaryDetailDo) Debug() ISysDictionaryDetailDo {
	return s.withDO(s.DO.Debug())
}

func (s sysDictionaryDetailDo) WithContext(ctx context.Context) ISysDictionaryDetailDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysDictionaryDetailDo) ReadDB() ISysDictionaryDetailDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysDictionaryDetailDo) WriteDB() ISysDictionaryDetailDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysDictionaryDetailDo) Session(config *gorm.Session) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysDictionaryDetailDo) Clauses(conds ...clause.Expression) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysDictionaryDetailDo) Returning(value interface{}, columns ...string) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysDictionaryDetailDo) Not(conds ...gen.Condition) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysDictionaryDetailDo) Or(conds ...gen.Condition) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysDictionaryDetailDo) Select(conds ...field.Expr) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysDictionaryDetailDo) Where(conds ...gen.Condition) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysDictionaryDetailDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysDictionaryDetailDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysDictionaryDetailDo) Order(conds ...field.Expr) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysDictionaryDetailDo) Distinct(cols ...field.Expr) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysDictionaryDetailDo) Omit(cols ...field.Expr) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysDictionaryDetailDo) Join(table schema.Tabler, on ...field.Expr) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysDictionaryDetailDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysDictionaryDetailDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysDictionaryDetailDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysDictionaryDetailDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysDictionaryDetailDo) Group(cols ...field.Expr) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysDictionaryDetailDo) Having(conds ...gen.Condition) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysDictionaryDetailDo) Limit(limit int) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysDictionaryDetailDo) Offset(offset int) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysDictionaryDetailDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysDictionaryDetailDo) Unscoped() ISysDictionaryDetailDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysDictionaryDetailDo) Create(values ...*models.SysDictionaryDetail) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysDictionaryDetailDo) CreateInBatches(values []*models.SysDictionaryDetail, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysDictionaryDetailDo) Save(values ...*models.SysDictionaryDetail) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysDictionaryDetailDo) First() (*models.SysDictionaryDetail, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDictionaryDetail), nil
	}
}

func (s sysDictionaryDetailDo) Take() (*models.SysDictionaryDetail, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDictionaryDetail), nil
	}
}

func (s sysDictionaryDetailDo) Last() (*models.SysDictionaryDetail, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDictionaryDetail), nil
	}
}

func (s sysDictionaryDetailDo) Find() ([]*models.SysDictionaryDetail, error) {
	result, err := s.DO.Find()
	return result.([]*models.SysDictionaryDetail), err
}

func (s sysDictionaryDetailDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.SysDictionaryDetail, err error) {
	buf := make([]*models.SysDictionaryDetail, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysDictionaryDetailDo) FindInBatches(result *[]*models.SysDictionaryDetail, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysDictionaryDetailDo) Attrs(attrs ...field.AssignExpr) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysDictionaryDetailDo) Assign(attrs ...field.AssignExpr) ISysDictionaryDetailDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysDictionaryDetailDo) Joins(fields ...field.RelationField) ISysDictionaryDetailDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysDictionaryDetailDo) Preload(fields ...field.RelationField) ISysDictionaryDetailDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysDictionaryDetailDo) FirstOrInit() (*models.SysDictionaryDetail, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDictionaryDetail), nil
	}
}

func (s sysDictionaryDetailDo) FirstOrCreate() (*models.SysDictionaryDetail, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDictionaryDetail), nil
	}
}

func (s sysDictionaryDetailDo) FindByPage(offset int, limit int) (result []*models.SysDictionaryDetail, count int64, err error) {
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

func (s sysDictionaryDetailDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysDictionaryDetailDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysDictionaryDetailDo) Delete(models ...*models.SysDictionaryDetail) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysDictionaryDetailDo) withDO(do gen.Dao) *sysDictionaryDetailDo {
	s.DO = *do.(*gen.DO)
	return s
}
