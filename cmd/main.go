package main

import (
	"product-api/controller"
	"product-api/db"
	"product-api/repository"
	"product-api/usecase"

	"github.com/gin-gonic/gin"
)

func main(){
	//definindo objeto gin
	server := gin.Default()

	dbConnection, err:= db.ConnectDB()
	if err != nil{
		panic(err)
	}

	//Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	UserRepository := repository.NewUserRepository(dbConnection)

	//Camada usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	UserUseCase := usecase.NewUserUseCase(UserRepository)
	//Camadas de controllers
	ProductController := controller.NewProductController(ProductUseCase)
	UserController := controller.NewUserController(UserUseCase)

	//primeiro parametro: rota
	//segundo parametro: funcao que sera executada quando a rota for acessada
	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.GET("/users", UserController.GetUsers)


	server.Run(":8000") // por padrao roda na porta 8080
}