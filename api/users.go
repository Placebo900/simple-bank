package api

import (
	"database/sql"
	"log"
	"net/http"

	db "github.com/Placebo900/simple-bank/db/sqlc"
	"github.com/Placebo900/simple-bank/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username string `json:"username" binding:"required,alphanum"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (s *Server) createUser(c *gin.Context) {
	arg := createUserRequest{}
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	password, err := util.HashPassword(arg.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := s.store.CreateUser(c,
		db.CreateUserParams{
			Username:       arg.Username,
			HashedPassword: password,
			FullName:       arg.FullName,
			Email:          arg.Email,
		})
	if err != nil {
		log.Println(err.(*pq.Error).Code.Name())
		switch err.(*pq.Error).Code.Name() {
		case "unique_violation":
			c.JSON(http.StatusForbidden, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	resp := createUserResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}
	c.JSON(http.StatusOK, resp)
}

type getUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
}

func (s *Server) getUser(c *gin.Context) {
	arg := getUserRequest{}
	if err := c.ShouldBindUri(&arg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := s.store.GetUser(c, arg.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	resp := createUserResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}
	c.JSON(http.StatusOK, resp)
}
