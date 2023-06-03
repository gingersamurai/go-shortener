package grpc_server

import (
	"context"
	"errors"
	"go-shortener/api/link"
	"go-shortener/internal/config"
	"go-shortener/internal/usecase"
	"strings"
	"time"
)

type Handler struct {
	link.UnimplementedLinkShortenerServiceServer

	hostAddr       string
	handleTimeout  time.Duration
	linkInteractor *usecase.LinkInteractor
}

func NewHandler(handlerConfig config.HandlerConfig, linkInteractor *usecase.LinkInteractor) *Handler {
	return &Handler{
		hostAddr:       handlerConfig.HostAddr,
		handleTimeout:  handlerConfig.HandleTimeout,
		linkInteractor: linkInteractor,
	}
}

func (h *Handler) AddLink(ctx context.Context, in *link.Link) (*link.Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()

	mapping, err := h.linkInteractor.AddLink(ctx, in.Value)
	if err != nil {
		return nil, err
	}
	result := &link.Link{Value: h.hostAddr + "/" + mapping}

	return result, nil
}

func (h *Handler) GetLink(ctx context.Context, in *link.Link) (*link.Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), h.handleTimeout)
	defer cancel()
	id := strings.LastIndex(in.Value, "/")
	if id == -1 {
		return nil, errors.New("mapping not found")
	}
	mapping := in.Value[id+1:]
	fullLink, err := h.linkInteractor.GetLink(ctx, mapping)
	if err != nil {
		return nil, err
	}

	result := &link.Link{Value: fullLink}
	return result, nil
}
