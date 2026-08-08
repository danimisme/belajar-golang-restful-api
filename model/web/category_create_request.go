package web

type CategoryCreateRequest struct {
	Name string `validate:"requeid,max=200,min=1"`
}