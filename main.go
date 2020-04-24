package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/thehunter365/gorunner/runner"
)

func main() {
	code := []string{
		"package main",
		"import \"fmt\"",
		"func main() { ",
		"      fmt.Println(\"Salut c'est moi\")",
		"}",
	}
	jC, err := json.Marshal(runner.RawCode{CodeLines: code})
	if err != nil {
		log.Print(err)
	}
	r := runner.NewRunner(runner.GO, string(jC))
	out := r.StartRunner()
	fmt.Println(out)
}
