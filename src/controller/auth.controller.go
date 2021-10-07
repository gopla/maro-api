package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gopla/maro/src/helper"
	"github.com/gopla/maro/src/model"
	"github.com/gopla/maro/src/service"
	"golang.org/x/crypto/bcrypt"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	loginService service.LoginService
	jwtService  service.JWTService
}

func NewAuthController(loginService service.LoginService, jwtService service.JWTService) AuthController {
	return &authController{
		loginService: loginService,
		jwtService:  jwtService,
	}
}

func VerifyCredential(username string, password string) interface{} {
	res := CompareUser(username, password)
	if v, ok := res.(model.User); ok {
		comparedPass := ComparePassword(v.Password, []byte(password))
		if v.Username == username && comparedPass{
			return res
		}
		return false
	}
	return false
}

func CompareUser(username string, password string) interface{} {
	var user model.User
	res := model.DB.Where("username = ?", username).Take(&user)
	if res.Error == nil {
		return user
	}

	return nil
}

func ComparePassword(hashedPass string, plainPass []byte) bool  {
	byteHash := []byte(hashedPass)
	err := bcrypt.CompareHashAndPassword(byteHash,plainPass)
	if err != nil {
		log.Print(err)
		return false
	}

	return true
}

func DuplicateUsername(username string) bool {
	res := CompareUser(username, "")
	if v, ok := res.(model.User); ok {
		return v.Username != username
	}
	return true
}

func (a *authController)Login(ctx *gin.Context){
	var input model.Login
	if err := ctx.ShouldBindJSON(&input); err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
    ctx.JSON(http.StatusBadRequest, res)
    return
  }

	authResult := VerifyCredential(input.Username,input.Password)
	if v, ok := authResult.(model.User); ok {
		generatedToken := a.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID),10))
		v.Token = generatedToken

		response := helper.BuildResponse(true, "Logged In",v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.BuildErrorResponse("Invalid",helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (a *authController)Register(ctx *gin.Context)  {
	var input model.Register
	if err := ctx.ShouldBindJSON(&input); err != nil {
		res := helper.BuildErrorResponse(err.Error(),err)
    ctx.JSON(http.StatusBadRequest, res)
    return
  }

	dupe := DuplicateUsername(input.Username)
	

	if !dupe{
		resp := helper.BuildErrorResponse("Duplicate Username",helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict,resp)
	}else{
		user := model.User{Name: input.Name, Username: input.Username, Password: helper.HashAndSalt([]byte(input.Password))}
		model.DB.Create(&user)

		token:=a.jwtService.GenerateToken(strconv.FormatUint(uint64(user.ID),10))
		user.Token = token
		response := helper.BuildResponse(true, "Logged In",user)
		ctx.JSON(http.StatusOK, response)
	}
}
