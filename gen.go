// +build ignore

// This program generates methods.go, request.go and response.go. It can be invoked by running
// go generate
package main

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "-c", "if [ -f 'venv/bin/activate' ]; then source venv/bin/activate; fi ; python generator.py 1>&2 stderr")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stderr.Close()

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Waiting for generate telegram API...")

	slurp, _ := ioutil.ReadAll(stderr)
	log.Printf("%s", slurp)
	if err = cmd.Wait(); err != nil {
		log.Fatalf("Command finished with error: %v", err)
	}
	log.Printf("Command finished without error")
}
