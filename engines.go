package gogpt

import (
	"fmt"
	"net/http"
)

// Engine struct represents engine from OpenAPI API.
type Engine struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Owner  string `json:"owner"`
	Ready  bool   `json:"ready"`
}

// EnginesList is a list of engines.
type EnginesList struct {
	Engines []Engine `json:"data"`
}

// ListEngines Lists the currently available engines, and provides basic
// information about each option such as the owner and availability.
func (c *Client) ListEngines() (engines EnginesList, err error) {
	req, err := http.NewRequest(http.MethodGet, c.fullURL("/engines"), nil)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &engines)
	return
}

// GetEngine Retrieves an engine instance, providing basic information about
// the engine such as the owner and availability.
func (c *Client) GetEngine(
	engineID string,
) (engine Engine, err error) {
	urlSuffix := fmt.Sprintf("/engines/%s", engineID)
	req, err := http.NewRequest(http.MethodGet, c.fullURL(urlSuffix), nil)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &engine)
	return
}
