package dto

//UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
	Nik      string `json:"nik" form:"nik" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Role    string `json:"role" form:"role" binding:"required"`
}

//UserCreateDTO is used by client when create user
// type UserCreateDTO struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password string `json:"password, omitempty" form:"password, omitempty" validate:"min:6" binding:"required" `
// }
