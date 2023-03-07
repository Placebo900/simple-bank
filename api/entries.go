package api

import (
	"net/http"

	db "github.com/Placebo900/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type getEntryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getEntry(c *gin.Context) {
	arg := getEntryRequest{}
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	entry, err := s.store.GetEntry(c, arg.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, entry)
}

type listEntryRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) listEntry(c *gin.Context) {
	arg := listEntryRequest{}
	if err := c.BindQuery(&arg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	entry, err := s.store.ListEntry(c, db.ListEntryParams{Limit: arg.PageSize, Offset: (arg.PageID - 1) * arg.PageSize})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, entry)
}
