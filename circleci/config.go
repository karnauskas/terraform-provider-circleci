package circleci

import (
	"fmt"

	circleci "github.com/ryanlower/go-circleci"
)

type Config struct {
	ApiToken string
}

type CircleCIClient struct {
	conn *circleci.Client
}

func (c *Config) Client() (interface{}, error) {
	var client CircleCIClient

	if c.ApiToken == "" {
		return nil, fmt.Errorf("[Err] No API Token for CircleCI")
	}

	fconn := circleci.NewClient(c.ApiToken)

	client.conn = fconn

	return &client, nil
}
