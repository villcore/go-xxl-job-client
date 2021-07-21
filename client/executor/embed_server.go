package executor

import (
	"net/http"
	"strconv"
)

type EmbedServer interface {
	Start() error
	Stop() error
}

type EmbedHttpServer struct {
	config            JobConfig
	clientExecutorBiz ClientExecutorBiz
	httpServer        *http.Server
}

func (receiver *EmbedHttpServer) Start() error {

	httpServerMux := http.NewServeMux()
	for pattern, handler := range receiver.httpPatterns() {
		httpServerMux.Handle(pattern, handler)
	}

	receiver.httpServer = &http.Server{
		Addr:    ":" + strconv.Itoa(int(receiver.config.Port)),
		Handler: httpServerMux,
	}
	return receiver.httpServer.ListenAndServe()
}

func (receiver *EmbedHttpServer) Stop() error {
	if receiver.httpServer != nil {
		return receiver.httpServer.Close()
	}
	return nil
}

func (receiver *EmbedHttpServer) httpPatterns() map[string]http.Handler {
	return map[string]http.Handler{
		"/beat":     receiver.beatHttpHandler(),
		"/idleBeat": receiver.beatHttpHandler(),
		"/run":      receiver.beatHttpHandler(),
		"/kill":     receiver.beatHttpHandler(),
		"/log":      receiver.beatHttpHandler(),
	}
}

func (receiver *EmbedHttpServer) beatHttpHandler() http.Handler {
	receiver.clientExecutorBiz
}
