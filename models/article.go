package models

import "time"

type Article struct {
	Id          int64
	Uid         string
	Title       string
	Content     string
	Views       int64
	Comments    int64
	CreatedTime time.Time
}

func (a *Article) TableName() string {
	return TableName("article")
}
