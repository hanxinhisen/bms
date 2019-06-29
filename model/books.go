// Created by Hisen at 2019-06-25.
package model

import "time"

type BookLists struct {
	Books []*Book
}

type Book struct {
	ID          int       `json:"id" column:"id" form:"id"`
	Name        string    `json:"name" db:"name" form:"name" binding:"required"`
	Author      string    `json:"author" db:"author" form:"author" binding:"required"`
	Price       float64   `json:"price" db:"price" form:"price" binding:"required"`
	PublishDate time.Time `json:"publish_date" db:"publish_date" form:"publish_date" binding:"required" time_format:"2006-01-02 15:04:05"`
	CreateAt    time.Time `json:"create_at" db:"create_at" time_format:"2006-01-02 15:04:05"`
	UpdateAt    time.Time `json:"update_at" db:"update_at" time_format:"2006-01-02 15:04:05"`
}
