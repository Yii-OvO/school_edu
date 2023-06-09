package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/SupenBysz/gf-admin-community"
	_ "github.com/SupenBysz/gf-admin-company-modules"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	"github.com/kysion/edu-share/example/internal/boot"
)

func main() {
	boot.Main.Run(gctx.New())
}
