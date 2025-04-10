package usecase

type usecase struct {
	jwt  jwt
	repo repository
}

func New(jwt jwt, repo repository) *usecase {
	return &usecase{
		jwt:  jwt,
		repo: repo,
	}
}
