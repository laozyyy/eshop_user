package req

type LoginDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
