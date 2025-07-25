// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTempRole = "temp_roles"

// TempRole mapped from table <temp_roles>
type TempRole struct {
	ID           int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RoleID       int32          `gorm:"column:role_id;not null" json:"role_id"`
	Name         string         `gorm:"column:name;not null" json:"name"`
	Description  string         `gorm:"column:description;not null" json:"description"`
	Type         string         `gorm:"column:type;not null" json:"type"`
	InactiveDate *time.Time     `gorm:"column:inactive_date" json:"inactive_date"`
	Resources    *string        `gorm:"column:resources" json:"resources"`
	CreatedBy    string         `gorm:"column:created_by;not null" json:"created_by"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedBy    string         `gorm:"column:updated_by;not null" json:"updated_by"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedBy    *string        `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName TempRole's table name
func (*TempRole) TableName() string {
	return TableNameTempRole
}
