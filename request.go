package ebay

import (
	"github.com/cubixle2/go-ebay/commands"
	"github.com/cubixle2/go-ebay/config"
)

const xmlns = "urn:ebay:apis:eBLBaseComponents"

type Request struct {
	Config  *config.Config
	Command commands.Command
}
