package main

//goland:noinspection GoUnsortedImport
import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/Wyatex/SakuraAdminGF/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
