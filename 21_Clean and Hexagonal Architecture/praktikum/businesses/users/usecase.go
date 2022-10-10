package users

import "praktikum/app/middlewares"

type UserUsecase struct {
	userRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		userRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *UserUsecase) CreateUser(userDomain *Domain) Domain {
	return uu.userRepository.CreateUser(userDomain)
}

func (uu *UserUsecase) Login(userDomain *Domain) string {
	user := uu.userRepository.GetByEmail(userDomain)

	if user.ID == 0 {
		return ""
	}

	token := uu.jwtAuth.GenerateToken(int(user.ID))

	return token
}

func (uu *UserUsecase) GetAllusers() []Domain {
	return uu.userRepository.GetAllusers()
}
