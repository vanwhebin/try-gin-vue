package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"vanwhebin/try-gin-vue/model"
	"vanwhebin/try-gin-vue/repo"
	"vanwhebin/try-gin-vue/response"
	"vanwhebin/try-gin-vue/validate"
)

type IPost interface {
	Rest
}

type Post struct {
	repo repo.PostRepository
}

func (p Post) Create(ctx *gin.Context) {
	var request = validate.CreatePostValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, gin.H{"params": request}, err.Error())
		return
	}
	// 获取登录用户 user
	user, _ := ctx.Get("user")

	// 创建post
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: request.CategoryId,
		Title:      request.Title,
		HeadImg:    request.HeadImg,
		Content:    request.Content,
	}

	//post, err := p.repo.Create(post)
	//if err != nil {
	//	panic(err)
	//	return
	//}

	response.Success(ctx, gin.H{"post": post}, "创建成功")
}

func (p Post) Update(ctx *gin.Context) {
	var request = validate.CreatePostValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, gin.H{"params": request}, err.Error())
		return
	}
}

func (p Post) Show(ctx *gin.Context) {
	panic("implement me")
}

func (p Post) Delete(ctx *gin.Context) {
	panic("implement me")
}

func NewPost() IPost {
	r := repo.NewPostRepository()
	_ = r.DB.AutoMigrate(&model.Post{})
	return Post{repo: r}
}
