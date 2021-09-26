package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

func helloTask(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello http\n")
	w.Write(([]byte("hello")))
}

func listTask(w http.ResponseWriter, req *http.Request) {
	dir, _ := GetCurrentPath()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("ioutil.ReadDir(%s), err: %s", dir, err)
	}
	for _, f := range files {
		fmt.Fprintf(w, f.Name()+"\n")
	}
}

func startHttpServer() {
	router := mux.NewRouter()

	router.HandleFunc("/api/hello", helloTask)
	router.HandleFunc("/api/list", listTask)
	curdir, _ := GetCurrentPath()
	pathSep := string(os.PathSeparator)
	filePath := curdir + pathSep + "file"
	// 好像不行,垃圾
	router.PathPrefix("/file/").Handler(http.StripPrefix("/file/", http.FileServer(http.Dir(filePath))))

	srv := &http.Server{
		Handler: router,
		Addr:    ":10086",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func GetCurrentPath() (dir string, err error) {
	log.Println(os.Args)
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Println("Exec.LookPath (%s), err: %s", os.Args[0], err)
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Println("filepath.Abs(%s), err: %s", path, err)
	}
	dir = filepath.Dir(absPath)
	return dir, nil
}

func main() {
	startHttpServer()
}
