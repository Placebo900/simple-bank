package api

import (
	"database/sql"
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

type UserResponse struct {
	Username string `json:"username" binding:"required,alphanum"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}
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
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				c.JSON(http.StatusForbidden, err.Error())
				return
			}
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	resp := newUserResponse(user)
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
	resp := newUserResponse(user)
	c.JSON(http.StatusOK, resp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"access_token" binding:"required"`
}

func (s *Server) loginUser(c *gin.Context) {
	arg := loginUserRequest{}
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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

	err = util.CheckPassword(arg.Password, user.HashedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	accessToken, err := s.tokenMaker.CreateToken(user.Username, s.config.AccessTokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	res := loginUserResponse{
		User:        newUserResponse(user),
		AccessToken: accessToken,
	}
	c.JSON(http.StatusOK, res)
}
