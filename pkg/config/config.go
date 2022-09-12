package config

import (
	"net/url"
)

type Config struct {
	Parallel int
	URLs     []url.URL
}
