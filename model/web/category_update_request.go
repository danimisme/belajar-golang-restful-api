package web

type CategoryUpdateRequest struct {
	Id int `validate:"required" json:"-"`
	Name string `validate:"requeid,max=200,min=1" json:"name"`
}