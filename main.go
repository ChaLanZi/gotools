package main

import (
	"github.com/ChaLanZi/gotools/cmd/cli"
	"log"
)

const templateText = `
Output 0: {{title .Name1}}
Output 1: {{title .Name2}}
Output 2: {{.Name3 | title}}
`

func main() {
	err := cli.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute error: %v", err)
	}
}
