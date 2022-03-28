package auth

import _authRepository "latihan/coba-project/repo/auth"

type AuthUseCase struct {
	authRepository _authRepository.AuthRepositoryInterface
}

func NewAuthUseCase(authRepo _authRepository.AuthRepositoryInterface) AuthUseCaseInterface {
	return &AuthUseCase{
		authRepository: authRepo,
	}
}
func (auc *AuthUseCase) Login(email string, password string) (string, error) {
	token, err := auc.authRepository.Login(email, password)
	return token, err
}
