package register

import (
	registerDto "edtech.id/internal/register/dto"
	userDto "edtech.id/internal/user/dto"
	userUseCase "edtech.id/internal/user/usecase"

	mail "edtech.id/pkg/mail/gomail"
)

type RegisterUseCase interface {
	Register(registerDto userDto.UserRequestBody) error
}

type RegisterUseCaseImpl struct {
	userUseCase userUseCase.UserUseCase
	mail        mail.SmtpMail
}

func NewRegisterUseCase(
	userUseCase userUseCase.UserUseCase,
	mail mail.SmtpMail) RegisterUseCase {
	return &RegisterUseCaseImpl{userUseCase, mail}
}

func (ru *RegisterUseCaseImpl) Register(userDto userDto.UserRequestBody) error {
	// TODO: Check module user
	user, err := ru.userUseCase.Create(userDto)

	if err != nil {
		return err
	}

	// TODO: Send email
	email := registerDto.CreateEmailVerification{
		SUBJECT:           "Kode Verifikasi Email",
		EMAIL:             user.Email,
		VERIFICATION_CODE: user.CodeVerified,
	}

	go ru.mail.SmtpSendVerificationEmail(user.Email, email)

	return nil
}
