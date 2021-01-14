package validate

type CreateCategoryValidate struct {
	Name string `json:"name" binding:"required"`
}
