package cmd

import (
	"fmt"

	"github.com/awiipp/ranpo/internal/client"
	"github.com/awiipp/ranpo/internal/config"
	"github.com/awiipp/ranpo/internal/renderer"
	"github.com/awiipp/ranpo/internal/store"
	"github.com/awiipp/ranpo/pkg/models"
)

// Shared HTTP execution path for all non-interactive commands.
func executeRequest(method, url, body string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}

	env, _ := store.LoadEnv(cfg.ActiveEnv)

	// Resolve auth
	auth := models.AuthConfig{Type: "none"}

	if flagToken != "" {
		auth = models.AuthConfig{Type: "bearer", Token: flagToken}
	} else if env != nil {
		if token, ok := env.Variables["TOKEN"]; ok && token != "" {
			auth = models.AuthConfig{Type: "bearer", Token: token}
		}
	} else if cfg.DefaultAuth.Token != "" {
		auth = models.AuthConfig{Type: "bearer", Token: cfg.DefaultAuth.Token}
	}

	req := &models.Request{
		Method:  method,
		URL:     url,
		Headers: parseHeaders(flagHeaders),
		Body:    body,
		Auth:    auth,
	}

	// Save if requested
	if flagSave != "" {
		req.Name = flagSave

		col, _ := store.LoadCollection(flagCollection)
		if col == nil {
			col = &models.Collection{Name: flagCollection}
		}

		replaced := false
		for i, r := range col.Requests {
			if r.Name == flagSave {
				col.Requests[i] = *req
				replaced = true
				break
			}
		}

		if !replaced {
			col.Requests = append(col.Requests, *req)
		}

		if err := store.SaveCollection(col); err != nil {
			fmt.Printf("warning: could not save request: %v\n", err)
		} else {
			fmt.Printf("saved %q to collection %q\n", flagSave, flagCollection)
		}
	}

	resp, err := client.Execute(req, env)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	fmt.Print(renderer.RenderResponse(resp.StatusCode, resp.Status, resp.Body, resp.Duration))
	return nil
}
