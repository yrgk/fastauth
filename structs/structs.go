package structs

// You can extend it
type (
	RegisterRequest struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	LoginRequest struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
)
