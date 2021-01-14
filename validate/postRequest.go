package validate

type CreatePostValidate struct {
	Title      string `json:"title" binding:"required,max=30"`
	CategoryId uint   `json:"category_id" binding:"required"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" binding:"required"`
}
