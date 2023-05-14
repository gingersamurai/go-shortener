package http_server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go-shortener/internal/config"
	"go-shortener/internal/usecase"
	"net/http"
	"time"
)

type Handler struct {
	engine         *gin.Engine
	hostAddr       string
	handleTimeout  time.Duration
	linkInteractor *usecase.LinkInteractor
}

func NewHandler(handlerConfig config.HandlerConfig, linkInteractor *usecase.LinkInteractor) *Handler {
	engine := gin.Default()

	handler := &Handler{
		engine:         engine,
		hostAddr:       handlerConfig.HostAddr,
		handleTimeout:  handlerConfig.HandleTimeout,
		linkInteractor: linkInteractor,
	}

	handler.engine.POST("/shorten", handler.AddLinkHandler)
	handler.engine.GET("/:mapping", handler.GetLinkHandler)

	return handler
}

// POST /shorten

func (hd *Handler) AddLinkHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), hd.handleTimeout)
	defer cancel()

	type RequestLink struct {
		Source string `json:"source"`
	}

	var requestLink RequestLink
	if err := c.ShouldBindJSON(&requestLink); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	mapping, err := hd.linkInteractor.AddLink(ctx, requestLink.Source)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	type ResponseLink struct {
		ShortenLink string `json:"shorten_link"`
	}

	responseLink := ResponseLink{
		ShortenLink: hd.hostAddr + "/" + mapping,
	}

	c.JSON(http.StatusOK, responseLink)

}

//get /:mapping

func (hd *Handler) GetLinkHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), hd.handleTimeout)
	defer cancel()

	mapping := c.Params.ByName("mapping")
	source, err := hd.linkInteractor.GetLink(ctx, mapping)
	if err != nil {
		c.String(http.StatusBadRequest, errors.Unwrap(err).Error())
		return
	}
	c.Redirect(http.StatusMovedPermanently, source)
}
