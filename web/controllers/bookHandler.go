// Created by Hisen at 2019-06-25.
package controllers

import (
	"code.hanx.xin/bms/database"
	"code.hanx.xin/bms/logger"
	"code.hanx.xin/bms/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBookByIDHandlers(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		SuccessResponseHandle(c, err.Error(), nil)
		return
	}

	book, err := database.GetBookInfoByID(idInt)
	if err != nil {
		ErrorResponseHandle(c, http.StatusNotFound, err.Error())
		return
	}

	var books []*model.Book
	books = append(books, book)
	// 纯属测试无实际意义,后续可根据trace_id记录和过滤日志
	traceId, ok := c.Get("trace_id")
	if !ok {
		traceId = "xxxx"
	}
	logger.Logger.Debugf("trace_id:%v", traceId)
	SuccessResponseHandle(c, "success", books)
}

func GetBooksListHandlers(c *gin.Context) {
	books, err := database.GetBookList()
	if err != nil {
		ErrorResponseHandle(c, http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponseHandle(c, "success", books)
}

func PostBooksListHandlers(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBind(&book); err != nil {
		logger.Logger.Errorf("解析参数出错,%v", err)
		ErrorResponseHandle(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := database.AddBook(&book); err != nil {
		ErrorResponseHandle(c, http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponseHandle(c, "success", nil)
}

func PutBooksListHandlers(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBind(&book); err != nil {
		logger.Logger.Errorf("解析参数出错,%v", err)
		ErrorResponseHandle(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := database.UpdateBook(&book); err != nil {
		ErrorResponseHandle(c, http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponseHandle(c, "success", nil)
}

func DeleteBooksListHandlers(c *gin.Context) {
	id := c.Param("id")
	// TODO 强转成int,防止sql注入,后续判断用户输入
	idInt, err := strconv.Atoi(id)

	if err != nil {
		ErrorResponseHandle(c, http.StatusOK, err.Error())
		return
	}
	err = database.DeleteBookByID(idInt)
	if err != nil {
		ErrorResponseHandle(c, http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponseHandle(c, "success", nil)
}
