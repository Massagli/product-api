package usecase

import (
	"product-api/model"
	"product-api/repository"
)

//Recebe a estrutura do repository
type UserUsecase struct {
	repository repository.UserRepository
}

//Essa função cria uma nova instância de UserUsecase.
//Ela recebe um repositório (repo) e o injeta na struct.
func NewUserUseCase(repo repository.UserRepository) UserUsecase{
	return UserUsecase{
		repository: repo,
	}
}

//Cuida de toda a lógica e funcionalidade do produto
func(us *UserUsecase) GetUsers() ([]model.User, error){
	return us.repository.GetUser()
}