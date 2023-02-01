package wshandler

import (
	"github.com/dwnGnL/validation/internal/application"
	"github.com/dwnGnL/validation/lib/e"
	"github.com/dwnGnL/validation/lib/goerrors"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func newWsHandler() *handler {
	return &handler{}
}

func (h handler) wsContest(c *gin.Context) {
	app, err := application.GetAppFromRequest(c)
	if err != nil {
		goerrors.Log().Warn("fatal err: %w", err)
		e.With(err).Code(e.InternalErrorCode).Detail("GetAppFromRequest err").Write(c)
		return
	}
	app.TestService()
}

func GenRouting(r *gin.RouterGroup) {
	ws := newWsHandler()
	r.Any("/connect/:contestID", ws.wsContest)

}
