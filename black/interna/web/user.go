package web

import (
	"net/http"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

const emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
const passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`

type UserHandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		emailExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {

	server.GET("/user/login", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "login success")
	})

	server.POST("/user/signup", func(ctx *gin.Context) {

		type SinupStruct struct {
			Email           string `json:"email"`
			Password        string `json:"password"`
			ConfirmPassword string `json:"confirmpassword"`
		}

		var req SinupStruct
		err := ctx.Bind(&req)
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}

		ok, err := u.emailExp.MatchString(req.Email)
		if err != nil {
			ctx.String(http.StatusOK, "系统错误！")
			return
		}
		if !ok {
			ctx.String(http.StatusOK, "邮箱格式错误，请重新填写！")
			return
		}

		if req.Password != req.ConfirmPassword {
			ctx.String(http.StatusOK, "两次密码输入不匹配！")
			return
		}

		ok, err = u.passwordExp.MatchString(req.Password)
		if err != nil {
			ctx.String(http.StatusOK, "系统错误！")
			return
		}
		if !ok {
			ctx.String(http.StatusOK, "密码格式错误，请重新填写！")
			return
		}

		ctx.String(http.StatusOK, "signup success")
	})

	server.GET("/user/profile", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "profile success")
	})

	server.POST("/user/edit", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "edit success")
	})
}

func (u *UserHandler) SignUp(ctx *gin.Context) {

}

func (u *UserHandler) SignIn(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {

}

func (u *UserHandler) Edit(ctx *gin.Context) {

}
