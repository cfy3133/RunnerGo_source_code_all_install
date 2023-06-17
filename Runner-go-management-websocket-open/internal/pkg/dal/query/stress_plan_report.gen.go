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

	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/model"
)

func newStressPlanReport(db *gorm.DB, opts ...gen.DOOption) stressPlanReport {
	_stressPlanReport := stressPlanReport{}

	_stressPlanReport.stressPlanReportDo.UseDB(db, opts...)
	_stressPlanReport.stressPlanReportDo.UseModel(&model.StressPlanReport{})

	tableName := _stressPlanReport.stressPlanReportDo.TableName()
	_stressPlanReport.ALL = field.NewAsterisk(tableName)
	_stressPlanReport.ID = field.NewInt64(tableName, "id")
	_stressPlanReport.ReportID = field.NewString(tableName, "report_id")
	_stressPlanReport.ReportName = field.NewString(tableName, "report_name")
	_stressPlanReport.TeamID = field.NewString(tableName, "team_id")
	_stressPlanReport.PlanID = field.NewString(tableName, "plan_id")
	_stressPlanReport.RankID = field.NewInt64(tableName, "rank_id")
	_stressPlanReport.PlanName = field.NewString(tableName, "plan_name")
	_stressPlanReport.SceneID = field.NewString(tableName, "scene_id")
	_stressPlanReport.SceneName = field.NewString(tableName, "scene_name")
	_stressPlanReport.TaskType = field.NewInt32(tableName, "task_type")
	_stressPlanReport.TaskMode = field.NewInt32(tableName, "task_mode")
	_stressPlanReport.ControlMode = field.NewInt32(tableName, "control_mode")
	_stressPlanReport.DebugMode = field.NewString(tableName, "debug_mode")
	_stressPlanReport.Status = field.NewInt32(tableName, "status")
	_stressPlanReport.Remark = field.NewString(tableName, "remark")
	_stressPlanReport.RunUserID = field.NewString(tableName, "run_user_id")
	_stressPlanReport.CreatedAt = field.NewTime(tableName, "created_at")
	_stressPlanReport.UpdatedAt = field.NewTime(tableName, "updated_at")
	_stressPlanReport.DeletedAt = field.NewField(tableName, "deleted_at")

	_stressPlanReport.fillFieldMap()

	return _stressPlanReport
}

type stressPlanReport struct {
	stressPlanReportDo stressPlanReportDo

	ALL         field.Asterisk
	ID          field.Int64
	ReportID    field.String // 报告ID
	ReportName  field.String // 报告名称
	TeamID      field.String // 团队ID
	PlanID      field.String // 计划ID
	RankID      field.Int64  // 序号ID
	PlanName    field.String // 计划名称
	SceneID     field.String // 场景ID
	SceneName   field.String // 场景名称
	TaskType    field.Int32  // 任务类型
	TaskMode    field.Int32  // 压测模式
	ControlMode field.Int32  // 控制模式：0-集中模式，1-单独模式
	DebugMode   field.String // debug模式：stop-关闭，all-开启全部日志，only_success-开启仅成功日志，only_error-开启仅错误日志
	Status      field.Int32  // 报告状态1:进行中，2:已完成
	Remark      field.String // 备注
	RunUserID   field.String // 启动人id
	CreatedAt   field.Time   // 创建时间（执行时间）
	UpdatedAt   field.Time   // 修改时间
	DeletedAt   field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (s stressPlanReport) Table(newTableName string) *stressPlanReport {
	s.stressPlanReportDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s stressPlanReport) As(alias string) *stressPlanReport {
	s.stressPlanReportDo.DO = *(s.stressPlanReportDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *stressPlanReport) updateTableName(table string) *stressPlanReport {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.ReportID = field.NewString(table, "report_id")
	s.ReportName = field.NewString(table, "report_name")
	s.TeamID = field.NewString(table, "team_id")
	s.PlanID = field.NewString(table, "plan_id")
	s.RankID = field.NewInt64(table, "rank_id")
	s.PlanName = field.NewString(table, "plan_name")
	s.SceneID = field.NewString(table, "scene_id")
	s.SceneName = field.NewString(table, "scene_name")
	s.TaskType = field.NewInt32(table, "task_type")
	s.TaskMode = field.NewInt32(table, "task_mode")
	s.ControlMode = field.NewInt32(table, "control_mode")
	s.DebugMode = field.NewString(table, "debug_mode")
	s.Status = field.NewInt32(table, "status")
	s.Remark = field.NewString(table, "remark")
	s.RunUserID = field.NewString(table, "run_user_id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *stressPlanReport) WithContext(ctx context.Context) *stressPlanReportDo {
	return s.stressPlanReportDo.WithContext(ctx)
}

func (s stressPlanReport) TableName() string { return s.stressPlanReportDo.TableName() }

func (s stressPlanReport) Alias() string { return s.stressPlanReportDo.Alias() }

func (s *stressPlanReport) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *stressPlanReport) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 19)
	s.fieldMap["id"] = s.ID
	s.fieldMap["report_id"] = s.ReportID
	s.fieldMap["report_name"] = s.ReportName
	s.fieldMap["team_id"] = s.TeamID
	s.fieldMap["plan_id"] = s.PlanID
	s.fieldMap["rank_id"] = s.RankID
	s.fieldMap["plan_name"] = s.PlanName
	s.fieldMap["scene_id"] = s.SceneID
	s.fieldMap["scene_name"] = s.SceneName
	s.fieldMap["task_type"] = s.TaskType
	s.fieldMap["task_mode"] = s.TaskMode
	s.fieldMap["control_mode"] = s.ControlMode
	s.fieldMap["debug_mode"] = s.DebugMode
	s.fieldMap["status"] = s.Status
	s.fieldMap["remark"] = s.Remark
	s.fieldMap["run_user_id"] = s.RunUserID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s stressPlanReport) clone(db *gorm.DB) stressPlanReport {
	s.stressPlanReportDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s stressPlanReport) replaceDB(db *gorm.DB) stressPlanReport {
	s.stressPlanReportDo.ReplaceDB(db)
	return s
}

type stressPlanReportDo struct{ gen.DO }

func (s stressPlanReportDo) Debug() *stressPlanReportDo {
	return s.withDO(s.DO.Debug())
}

func (s stressPlanReportDo) WithContext(ctx context.Context) *stressPlanReportDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s stressPlanReportDo) ReadDB() *stressPlanReportDo {
	return s.Clauses(dbresolver.Read)
}

func (s stressPlanReportDo) WriteDB() *stressPlanReportDo {
	return s.Clauses(dbresolver.Write)
}

func (s stressPlanReportDo) Session(config *gorm.Session) *stressPlanReportDo {
	return s.withDO(s.DO.Session(config))
}

func (s stressPlanReportDo) Clauses(conds ...clause.Expression) *stressPlanReportDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s stressPlanReportDo) Returning(value interface{}, columns ...string) *stressPlanReportDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s stressPlanReportDo) Not(conds ...gen.Condition) *stressPlanReportDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s stressPlanReportDo) Or(conds ...gen.Condition) *stressPlanReportDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s stressPlanReportDo) Select(conds ...field.Expr) *stressPlanReportDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s stressPlanReportDo) Where(conds ...gen.Condition) *stressPlanReportDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s stressPlanReportDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *stressPlanReportDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s stressPlanReportDo) Order(conds ...field.Expr) *stressPlanReportDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s stressPlanReportDo) Distinct(cols ...field.Expr) *stressPlanReportDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s stressPlanReportDo) Omit(cols ...field.Expr) *stressPlanReportDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s stressPlanReportDo) Join(table schema.Tabler, on ...field.Expr) *stressPlanReportDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s stressPlanReportDo) LeftJoin(table schema.Tabler, on ...field.Expr) *stressPlanReportDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s stressPlanReportDo) RightJoin(table schema.Tabler, on ...field.Expr) *stressPlanReportDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s stressPlanReportDo) Group(cols ...field.Expr) *stressPlanReportDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s stressPlanReportDo) Having(conds ...gen.Condition) *stressPlanReportDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s stressPlanReportDo) Limit(limit int) *stressPlanReportDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s stressPlanReportDo) Offset(offset int) *stressPlanReportDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s stressPlanReportDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *stressPlanReportDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s stressPlanReportDo) Unscoped() *stressPlanReportDo {
	return s.withDO(s.DO.Unscoped())
}

func (s stressPlanReportDo) Create(values ...*model.StressPlanReport) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s stressPlanReportDo) CreateInBatches(values []*model.StressPlanReport, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s stressPlanReportDo) Save(values ...*model.StressPlanReport) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s stressPlanReportDo) First() (*model.StressPlanReport, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanReport), nil
	}
}

func (s stressPlanReportDo) Take() (*model.StressPlanReport, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanReport), nil
	}
}

func (s stressPlanReportDo) Last() (*model.StressPlanReport, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanReport), nil
	}
}

func (s stressPlanReportDo) Find() ([]*model.StressPlanReport, error) {
	result, err := s.DO.Find()
	return result.([]*model.StressPlanReport), err
}

func (s stressPlanReportDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StressPlanReport, err error) {
	buf := make([]*model.StressPlanReport, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s stressPlanReportDo) FindInBatches(result *[]*model.StressPlanReport, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s stressPlanReportDo) Attrs(attrs ...field.AssignExpr) *stressPlanReportDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s stressPlanReportDo) Assign(attrs ...field.AssignExpr) *stressPlanReportDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s stressPlanReportDo) Joins(fields ...field.RelationField) *stressPlanReportDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s stressPlanReportDo) Preload(fields ...field.RelationField) *stressPlanReportDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s stressPlanReportDo) FirstOrInit() (*model.StressPlanReport, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanReport), nil
	}
}

func (s stressPlanReportDo) FirstOrCreate() (*model.StressPlanReport, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StressPlanReport), nil
	}
}

func (s stressPlanReportDo) FindByPage(offset int, limit int) (result []*model.StressPlanReport, count int64, err error) {
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

func (s stressPlanReportDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s stressPlanReportDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s stressPlanReportDo) Delete(models ...*model.StressPlanReport) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *stressPlanReportDo) withDO(do gen.Dao) *stressPlanReportDo {
	s.DO = *do.(*gen.DO)
	return s
}
