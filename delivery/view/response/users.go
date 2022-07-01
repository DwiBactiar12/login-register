package response

import "time"

type Admin struct {
	ID      string    `json:"id"`
	Nama      string    `json:"nama"`
	Email     string    `json:"email"`
	HP        string    `json:"hp"`
	Role      string    `json:"role"`
	Umur      int       `json:"umur"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
}
type User struct {
	ID      string    `json:"id"`
	Email     string    `json:"email"`
	HP        string    `json:"hp"`
	Umur      int       `json:"umur"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
}

type Login struct {
	ID    uint   `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

type LoginDetail struct {
	User  Login  `json:"user"`
	Token string `json:"token"`
}
