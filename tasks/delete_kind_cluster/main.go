package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	//https://stackoverflow.com/questions/1877045/how-do-you-get-the-output-of-a-system-command-in-go
	fmt.Println("Cluster Name:", os.Getenv("PT_cluster_name"))
	out, err := exec.Command("kind", "delete", "cluster", "--name", os.Getenv("PT_cluster_name")).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" %s\n", out)
}

//env GOOS=linux GOARCH=amd64 go build -o ../../delete_kind_cluster_linux_amd64
//bolt task run k8s::delete_kind_cluster_linux_amd64 cluster_name=kind-cluster2 --nodes localhost --modulepath .
