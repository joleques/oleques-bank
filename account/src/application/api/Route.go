package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	renderChi "github.com/go-chi/render"
	"github.com/joleques/oleques-bank/account/src/application/dto"
	renderPkg "github.com/unrolled/render"
	"net/http"
	"time"
)

var render *renderPkg.Render
var logger http.Handler

func Start(status string) {
	contentType := middleware.AllowContentType("application/json")
	render = renderPkg.New()
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(contentType)
	r.Use(renderChi.SetContentType(renderChi.ContentTypeJSON))
	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := dto.ApiResponse{StatusCode: 200, Message: status}
		render.JSON(w, 200, response)
	})

	r.Route("/accounts", func(r chi.Router) {
		r.Post("/", CreateAccount)
		r.Get("/", List)
		r.Get("/{account_id}/balance", Get)
		r.Post("/{account_id}/balance", Update)
	})

	panic(http.ListenAndServe(":3000", r))
}
