package main

import (
	_ "github.com/ledboot/OwlHook/internal/boot"
	_ "github.com/ledboot/OwlHook/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/ledboot/OwlHook/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
