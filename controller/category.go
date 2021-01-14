package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"vanwhebin/try-gin-vue/common"
	"vanwhebin/try-gin-vue/model"
	"vanwhebin/try-gin-vue/response"
)

type ICategory interface {
	Rest
}

type Category struct {
	DB *gorm.DB
}

func (c Category) Create(ctx *gin.Context) {
	var requestCategory = model.Category{}
	_ = ctx.Bind(&requestCategory)

	name := requestCategory.Name
	if name == "" {
		response.Fail(ctx, gin.H{"params": requestCategory.Name}, "分类名不能为空")
		return
	}
	c.DB.Create(&requestCategory)
	response.Response(ctx, http.StatusCreated, 201, gin.H{"category": requestCategory}, "创建成功")

}

func (c Category) Update(ctx *gin.Context) {
	var requestCategory = model.Category{}
	_ = ctx.Bind(&requestCategory)

	name := requestCategory.Name
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory model.Category
	if err := c.DB.First(&updateCategory, categoryId).Error; err != nil {
		response.Fail(ctx, gin.H{}, "数据错误")
		return
	}
	c.DB.Model(&updateCategory).Update("name", name)
	response.Success(ctx, gin.H{"category": updateCategory}, "更新成功")
}

func (c Category) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var findCategory model.Category
	if err := c.DB.First(&findCategory, categoryId).Error; err != nil {
		response.Fail(ctx, gin.H{}, "数据错误")
		return
	}
	response.Success(ctx, gin.H{"category": findCategory}, "OK")
}

func (c Category) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, gin.H{}, "删除失败")
		return
	}

	response.Response(ctx, http.StatusNoContent, 204, nil, "Deleted")
}

func NewCategory() ICategory {
	db := common.GetDB()
	_ = db.AutoMigrate(&model.Category{})
	return Category{DB: db}
}

// 分类的增删改查

// 控制器方法  增删改查
// 资源控制器如何写
// 要有一个基类，实现基类的方法
// go中是通过接口的组合

func Create(ctx *gin.Context) {
	// 数据验证

	// 执行数据库IO操作  简单或者复杂

	// 返回http响应

	// 统一的异常处理  中间件  在理解http响应请求流程就能完成web框架的编写， 在solid原则基础上完善

}
