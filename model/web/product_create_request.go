package web

type ProductCreateRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
}