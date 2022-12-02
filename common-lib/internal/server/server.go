package server

import (
	"net/http"
)

type HttpHandlerFn func(h http.Handler)
