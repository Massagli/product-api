package controller

import (
	"product-api/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

type produtcController struct{


	//useCase
	productUseCase useCase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) produtcController{
	return produtcController{
		productUseCase: usecase
	}
}

func(p *produtcController) GetProducts(ctx *gin.Context){
	products := []model.Product{
		{
			ID: 1,
			Name: "BATATA FRITA",
			Price: 20,
		},

		{
			ID: 2,
			Name: "FILÉ",
			Price: 40,
		},
	}

	ctx.JSON(http.StatusOK, products)
}