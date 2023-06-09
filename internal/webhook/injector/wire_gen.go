// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package webhook

import (
	"edtech.id/internal/cart/repository"
	cart2 "edtech.id/internal/cart/usecase"
	"edtech.id/internal/class_room/repository"
	class_room2 "edtech.id/internal/class_room/usecase"
	"edtech.id/internal/discount/repository"
	discount2 "edtech.id/internal/discount/usecase"
	"edtech.id/internal/order/repository"
	order2 "edtech.id/internal/order/usecase"
	"edtech.id/internal/order_detail/repository"
	order_detail2 "edtech.id/internal/order_detail/usecase"
	"edtech.id/internal/payment/usecase"
	"edtech.id/internal/product/repository"
	product2 "edtech.id/internal/product/usecase"
	"edtech.id/internal/webhook/delivery/http"
	webhook2 "edtech.id/internal/webhook/usecase"
	"edtech.id/pkg/fileupload/cloudinary"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *webhook.WebhookHandler {
	orderRepository := order.NewOrderRepository(db)
	cartRepository := cart.NewCartRepository(db)
	cartUsecase := cart2.NewCartUsecase(cartRepository)
	discountRepository := discount.NewDiscountRepository(db)
	discountUsecase := discount2.NewDiscountUsecase(discountRepository)
	productRepository := product.NewProductRepository(db)
	fileUpload := fileupload.NewFileUpload()
	productUseCase := product2.NewProductUsecase(productRepository, fileUpload)
	orderDetailRepository := order_detail.NewOrderDetailRepository(db)
	orderDetailUsecase := order_detail2.NewOrderDetailUsecase(orderDetailRepository)
	paymentUsecase := payment.NewPaymentUsecase()
	orderUsecase := order2.NewOrderUsecase(orderRepository, cartUsecase, discountUsecase, productUseCase, orderDetailUsecase, paymentUsecase)
	classRoomRepository := class_room.NewClassRoomRepository(db)
	classRoomUsecase := class_room2.NewClassRoomUsecase(classRoomRepository)
	webhookUsecase := webhook2.NewWebhookUsecase(orderUsecase, classRoomUsecase)
	webhookHandler := webhook.NewWebhookHandler(webhookUsecase)
	return webhookHandler
}
