package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("./run_test.sh", "yo", "how", "r", "u", "?")
	var outputBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(outputBuffer.String())
}
