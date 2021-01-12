package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"vanwhebin/try-gin-vue/common"
	"vanwhebin/try-gin-vue/model"
	"vanwhebin/try-gin-vue/util"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 || len(password) > 18 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位大于18位"})
		return
	}
	// 如果用户名没有传则创建一个随机字符串
	if name == "" {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	//ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "创建成功"})
	// 用户已存在
	if model.IsTelephoneExists(common.DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 400, "msg": "该手机号已注册"})
		return
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "生成密码错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(pwd),
	}
	result := common.DB.Create(&newUser)
	fmt.Println(result)

	//model.CreateUser(name, telephone, password)
	ctx.JSON(http.StatusCreated, gin.H{"code": 201, "msg": "创建成功"})
}

func Login(ctx *gin.Context) {
	// 验证表单数据
	var user model.User
	//name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 || len(password) > 18 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位大于18位"})
		return
	}
	// 查找数据表单数据对比
	common.DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusExpectationFailed, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	if 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password));err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{"code": 422, "msg": "密码或用户错误"})
		return
	}
	token := "34132412412"
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": token})

	// 返回token
}
