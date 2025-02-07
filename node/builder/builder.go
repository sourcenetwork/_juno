package builder

import (
	"fmt"

	"github.com/forbole/juno/v5/node"
	nodeconfig "github.com/forbole/juno/v5/node/config"
	"github.com/forbole/juno/v5/node/local"
	"github.com/forbole/juno/v5/node/remote"
	"github.com/forbole/juno/v5/types/params"
)

func BuildNode(cfg nodeconfig.Config, encodingConfig params.EncodingConfig) (node.Node, error) {
	switch cfg.Type {
	case nodeconfig.TypeRemote:
		return remote.NewNode(cfg.Details.(*remote.Details), encodingConfig.Codec)
	case nodeconfig.TypeLocal:
		return local.NewNode(cfg.Details.(*local.Details), encodingConfig.TxConfig, encodingConfig.Codec)
	case nodeconfig.TypeNone:
		return nil, nil

	default:
		return nil, fmt.Errorf("invalid node type: %s", cfg.Type)
	}
}
