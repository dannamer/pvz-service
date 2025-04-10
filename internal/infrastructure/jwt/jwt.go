package jwt

type jwt struct {
	key string
}

func New(key string) *jwt {
	return &jwt{}
}
