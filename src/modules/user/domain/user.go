package domain

type UserFilter struct {
	Name     string
	Username string
}

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}

type RegisterPayload struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
