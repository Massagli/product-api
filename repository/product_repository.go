package repository

import (
	"database/sql"
	"fmt"
	"product-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

// Inicializa a estrutura
func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

// função que busca os dados no banco
func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
	}
	rows.Close()
	productList = append(productList, productObj)

	return productList, nil

}

func(pr *ProductRepository) CreateProduct(product model.Product) (int, error){

	var id int
	query := "INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id"
	row, err := pr.connection.Prepare(query)
	if err != nil{
		fmt.Println(err)
		return 0, err
	}

	err = row.QueryRow(product.Name, product.Price).Scan(&id)
		if err != nil{
		fmt.Println(err)
		return 0, err
	}
	row.Close()
	return id, nil
}