package api

import (
	"fmt"
	"myblog/middleware"
	"myblog/model"
	"myblog/utils/errmsg"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	username, _ := c.GetPostForm("username") //从表单获得
	password, _ := c.GetPostForm("password")
	//_ = c.ShouldBindJSON(&user)
	user.Username = username
	user.Password = password
	fmt.Printf("user: %v\n", user)
	var token string
	code := model.CheckLogin(user.Username, user.Password)
	if code == errmsg.SUCCESS {
		setToken(c, user.Username)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    user.Username,
			"message": errmsg.GetErrMsg(code),
			"token":   token,
		})
	}
}

// token生成函数
func setToken(c *gin.Context, username string) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "Zyj",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    username,
		"message": errmsg.GetErrMsg(200),
		"token":   token,
	})
	return
}
