package web

type CategoryUpdateRequest struct {
	Id int `validate:"required"`
	Name string `validate:"requeid,max=200,min=1"`
}