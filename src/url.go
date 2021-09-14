package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"regexp"
)

var (
	Url      string
	parseUrl = regexp.MustCompile(`(http(s?):)([/|.|\w|\s|-])*\.(?:.* )`)
)

func init() {

	flag.StringVar(&Url, "t", "", "bitly url")
	flag.Parse()
}

func main() {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("sh", "-c", `curl `+Url)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	peo := cmd.Run()
	if peo != nil {
		fmt.Println(peo)
	}
	// capture the stderr and stdout
	out := stdout.String() + stderr.String()
	coolOut := string(parseUrl.Find([]byte(out)))
	fmt.Println(coolOut)
}
