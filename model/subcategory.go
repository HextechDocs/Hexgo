package model

import "time"

type Subcategory struct {
	Id          int64
	Uuid        string    `xorm:"char(36) notnull unique index"`
	Slug        string    `xorm:"varchar(255) notnull index"`
	DisplayName string    `xorm:"varchar(255) notnull"`
	Category    *Category `xorm:"-"`
	CategoryId  int64
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
	Version     int32     `xorm:"version"`
}

func (sc *Subcategory) GetParentCategory() *Category {
	if sc.Category == nil {
		var category Category
		success, err := x.Where("id = ?", sc.CategoryId).Get(&category)

		if !success || err != nil {
			return nil
		}

		sc.Category = &category
		return &category
	} else {
		return sc.Category
	}
}

func GetSubcategoryByUuid(uuid string) (*Subcategory, error) {
	var subcategory Subcategory
	success, err := x.Where("uuid = ?", uuid).Get(&subcategory)

	if err != nil {
		return nil, err
	}

	if !success {
		return nil, nil
	}

	return &subcategory, nil
}
