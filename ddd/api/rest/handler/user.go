package handler

import (
	"net/http"

	"ddd/api/rest/dto"
	"ddd/application"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userSvc *application.UserService
}

func NewUserHandler(userSvc *application.UserService) *UserHandler {
	return &UserHandler{userSvc: userSvc}
}

// AddUser add a user to the system
// @Summary add a user to the system
// @description add a user to the system so that he will have an identity
// @tags user
// @id AddUser
// @accept application/json
// @produce application/json
// @param data body dto.AddUserReq false "AddUserReq"
// @success 200 {object} dto.JSONResult{}
// @router /ddd/api/v1/user [post]
func (d *UserHandler) AddUser(c *gin.Context) {

	// get request parameters
	var req dto.AddUserReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, &dto.JSONResult{Code: 1, Message: "failed"})
		return
	}

	// deal request in application layer
	err := d.userSvc.AddUser(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.JSONResult{Code: 2, Message: "failed"})
		return
	}

	// response to the request
	c.JSON(http.StatusOK, &dto.JSONResult{Message: "succeed"})
}

// GetUsers find users
// @Summary find users
// @description find users from the given condition
// @tags user
// @id GetUsers
// @accept application/json
// @produce application/json
// @param data query dto.GetUserReq false "get user req"
// @success 200 {object} dto.JSONResult{data=dto.GetUserResp}
// @router /ddd/api/v1/users [get]
func (d *UserHandler) GetUsers(c *gin.Context) {

	// get request parameters
	var req dto.GetUserReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, &dto.JSONResult{Code: 1, Message: "failed"})
		return
	}

	// deal request in application layer
	resp, err := d.userSvc.GetUsers(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.JSONResult{Code: 2, Message: "failed"})
		return
	}

	// response to the request
	c.JSON(http.StatusOK, &dto.JSONResult{Message: "succeed", Data: resp})
}
