package app

import (
	"github.com/arifin2018/facebook/helpers/handlers"
	"github.com/arifin2018/facebook/models/requestValidator"
)

func InitProvider() {
	handlers.SetupDefaultConfigJwt()
}
func InitProviderRequestValidator() {
	requestValidator.InitProviderValidatorRequestPost()
	requestValidator.InitProviderValidatorRequestPostImage()
	requestValidator.InitProviderValidatorRequestPermission()
	requestValidator.InitProviderValidatorRequestRole()
}
