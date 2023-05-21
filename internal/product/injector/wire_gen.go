// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package product

import (
	"edtech.id/internal/product/delivery/http"
	product2 "edtech.id/internal/product/repository"
	product3 "edtech.id/internal/product/usecase"
	"edtech.id/pkg/fileupload/cloudinary"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *product.ProductHandler {
	productRepository := product2.NewProductRepository(db)
	fileUpload := fileupload.NewFileUpload()
	productUseCase := product3.NewProductUsecase(productRepository, fileUpload)
	productHandler := product.NewProductHandler(productUseCase)
	return productHandler
}
