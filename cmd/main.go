package main

import (
	"product-api/controller"
	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()


	ProductUseCase := usecase.NewProductUseCase()
	//Camadas de controllers
	ProductController := controller.NewProductController(productUseCase)

	//primeiro parametro: rota
	//segundo parametro: funcao que sera executada quando a rota for acessada
	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000") // por padrao roda na porta 8080
}