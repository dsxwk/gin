package main

import (
	_ "gin/app/command"
	_ "gin/app/command/cronjob"
	"gin/utils/cli"
	_ "gin/utils/cli/db"
	_ "gin/utils/cli/make"
	_ "gin/utils/cli/route"
)

func main() {
	cli.Execute()
}
