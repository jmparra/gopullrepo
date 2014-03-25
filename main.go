package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	repopath := flag.String("repo", "~/repo", "Path to git repo to pull")
	address := flag.String("address", "127.0.0.1:8080", "Listen address or just port like \":8080\"")
	httppath := flag.String("path", "/", "Path to listen at, other paths will be 404")
	flag.Parse()

	http.HandleFunc(*httppath, func(w http.ResponseWriter, r *http.Request) {
		os.Chdir(*repopath)
		cmd := exec.Command("git", "pull")
		out, err := cmd.Output()
		fmt.Fprintf(w, "Ok: %s", string(out[:]))
                if err != nil {
                        fmt.Fprintf(w, "Error: %s", err)
                }
	})

	log.Fatal(http.ListenAndServe(*address, nil))
}
