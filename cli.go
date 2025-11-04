package main

import (
	_ "gin/app/command"
	_ "gin/app/listener"
	"gin/utils/cli"
	_ "gin/utils/cli/db"
	_ "gin/utils/cli/event"
	_ "gin/utils/cli/make"
	_ "gin/utils/cli/route"
)

func main() {
	cli.Execute()
}
