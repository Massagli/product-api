package controller

import (
	"product-api/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
)

type produtcController struct{
	//useCase
	productUseCase usecase.ProductUsecase
}
 
//Inicializa a estrutura
func NewProductController(usecase usecase.ProductUsecase) produtcController{
	return produtcController{
		productUseCase: usecase,
	}
} 

func(p *produtcController) GetProducts(ctx *gin.Context){
	products, err := p.productUseCase.GetProducts()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}