package controller

import (
	"net/http"
	"product-api/model"
	"product-api/usecase"

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

func(p *produtcController) CreateProduct(ctx *gin.Context){
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}