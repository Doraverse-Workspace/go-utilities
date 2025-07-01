package cf

import (
	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func NewClient(apiToken string) *cloudflare.Client {
	c := cloudflare.NewClient(option.WithAPIToken(apiToken))
	return c
}
