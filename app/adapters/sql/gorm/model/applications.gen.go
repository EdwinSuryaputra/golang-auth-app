// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameApplication = "applications"

// Application mapped from table <applications>
type Application struct {
	ID          int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string         `gorm:"column:name;not null" json:"name"`
	PublicName  string         `gorm:"column:public_name;not null" json:"public_name"`
	Description string         `gorm:"column:description;not null" json:"description"`
	CreatedBy   string         `gorm:"column:created_by;not null" json:"created_by"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedBy   string         `gorm:"column:updated_by;not null" json:"updated_by"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedBy   *string        `gorm:"column:deleted_by" json:"deleted_by"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Application's table name
func (*Application) TableName() string {
	return TableNameApplication
}
