package app

import (
	"github.com/arifin2018/facebook/helpers/handlers"
	requestvalidator "github.com/arifin2018/facebook/models/RequestValidator"
)

func InitProvider() {
	handlers.SetupDefaultConfigJwt()
}
func InitProviderRequestValidator() {
	requestvalidator.InitProviderValidatorRequestPost()
}
