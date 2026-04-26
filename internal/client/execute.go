package client

import (
	"strings"
	"time"

	"github.com/awiipp/ranpo/pkg/models"
	"github.com/go-resty/resty/v2"
)

// Execute sends the request, applying environment variable resolution and auth.
func Execute(req *models.Request, env *models.Environment) (*models.Response, error) {
	vars := map[string]string{}
	if env != nil {
		vars = env.Variables
	}

	url := Resolve(req.URL, vars)

	c := resty.New().SetTimeout(30 * time.Second)
	r := c.R()

	buildRequest(r, req, vars)

	start := time.Now()

	var (
		resp *resty.Response
		err  error
	)

	switch strings.ToUpper(req.Method) {
	case "GET":
		resp, err = r.Get(url)
	case "POST":
		resp, err = r.Post(url)
	case "PUT":
		resp, err = r.Put(url)
	case "PATCH":
		resp, err = r.Patch(url)
	case "DELETE":
		resp, err = r.Delete(url)
	}

	if err != nil {
		return nil, err
	}

	return &models.Response{
		StatusCode: resp.StatusCode(),
		Status:     resp.Status(),
		Headers:    resp.Header(),
		Body:       resp.Body(),
		Duration:   time.Since(start),
	}, nil
}
