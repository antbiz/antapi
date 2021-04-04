package model

// Project .
type Project struct {
	ID          string `orm:"id"`
	Title       string `orm:"title"`
	Name        string `orm:"name"`
	Description string `orm:"description"`
	Cover       string `orm:"cover"`
	IsPublic    bool   `orm:"is_public"`
}
