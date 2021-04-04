package model

import "github.com/gogf/gf/os/gtime"

// Project .
type Project struct {
	ID          string      `orm:"id"`
	CreatedAt   *gtime.Time `orm:"created_at"`
	UpdatedAt   *gtime.Time `orm:"updated_at"`
	DeletedAt   *gtime.Time `orm:"deleted_at"`
	CreatedBy   string      `orm:"created_by"`
	UpdatedBy   string      `orm:"updated_by"`
	Title       string      `orm:"title"`
	Name        string      `orm:"name"`
	Description string      `orm:"description"`
	Cover       string      `orm:"cover"`
}
