package circleci

import (
	"fmt"

	"github.com/jszwedko/go-circleci"
)

type Config struct {
	Token string
}

const noToken = `
[Err] No API Token for CircleCI

`

func (c *Config) Client() (*circleci.Client, error) {
	if c.Token == "" {
		return nil, fmt.Errorf(noToken)
	}

	client := circleci.Client{
		Token: c.Token,
	}

	return &client, nil
}
