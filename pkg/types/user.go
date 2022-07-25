package types

type User struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Type      string `json:"type" validate:"required"`
}


type Token struct {
	User_Token        string
	User_Refreshtoken string
}
