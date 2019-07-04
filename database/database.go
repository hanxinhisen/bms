// Created by Hisen at 2019-06-25.
package database

import (
	"code.hanx.xin/bms/config"
	"code.hanx.xin/bms/logger"
	"code.hanx.xin/bms/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabaseConnect() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.DbConfig.AuthUser, config.DbConfig.AuthPassword, config.DbConfig.IPAddr, config.DbConfig.Port, config.DbConfig.DatabaseName)
	if DB, err = sql.Open("mysql", dsn); err != nil {
		return
	}
	DB.SetMaxOpenConns(2000)
	DB.SetMaxIdleConns(200)
	if err = DB.Ping(); err != nil {
		return
	}
	return

}

func AddBook(book *model.Book) (err error) {
	sqlStr := fmt.Sprintf("insert into %s(name,author,price,publish_date) values (?,?,?,?)", config.DbConfig.TableName)
	if _, err = DB.Exec(sqlStr, book.Name, book.Author, book.Price, book.PublishDate); err != nil {
		return err
	}
	return nil

}

func UpdateBook(book *model.Book) (err error) {
	sqlStr := fmt.Sprintf("update %s set name=?, author=?, price=?, publish_date = ? where id= ?", config.DbConfig.TableName)
	if _, err = DB.Exec(sqlStr, book.Name, book.Author, book.Price, book.PublishDate, book.ID); err != nil {
		return err
	}
	return nil

}

func GetBookInfoByID(id int) (book *model.Book, err error) {
	book = &model.Book{}
	sqlStr := fmt.Sprintf("select id,name,author,price,publish_date,create_at,update_at from %s where id=?", config.DbConfig.TableName)
	if err := DB.QueryRow(sqlStr, id).Scan(&book.ID, &book.Name, &book.Author, &book.Price, &book.PublishDate, &book.CreateAt, &book.UpdateAt); err != nil {
		logger.Logger.Debugf("error -> %v", err)
		return nil, err
	}
	return
}

func GetBookList() (books []*model.Book, err error) {
	books = make([]*model.Book, 0)

	sqlStr := fmt.Sprintf("select id,name,author,price,publish_date,create_at,update_at from %s order by publish_date", config.DbConfig.TableName)
	rows, err := DB.Query(sqlStr)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		book := &model.Book{}
		err = rows.Scan(&book.ID, &book.Name, &book.Author, &book.Price, &book.PublishDate, &book.CreateAt, &book.UpdateAt)
		if err != nil {
			return
		}
		books = append(books, book)
	}
	return
}

func DeleteBookByID(id int) (err error) {
	sqlStr := fmt.Sprintf("delete from %s where id=?", config.DbConfig.TableName)
	ret, err := DB.Exec(sqlStr, id)
	if err != nil {
		return
	}
	i, _ := ret.RowsAffected()

	logger.Logger.Infof("成功删除%d记录,删除%d行", id, i)
	return
}
