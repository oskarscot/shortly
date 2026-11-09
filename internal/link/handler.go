package link

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpx "github.com/oskarscot/shortly/internal/platform"
)

type Handler struct {
	svc Service
}

func NewHandler(s Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) Setup(c *chi.Mux) {
	c.Route("/api/v1", func(r chi.Router) {
		r.Post("/links", h.createLink)
		r.Get("/links", h.getLinks)
	})
	c.Get("/s/{code}", h.redirectLink)
}

func (h *Handler) createLink(w http.ResponseWriter, r *http.Request) {
	var req LinkCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	ctx := r.Context()

	link, err := h.svc.CreateLink(ctx, req)
	if err != nil {
		if errors.Is(err, ErrInvalidUrl) {
			httpx.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.JSON(w, http.StatusCreated, link)
}

func (h *Handler) getLinks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	links, err := h.svc.ListLinks(ctx)

	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.JSON(w, http.StatusOK, links)
}

func (h *Handler) redirectLink(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if code == "" {
		httpx.Error(w, http.StatusBadRequest, "code is missing")
		return
	}

	ctx := r.Context()

	link, err := h.svc.FindByCode(ctx, code)
	if err != nil {
		if errors.Is(err, ErrLinkNotFound) {
			httpx.Error(w, http.StatusNoContent, err.Error())
			return
		}
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.Redirect(w, r, link.Url, http.StatusSeeOther)
}
