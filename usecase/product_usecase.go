package usecase

import (
	"product-api/model"
	"product-api/repository"
)

type ProductUsecase struct {
	//Repository
	repository repository.ProductRepository 
}

//Inicializa a estrutura
func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

// trata as regras de negócio da rota de get product
func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error){
	productId, err := pu.repository.CreateProduct(product)
	if err != nil{
		return model.Product{}, err
	}

	product.ID = productId
	
	return product, nil
}

 