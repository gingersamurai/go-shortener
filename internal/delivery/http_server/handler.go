package http_server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go-shortener/internal/usecase"
	"net/http"
	"time"
)

type Handler struct {
	host           string
	handleTimeout  time.Duration
	linkInteractor *usecase.LinkInteractor
}

func NewHandler(host string, handleTimeout time.Duration, linkInteractor *usecase.LinkInteractor) *Handler {
	return &Handler{
		host:           host,
		handleTimeout:  handleTimeout,
		linkInteractor: linkInteractor,
	}
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
		ShortenLink: hd.host + "/" + mapping,
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
