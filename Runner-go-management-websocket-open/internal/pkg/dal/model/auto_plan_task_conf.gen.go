// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAutoPlanTaskConf = "auto_plan_task_conf"

// AutoPlanTaskConf mapped from table <auto_plan_task_conf>
type AutoPlanTaskConf struct {
	ID               int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 配置ID
	PlanID           string         `gorm:"column:plan_id;not null;default:0" json:"plan_id"`                         // 计划ID
	TeamID           string         `gorm:"column:team_id;not null" json:"team_id"`                                   // 团队ID
	TaskType         int32          `gorm:"column:task_type;not null" json:"task_type"`                               // 任务类型：1-普通模式，2-定时任务
	TaskMode         int32          `gorm:"column:task_mode;not null;default:1" json:"task_mode"`                     // 运行模式：1-按照用例执行
	SceneRunOrder    int32          `gorm:"column:scene_run_order;not null;default:1" json:"scene_run_order"`         // 场景运行次序：1-顺序执行，2-同时执行
	TestCaseRunOrder int32          `gorm:"column:test_case_run_order;not null;default:1" json:"test_case_run_order"` // 用例运行次序：1-顺序执行，2-同时执行
	RunUserID        string         `gorm:"column:run_user_id;not null;default:0" json:"run_user_id"`                 // 运行人用户ID
	CreatedAt        time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`   // 创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`   // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                      // 删除时间
}

// TableName AutoPlanTaskConf's table name
func (*AutoPlanTaskConf) TableName() string {
	return TableNameAutoPlanTaskConf
}
