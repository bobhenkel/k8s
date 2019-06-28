package main

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
// to run:
// go run 03-live-progress-and-capture-v3.go

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

func main() {
	cmd := exec.Command("kind", "create", "cluster", "--name", os.Getenv("PT_cluster_name"), "--image", os.Getenv("PT_image"))

	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/exec"
// )

// func main() {
// 	//https://stackoverflow.com/questions/1877045/how-do-you-get-the-output-of-a-system-command-in-go
// 	fmt.Println("Cluster Name:", os.Getenv("PT_cluster_name"))
// 	fmt.Println("Image:", os.Getenv("PT_image"))
// 	//out, err := exec.Command("kind", "delete", "cluster", "--name", os.Getenv("CLUSTER_NAME")).CombinedOutput()
// 	out, err := exec.Command("kind", "create", "cluster", "--name", os.Getenv("PT_cluster_name"), "--image", os.Getenv("PT_image")).CombinedOutput()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf(" %s\n", out)
// }

//cd into directory main.go is in
//This compile will place the binary in the tasks folder with the correct name.
//env GOOS=linux GOARCH=amd64 go build -o ../../create_kind_cluster-linux-amd64
//bolt task run k8s::create_kind_cluster_linux_amd64 cluster_name=kind-cluster2 image="kindest/node:v1.13.7" --nodes localhost --modulepath .
