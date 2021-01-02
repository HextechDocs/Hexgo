package model

import "time"

type Marker struct {
	Id          int64
	Uuid        string    `xorm:"char(36) notnull unique index"`
	Slug        string    `xorm:"varchar(255) notnull unique index"`
	DisplayName string    `xorm:"varchar(255) notnull"`
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
	Version     int32     `xorm:"version"`
}

func GetMarkerBySlug(slug string) (*Marker, error) {
	var marker Marker
	success, err := x.Where("slug = ?", slug).Get(&marker)

	if err != nil {
		return nil, err
	}

	if !success {
		return nil, nil
	}

	return &marker, nil
}
