package runner

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/thehunter365/gorunner/utils"
)

//Lang type
type Lang int

const (
	//GO golang code type
	GO Lang = iota
	//JAVA code
	JAVA
	//PYTHON code
	PYTHON
)

//RawCode type
type RawCode struct {
	CodeLines []string `json:"codeLines"`
}

//Runner type
type Runner struct {
	RawCode string

	Lang      Lang
	CodeLines []string
	Return    Response

	TimeOut time.Duration
}

//Response struct
type Response struct {
	Messages []string
}

//NewRunner func
func NewRunner(lang Lang, jsonCode RawCode) *Runner {
	return &Runner{
		CodeLines: jsonCode.CodeLines,
		Lang:    lang,
		TimeOut: 5,
	}
}

//ParseCode from json to plain old text
func (r *Runner) ParseCode() (code []string) {
	var rc RawCode
	if r.RawCode == "" {
		return
	}
	err := json.Unmarshal([]byte(r.RawCode), &rc)
	handleErr(err)
	code = rc.CodeLines
	r.CodeLines = code
	return
}

//StartRunner func
func (r *Runner) StartRunner() (out []string) {
	out = r.execCode()
	r.Return = Response{out}
	return
}

//ExecCode func
func (r *Runner) execCode() (stdout []string) {
	c1 := make(chan []byte, 1)

	if len(r.CodeLines) == 0 {
		r.ParseCode()
	}

	utils.FileWrite("tmp.go", r.CodeLines)
	cmd := exec.Command("go", "run", "../tmp.go")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	go func() {
		out, err := cmd.CombinedOutput()
		c1 <- out
		if err != nil {
			log.Fatalln(err)
		}
	}()
	select {
	case <-time.After(r.TimeOut * time.Second):
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		log.Println("process timed out")
	case o := <-c1:
		stdout = strings.Split(string(o), "\n")
		log.Print("process finished successfully")
	}
	utils.FileDelete("tmp.go")
	return
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
