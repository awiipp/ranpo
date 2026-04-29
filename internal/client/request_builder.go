package client

import (
	"strings"

	"github.com/awiipp/ranpo/pkg/models"
	"github.com/go-resty/resty/v2"
)

func buildRequest(r *resty.Request, req *models.Request, vars map[string]string) {
	// Headers
	for k, v := range req.Headers {
		r.SetHeader(k, Resolve(v, vars))
	}

	// Auth
	switch strings.ToLower(req.Auth.Type) {
	case "bearer":
		token := Resolve(req.Auth.Token, vars)
		r.SetAuthToken(token)
	case "basic":
		r.SetBasicAuth(req.Auth.User, req.Auth.Pass)
	}

	// Body
	if req.Body != "" {
		if r.Header.Get("Content-Type") == "" {
			r.SetHeader("Content-Type", "application/json")
		}
		r.SetBody(Resolve(req.Body, vars))
	}
}
