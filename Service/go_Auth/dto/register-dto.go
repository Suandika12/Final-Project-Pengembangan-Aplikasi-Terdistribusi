package dto

//RegisterDTO used when client post from /register url
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
	Nik      string `json:"nik" form:"nik" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Role    string `json:"role" form:"role" binding:"required"`

}
