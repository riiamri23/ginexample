package requests

type CreatePlatformsRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type UpdatePlatformsRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}
