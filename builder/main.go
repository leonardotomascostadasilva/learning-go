package main

import "fmt"

// Struct que será construída
type Product struct {
	Name        string
	Description string
}

// Builder para construir o produto
type ProductBuilder struct {
	product Product
}

// Método estático para criar um novo builder
func CreateProductBuilder() *ProductBuilder {
	return &ProductBuilder{}
}

// Método para adicionar o nome ao produto
func (pb *ProductBuilder) AddName(name string) *ProductBuilder {
	pb.product.Name = name
	return pb
}

// Método para adicionar a descrição ao produto
func (pb *ProductBuilder) AddDescription(description string) *ProductBuilder {
	pb.product.Description = description
	return pb
}

// Método para construir o produto
func (pb *ProductBuilder) Build() Product {
	return pb.product
}

func main() {
	// Criando e construindo o produto
	product := CreateProductBuilder().
		AddName("Nome do Produto").
		AddDescription("Descrição do Produto").
		Build()

	// Exibindo o produto
	fmt.Printf("Name: %s, Description: %s\n", product.Name, product.Description)
}
