package items

type CreateItemSchema struct {
	Name string `form:"name" json:"name" binding:"required" validate:"required,min=4,max=32"`
}
