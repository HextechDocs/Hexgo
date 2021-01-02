/*
  Package model defines the GraphQL and ORM models used throughout the application
 */
package model

import "time"

type Author struct {
	Id                int64
	GithubId          int64     `xorm:"notnull unique index"`
	GithubUsername    string    `xorm:"VARCHAR(50) notnull"`
	GithubName        string    `xorm:"VARCHAR(255)"`
	ShouldDisplayName bool      `xorm:"notnull default false"`
	CreatedAt         time.Time `xorm:"created"`
	UpdatedAt         time.Time `xorm:"updated"`
	Version           int32     `xorm:"version"`
}
