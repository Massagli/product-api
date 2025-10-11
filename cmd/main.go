package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	//primeiro parametro: rota
	//segundo parametro: funcao que sera executada quando a rota for acessada
	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8000") // por padrao roda na porta 8080


}