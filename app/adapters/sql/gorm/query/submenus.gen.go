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

	"golang-auth-app/app/adapters/sql/gorm/model"
)

func newSubmenu(db *gorm.DB, opts ...gen.DOOption) submenu {
	_submenu := submenu{}

	_submenu.submenuDo.UseDB(db, opts...)
	_submenu.submenuDo.UseModel(&model.Submenu{})

	tableName := _submenu.submenuDo.TableName()
	_submenu.ALL = field.NewAsterisk(tableName)
	_submenu.ID = field.NewInt32(tableName, "id")
	_submenu.MenuID = field.NewInt32(tableName, "menu_id")
	_submenu.Name = field.NewString(tableName, "name")
	_submenu.PublicName = field.NewString(tableName, "public_name")
	_submenu.Description = field.NewString(tableName, "description")
	_submenu.CreatedBy = field.NewString(tableName, "created_by")
	_submenu.CreatedAt = field.NewTime(tableName, "created_at")
	_submenu.UpdatedBy = field.NewString(tableName, "updated_by")
	_submenu.UpdatedAt = field.NewTime(tableName, "updated_at")
	_submenu.DeletedBy = field.NewString(tableName, "deleted_by")
	_submenu.DeletedAt = field.NewField(tableName, "deleted_at")

	_submenu.fillFieldMap()

	return _submenu
}

type submenu struct {
	submenuDo submenuDo

	ALL         field.Asterisk
	ID          field.Int32
	MenuID      field.Int32
	Name        field.String
	PublicName  field.String
	Description field.String
	CreatedBy   field.String
	CreatedAt   field.Time
	UpdatedBy   field.String
	UpdatedAt   field.Time
	DeletedBy   field.String
	DeletedAt   field.Field

	fieldMap map[string]field.Expr
}

func (s submenu) Table(newTableName string) *submenu {
	s.submenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s submenu) As(alias string) *submenu {
	s.submenuDo.DO = *(s.submenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *submenu) updateTableName(table string) *submenu {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.MenuID = field.NewInt32(table, "menu_id")
	s.Name = field.NewString(table, "name")
	s.PublicName = field.NewString(table, "public_name")
	s.Description = field.NewString(table, "description")
	s.CreatedBy = field.NewString(table, "created_by")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedBy = field.NewString(table, "updated_by")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedBy = field.NewString(table, "deleted_by")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *submenu) WithContext(ctx context.Context) ISubmenuDo { return s.submenuDo.WithContext(ctx) }

func (s submenu) TableName() string { return s.submenuDo.TableName() }

func (s submenu) Alias() string { return s.submenuDo.Alias() }

func (s submenu) Columns(cols ...field.Expr) gen.Columns { return s.submenuDo.Columns(cols...) }

func (s *submenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *submenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["id"] = s.ID
	s.fieldMap["menu_id"] = s.MenuID
	s.fieldMap["name"] = s.Name
	s.fieldMap["public_name"] = s.PublicName
	s.fieldMap["description"] = s.Description
	s.fieldMap["created_by"] = s.CreatedBy
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_by"] = s.UpdatedBy
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_by"] = s.DeletedBy
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s submenu) clone(db *gorm.DB) submenu {
	s.submenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s submenu) replaceDB(db *gorm.DB) submenu {
	s.submenuDo.ReplaceDB(db)
	return s
}

type submenuDo struct{ gen.DO }

type ISubmenuDo interface {
	gen.SubQuery
	Debug() ISubmenuDo
	WithContext(ctx context.Context) ISubmenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISubmenuDo
	WriteDB() ISubmenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISubmenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISubmenuDo
	Not(conds ...gen.Condition) ISubmenuDo
	Or(conds ...gen.Condition) ISubmenuDo
	Select(conds ...field.Expr) ISubmenuDo
	Where(conds ...gen.Condition) ISubmenuDo
	Order(conds ...field.Expr) ISubmenuDo
	Distinct(cols ...field.Expr) ISubmenuDo
	Omit(cols ...field.Expr) ISubmenuDo
	Join(table schema.Tabler, on ...field.Expr) ISubmenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISubmenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISubmenuDo
	Group(cols ...field.Expr) ISubmenuDo
	Having(conds ...gen.Condition) ISubmenuDo
	Limit(limit int) ISubmenuDo
	Offset(offset int) ISubmenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISubmenuDo
	Unscoped() ISubmenuDo
	Create(values ...*model.Submenu) error
	CreateInBatches(values []*model.Submenu, batchSize int) error
	Save(values ...*model.Submenu) error
	First() (*model.Submenu, error)
	Take() (*model.Submenu, error)
	Last() (*model.Submenu, error)
	Find() ([]*model.Submenu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Submenu, err error)
	FindInBatches(result *[]*model.Submenu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Submenu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISubmenuDo
	Assign(attrs ...field.AssignExpr) ISubmenuDo
	Joins(fields ...field.RelationField) ISubmenuDo
	Preload(fields ...field.RelationField) ISubmenuDo
	FirstOrInit() (*model.Submenu, error)
	FirstOrCreate() (*model.Submenu, error)
	FindByPage(offset int, limit int) (result []*model.Submenu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISubmenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s submenuDo) Debug() ISubmenuDo {
	return s.withDO(s.DO.Debug())
}

func (s submenuDo) WithContext(ctx context.Context) ISubmenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s submenuDo) ReadDB() ISubmenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s submenuDo) WriteDB() ISubmenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s submenuDo) Session(config *gorm.Session) ISubmenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s submenuDo) Clauses(conds ...clause.Expression) ISubmenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s submenuDo) Returning(value interface{}, columns ...string) ISubmenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s submenuDo) Not(conds ...gen.Condition) ISubmenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s submenuDo) Or(conds ...gen.Condition) ISubmenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s submenuDo) Select(conds ...field.Expr) ISubmenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s submenuDo) Where(conds ...gen.Condition) ISubmenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s submenuDo) Order(conds ...field.Expr) ISubmenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s submenuDo) Distinct(cols ...field.Expr) ISubmenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s submenuDo) Omit(cols ...field.Expr) ISubmenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s submenuDo) Join(table schema.Tabler, on ...field.Expr) ISubmenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s submenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISubmenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s submenuDo) RightJoin(table schema.Tabler, on ...field.Expr) ISubmenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s submenuDo) Group(cols ...field.Expr) ISubmenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s submenuDo) Having(conds ...gen.Condition) ISubmenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s submenuDo) Limit(limit int) ISubmenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s submenuDo) Offset(offset int) ISubmenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s submenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISubmenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s submenuDo) Unscoped() ISubmenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s submenuDo) Create(values ...*model.Submenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s submenuDo) CreateInBatches(values []*model.Submenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s submenuDo) Save(values ...*model.Submenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s submenuDo) First() (*model.Submenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Submenu), nil
	}
}

func (s submenuDo) Take() (*model.Submenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Submenu), nil
	}
}

func (s submenuDo) Last() (*model.Submenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Submenu), nil
	}
}

func (s submenuDo) Find() ([]*model.Submenu, error) {
	result, err := s.DO.Find()
	return result.([]*model.Submenu), err
}

func (s submenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Submenu, err error) {
	buf := make([]*model.Submenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s submenuDo) FindInBatches(result *[]*model.Submenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s submenuDo) Attrs(attrs ...field.AssignExpr) ISubmenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s submenuDo) Assign(attrs ...field.AssignExpr) ISubmenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s submenuDo) Joins(fields ...field.RelationField) ISubmenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s submenuDo) Preload(fields ...field.RelationField) ISubmenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s submenuDo) FirstOrInit() (*model.Submenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Submenu), nil
	}
}

func (s submenuDo) FirstOrCreate() (*model.Submenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Submenu), nil
	}
}

func (s submenuDo) FindByPage(offset int, limit int) (result []*model.Submenu, count int64, err error) {
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

func (s submenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s submenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s submenuDo) Delete(models ...*model.Submenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *submenuDo) withDO(do gen.Dao) *submenuDo {
	s.DO = *do.(*gen.DO)
	return s
}
