package handler

import (
	"net/http"

	"go-grpc/client/domain/dto"
	userPB "go-grpc/client/proto"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	client userPB.UserServiceClient
}

func NewUserHandler(route *gin.Engine, client userPB.UserServiceClient) {
	handler := userHandler{client: client}
	route.POST("api/v1/user", handler.CreateUserHandler)
	route.GET("api/v1/user/:id", handler.GetUserHandler)
	route.GET("api/v1/users", handler.GetUsersHandler)
	route.PUT("api/v1/user/:id", handler.UpdateUserHandler)
	route.DELETE("api/v1/user/:id", handler.DeleteUserHandler)
}

func (h *userHandler) CreateUserHandler(ctx *gin.Context) {
	create := new(dto.UserValidator)
	if err := ctx.Bind(create); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.client.CreateUser(ctx, &userPB.CreateUserRequest{
		User: &userPB.User{
			Name:     create.Name,
			Password: create.Password,
			Email:    create.Email,
			Phone:    create.Phone,
		},
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user.User})
}

func (h *userHandler) GetUsersHandler(ctx *gin.Context) {
	user, err := h.client.ListUsers(ctx, &userPB.ListUsersRequest{
		To:   1,
		From: 10,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) GetUserHandler(ctx *gin.Context) {
	userParams := new(dto.GetUserId)
	if err := ctx.BindUri(userParams); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.client.GetUser(ctx, &userPB.GetUserRequest{Id: userParams.ID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) UpdateUserHandler(ctx *gin.Context) {
	userParams := new(dto.GetUserId)
	if err := ctx.BindUri(userParams); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userReq := new(dto.UserValidator)
	if err := ctx.Bind(userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.client.UpdateUser(ctx, &userPB.UpdateUserRequest{
		Id: userParams.ID,
		User: &userPB.User{
			Name:     userReq.Name,
			Password: userReq.Password,
			Email:    userReq.Email,
			Phone:    userReq.Phone,
		},
	})

	userDetail, err := h.client.GetUser(ctx, &userPB.GetUserRequest{Id: userParams.ID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, userDetail)
}

func (h *userHandler) DeleteUserHandler(ctx *gin.Context) {
	userParams := new(dto.GetUserId)
	if err := ctx.BindUri(userParams); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.client.DeleteUser(ctx, &userPB.DeleteUserRequest{Id: userParams.ID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}
