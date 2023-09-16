package handler

import (
	"net/http"

	svcErr "authentication-service/usecase/error"

	"github.com/gorilla/mux"
)

// show content type as json format
func configureDefaultHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
}

// allow cors for host
func configureCorsHeaders(w http.ResponseWriter, r *http.Request, origin string, allowedHeaders string) {
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
}

func RegisterNotFoundHandler(r *mux.Router) {
	r.NotFoundHandler = http.HandlerFunc(notFoundHandlerFunc)
}

func notFoundHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

// handle different errors
func processResponseErrorStatus(w http.ResponseWriter, err error, defaultErrorCode int) {
	if reqErr, ok := err.(*svcErr.ServiceError); ok {
		switch reqErr.Code {
		case svcErr.NO_DATA_FOUND:
			w.WriteHeader(http.StatusNotFound)
		case svcErr.PROCESSING_ERROR:
			w.WriteHeader(http.StatusInternalServerError)
		case svcErr.UNAUTHORIZED_REQUEST:
			w.WriteHeader(http.StatusForbidden)
		case svcErr.INVALID_REQUEST:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(defaultErrorCode)
		}
		return
	}
	w.WriteHeader(defaultErrorCode)
}
