package request

type Register struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ValidateRegister struct {
	Firstname string `validate:"required,min=3,max=50"`
	Lastname string `validate:"required,min=3,max=50"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=50"`
}

type ValidateLogin struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=50"`
}