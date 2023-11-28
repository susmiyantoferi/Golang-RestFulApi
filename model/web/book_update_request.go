package web

type BookUpdateRequest struct {
	Id      int    `validate:"required"`
	Title   string `validate:"required,min=1,max=255" json:"title"`
	Author  string `validate:"required,min=1,max=255" json:"author"`
	Descrip string `validate:"required,min=1,max=255" json:"descrip"`
}
