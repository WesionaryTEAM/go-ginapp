package service

import "go-ginapp/auth"

// it will be defined in an interface:
type sigInInterface interface {
	SignIn(auth.AuthDetails) (string, error)
}

type signInStruct struct{}

//let expose this interface:
var (
	Authorize sigInInterface = &signInStruct{}
)

func (si *signInStruct) SignIn(authD auth.AuthDetails) (string, error) {
	token, err := auth.CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token, nil
}
