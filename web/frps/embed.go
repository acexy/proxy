package frps

import (
	"embed"

	"github.com/acexy/proxy/assets"
)

//go:embed dist
var EmbedFS embed.FS

func init() {
	assets.Register(EmbedFS)
}
