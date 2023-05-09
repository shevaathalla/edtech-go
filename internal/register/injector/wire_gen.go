// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package register

import (
	"edtech.id/internal/register/delivery/http"
	register2 "edtech.id/internal/register/usecase"
	"edtech.id/internal/user/repository"
	user2 "edtech.id/internal/user/usecase"
	"edtech.id/pkg/mail/gomail"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *register.RegisterHandler {
	userRepository := user.NewUserRepository(db)
	userUseCase := user2.NewUserUseCase(userRepository)
	smtpMail := mail.NewSmtpMail()
	registerUseCase := register2.NewRegisterUseCase(userUseCase, smtpMail)
	registerHandler := register.NewRegisterHandler(registerUseCase)
	return registerHandler
}