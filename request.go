package ebay

import (
	"github.com/cubixle/go-ebay/commands"
	"github.com/cubixle/go-ebay/config"
)

const xmlns = "urn:ebay:apis:eBLBaseComponents"

type Request struct {
	Config  *config.Config
	Command commands.Command
}
