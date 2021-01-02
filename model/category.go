package model

import (
	"time"
)

type Category struct {
	Id            int64
	Uuid          string        `xorm:"char(36) notnull unique index"`
	Slug          string        `xorm:"varchar(255) notnull unique index"`
	DisplayName   string        `xorm:"varchar(255) notnull"`
	LogoUrl       string        `xorm:"varchar(500) notnull"`
	ReadmeUrl     string        `xorm:"varchar(500)"`
	Subcategories []Subcategory `xorm:"-"`
	CreatedAt     time.Time     `xorm:"created"`
	UpdatedAt     time.Time     `xorm:"updated"`
	Version       int32         `xorm:"version"`
}

func (c *Category) GetSubcategories() []Subcategory {
	if c.Subcategories == nil {
		var subcategories []Subcategory
		err := x.Where("category_id = ?", c.Id).Find(&subcategories)

		if err != nil {
			emptySubcat := make([]Subcategory, 0)
			c.Subcategories = emptySubcat
			return emptySubcat
		} else {
			c.Subcategories = subcategories
			return subcategories
		}
	} else {
		return c.Subcategories
	}
}

func GetCategories() ([]Category, error) {
	var categories []Category
	err := x.Find(&categories)

	if err != nil {
		return nil, err
	}

	return categories, nil
}
