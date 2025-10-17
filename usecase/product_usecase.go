package useCase

type ProductUseCase struct{
	//Repository
}

func NewProductUseCase() ProductUsecase{
	return ProductUsecase{}
}

// trata as regras de negócio da rota de get product
func (pu *ProductUsecase) GetProducts() ([]model.Product, error){
	return []model.Product{}, nil
}