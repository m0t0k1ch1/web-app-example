package appv1

import (
	"github.com/m0t0k1ch1/web-app-sample/backend/handler"
)

type AppServiceHandler struct {
	env *handler.Env
}

func NewAppServiceHandler(env *handler.Env) *AppServiceHandler {
	return &AppServiceHandler{
		env: env,
	}
}
