package secrets

import(
	"errors"
)
type SecretRepo interface {
	 GetSecret(name string) (string, error)
}

type inMemRepo struct {

}

func (f inMemRepo) GetSecret(name string) (string, error) {
	if name == "stew" {
		return "secret",nil
	}
	return "", errors.New("user not found")
}

func NewInMemRepo() SecretRepo {
	return &inMemRepo{}
}