package main

import (
	"net/http"

	"github.com/dominicgerman/dominicgerman.com/ui"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /about", dynamic.ThenFunc(app.about))
	mux.Handle("GET /blog", dynamic.ThenFunc(app.blog))
	mux.Handle("GET /posts/{id}", dynamic.ThenFunc(app.handlerPostView))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.handlerUserLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.handlerUserLoginPost))

	protected := dynamic.Append(app.requireAuthentication)

	mux.Handle("GET /post/create", protected.ThenFunc(app.handlerPostCreate))
	mux.Handle("GET /post/update/{id}", protected.ThenFunc(app.handlerPostUpdate))
	mux.Handle("POST /post/create", protected.ThenFunc(app.handlerPostCreatePost))
	mux.Handle("/post/update/{id}", protected.ThenFunc(app.handlerPostUpdatePut))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.handlerUserLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
