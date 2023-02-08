package application

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

type ContextValue int

const (
	ContextApp ContextValue = iota
)

func GetAppFromContext(ctx context.Context) (Core, error) {
	if app, ok := ctx.Value(ContextApp).(Core); ok {
		return app, nil
	}

	return nil, errors.New("cannot get app container from request")
}

func GetAppFromRequest(r *gin.Context) (Core, error) {
	return GetAppFromContext(r.Request.Context())
}

func GetRfromContext(ctx context.Context) (RegistrationCore, error) {
	if app, ok := ctx.Value(ContextApp).(RegistrationCore); ok {
		return app, nil
	}

	return nil, errors.New("cannot get app container from request")
}

func GetRfromRequest(r *gin.Context) (RegistrationCore, error) {
	return GetRfromContext(r.Request.Context())
}

func GetLoginFromContext(ctx context.Context) (LoginCore, error) {
	if app, ok := ctx.Value(ContextApp).(LoginCore); ok {
		return app, nil
	}
	return nil, errors.New("cannot get app container from request")
}

func GetLoginFromRequest(r *gin.Context) (LoginCore, error) {
	return GetLoginFromContext(r.Request.Context())
}
