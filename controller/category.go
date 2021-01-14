package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vanwhebin/try-gin-vue/model"
	"vanwhebin/try-gin-vue/repo"
	"vanwhebin/try-gin-vue/response"
	"vanwhebin/try-gin-vue/validate"
)

type ICategory interface {
	Rest
}

type Category struct {
	repo repo.CategoryRepository
}

func (c Category) Create(ctx *gin.Context) {
	var requestCategory = validate.CreateCategoryValidate{}
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, gin.H{"params": requestCategory.Name}, "分类名不能为空")
		return
	}
	//category := model.Category{Name: requestCategory.Name}
	//c.DB.Create(&category)
	category, err := c.repo.Create(requestCategory.Name)
	if err != nil {
		panic(err)
		//response.Fail(ctx, gin.H{"params": requestCategory.Name}, "创建失败")
		//return
	}

	response.Response(ctx, http.StatusCreated, 201, gin.H{"category": category}, "创建成功")

}

func (c Category) Update(ctx *gin.Context) {
	var requestCategory = validate.CreateCategoryValidate{}
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, gin.H{"params": requestCategory}, err.Error())
		return
	}
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//var updateCategory model.Category
	//if err := c.DB.First(&updateCategory, categoryId).Error; err != nil {
	//	response.Fail(ctx, gin.H{}, "数据错误")
	//	return
	//}
	//c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	updateCategory, err := c.repo.SelectByID(categoryId)
	if err != nil {
		response.Fail(ctx, gin.H{}, "数据错误")
		return
	}
	updateCategory, err = c.repo.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		response.Fail(ctx, gin.H{}, "更新失败")
		return
	}

	response.Success(ctx, gin.H{"category": updateCategory}, "更新成功")
}

func (c Category) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	//var findCategory model.Category
	//if err := c.DB.First(&findCategory, categoryId).Error; err != nil {
	//	response.Fail(ctx, gin.H{}, "数据错误")
	//	return
	//}
	findCategory, err := c.repo.SelectByID(categoryId)
	if err != nil {
		response.Fail(ctx, gin.H{}, "数据错误")
		return
	}

	response.Success(ctx, gin.H{"category": findCategory}, "OK")
}

func (c Category) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	//if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
	//	response.Fail(ctx, gin.H{}, "删除失败")
	//	return
	//}

	err := c.repo.DeleteByID(categoryId)
	if err != nil {
		//response.Fail(ctx, gin.H{}, "删除失败")
		panic(err)
		return
	}

	response.Response(ctx, http.StatusNoContent, 204, nil, "Deleted")
}

func NewCategory() ICategory {
	category := repo.NewCategoryRepository()
	_ = category.DB.AutoMigrate(&model.Category{})
	return Category{repo: category}
}


func Create(ctx *gin.Context) {
	// 数据验证

	// 执行数据库IO操作  简单或者复杂

	// 返回http响应

	// 统一的异常处理  中间件  在理解http响应请求流程就能完成web框架的编写， 在solid原则基础上完善

}
