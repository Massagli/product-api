package repository

import (
	"database/sql"
	"fmt"
	"product-api/model"
)

//Estrutura
type UserRepository struct{
	//iniciando banco
	connection *sql.DB
}

//Inicializa a estrutura
func NewUserRepository(connection *sql.DB) UserRepository{
	return UserRepository{
		connection: connection,
	}
}

func(user *UserRepository) GetUser() ([]model.User, error){
	//query de busca no banco
	query := "SELECT id, user_name, user_age FROM user"
	//trazendo resultado da query na variável "rows"
	rows, err := user.connection.Query(query)

	//verificando erro na query
	if err != nil{
		fmt.Print(err)
		return []model.User{}, err
	}

	//definindo objetos que vão carregar os dados
	var userList []model.User
	var userObj model.User

	//passando pela as linhas do resultados da query
	//rows.Next(): avança para a próxima linha do resultado
	//Retorna true se houver mais linhas, false quando acabar
	//É um loop que percorre todas as linhas retornadas pela query
	for rows.Next(){
		// lê os valores da linha atual e atribui às variáveis
		// atribuindo os resultado da query ao modelo criado
		// A ordem dos campos deve ser a mesma ordem retornada pelo SELECT *
		err = rows.Scan(
			&userObj.ID,
			&userObj.User_name,
			&userObj.User_age)

		if err != nil{
			fmt.Println(err)
			return []model.User{}, err
		}

		//jogando o resultado numa lista
		userList = append(userList, userObj)
	}
	rows.Close()

	return userList, nil
}