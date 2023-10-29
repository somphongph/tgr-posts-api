package posts

type cached interface {
	GetCache(string) (string, error)
	SetCache(string, interface{}, int) error
	SetShortCache(string, interface{}) error
	SetLongCache(string, interface{}) error
}

type Handler struct {
	cache cached
}

func NewHandler(cache cached) *Handler {
	return &Handler{cache: cache}
}

type Post struct {
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	SubModel string `json:"subModel"`
}
