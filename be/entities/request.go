package entities

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
