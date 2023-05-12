package http_server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-shortener/internal/usecase"
	"net/http"
)

type Handler struct {
	host           string
	linkInteractor *usecase.LinkInteractor
}

func NewHandler(host string, linkInteractor *usecase.LinkInteractor) *Handler {
	return &Handler{
		host:           host,
		linkInteractor: linkInteractor,
	}
}

// POST /shorten

func (hd *Handler) AddLinkHandler(c *gin.Context) {
	type RequestLink struct {
		Source string `json:"source"`
	}

	var requestLink RequestLink
	if err := c.ShouldBindJSON(&requestLink); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	mapping, err := hd.linkInteractor.AddLink(requestLink.Source)
	if err != nil {
		c.String(http.StatusBadRequest, errors.Unwrap(err).Error())
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
	mapping := c.Params.ByName("mapping")
	source, err := hd.linkInteractor.GetLink(mapping)
	if err != nil {
		c.String(http.StatusBadRequest, errors.Unwrap(err).Error())
	}
	c.Redirect(http.StatusMovedPermanently, source)
}
