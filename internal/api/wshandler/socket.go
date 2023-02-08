package wshandler

import (
	"github.com/dwnGnL/validation/internal/application"
	"github.com/dwnGnL/validation/internal/repository"
	"github.com/dwnGnL/validation/lib/e"
	"github.com/dwnGnL/validation/lib/goerrors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

func (h handler) Registration(c *gin.Context) {

	app, err := application.GetRfromRequest(c)
	var info *repository.Users

	if err := c.ShouldBindJSON(&info); err != nil {
		log.Println(err, "test")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = app.Registration(info)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, err.Error())
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//c.JSON(http.StatusBadRequest, gin.H{"done": string(test)})
	c.JSON(http.StatusOK, "done")

}

func (h handler) Login(c *gin.Context) {

	app, err := application.GetLoginFromRequest(c)
	if err != nil {
		log.Println(err)
		return
	}

	var info *repository.Users

	if err := c.ShouldBindJSON(&info); err != nil {
		log.Println(err, "test")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	token, err := app.Login(info)
	if token == "" {
		c.JSON(http.StatusOK, "Login/password incorrect")
	} else {
		c.JSON(http.StatusOK, token)
	}
}

func GenRouting(r *gin.RouterGroup) {
	ws := newWsHandler()
	r.Any("/connect/:contestID", ws.wsContest)
	r.Any("/registrate", ws.Registration)
	r.Any("/login", ws.Login)
}
