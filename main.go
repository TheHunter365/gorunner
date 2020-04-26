package main

import "github.com/thehunter365/gorunner/runner"

func main() {

	server := runner.NewServer(":8080")
	server.Start()
}

/*code := []string{
	"package main",
	"import \"fmt\"",
	"func main() { ",
	"      fmt.Println(\"ce programme est inutile\")",
	"}",
}

sCode := []string{
	"package main",
	"func main() {",
	"	var i int",
	"	for {i++}",
	"}",
}
jC, err := json.Marshal(runner.RawCode{CodeLines: code})
if err != nil {
	log.Print(err)
}
r := runner.NewRunner(runner.GO, string(jC))
out := r.StartRunner()
fmt.Println(out)

jC, err = json.Marshal(runner.RawCode{CodeLines: sCode})
if err != nil {
	log.Println(err)
}
r = runner.NewRunner(runner.GO, string(jC))
out = r.StartRunner()
fmt.Println(out)*/
