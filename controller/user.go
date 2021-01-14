package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"vanwhebin/try-gin-vue/common"
	"vanwhebin/try-gin-vue/dto"
	"vanwhebin/try-gin-vue/model"
	"vanwhebin/try-gin-vue/response"
	"vanwhebin/try-gin-vue/util"
)

func Register(ctx *gin.Context) {
	//name := ctx.PostForm("name")
	//telephone := ctx.PostForm("telephone")
	//password := ctx.PostForm("password")

	// TODO 使用map获取请求参数
	//var requestBodyMap = make(map[string]string)
	//_ = json.NewDecoder(ctx.Request.Body).Decode(&requestBodyMap)

	/*--------TODO 使用结构体的方式接受参数---------------*/
	//var requestUser = model.User{}
	//_ = json.NewDecoder(ctx.Request.Body).Decode(&requestUser)

	/*--------使用GIN的bind方法的方式接受参数---------------*/
	var requestUser = model.User{}
	_ = ctx.Bind(&requestUser)

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "手机号必须为11位")
		return
	}

	if len(password) < 6 || len(password) > 18 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "密码不能少于6位大于18位")
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
		response.Response(ctx, http.StatusUnprocessableEntity, 400, gin.H{}, "该手机号已注册")
		return
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, gin.H{}, "生成密码错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(pwd),
	}
	result := common.DB.Create(&newUser)
	log.Println(result)
	token, _ := common.ReleaseToken(newUser)

	response.Response(ctx, http.StatusCreated, 201, gin.H{"access_token": token}, "创建成功")
	//model.CreateUser(name, telephone, password)
}

func Login(ctx *gin.Context) {
	// 验证表单数据
	var user model.User
	var requestUser = model.User{}
	_ = ctx.Bind(&requestUser)

	telephone := requestUser.Telephone
	password := requestUser.Password

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "手机号必须为11位")
		return
	}

	if len(password) < 6 || len(password) > 18 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "密码不能少于6位大于18位")
		return
	}
	// 查找数据表单数据对比
	common.DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "用户不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "密码或用户错误")
		return
	}
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusServiceUnavailable, 500, gin.H{}, "系统异常")
		log.Printf("generate jwt token error: %s", err.Error())
	}
	response.Success(ctx, gin.H{"access_token": token}, "OK")
	// 返回token
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUseDto(user.(model.User))}, "OK")
}

func Logout(ctx *gin.Context) {

}
