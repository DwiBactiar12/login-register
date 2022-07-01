package request

type InsertUser struct {
	Nama     string `json:"nama" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	HP       string `json:"hp" validate:"required"`
	Umur     int    `json:"umur" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
}

type UpdateUser struct {
	Name     string `json:"nama" validate:"required"`
	Password string `json:"password"`
	HP       string `json:"hp"`
	Umur     int    `json:"umur" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
