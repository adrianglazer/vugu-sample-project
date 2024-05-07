//go:build ignore
// +build ignore

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vugu/vugu/devutil"
)

func main() {
	indexBuf, err := os.ReadFile("public/index.html")

	if err != nil {
		panic(err)
	}
	index := string(indexBuf)

	l := "127.0.0.1:8844"
	log.Printf("Starting HTTP Server at %q", l)

	wc := devutil.NewWasmCompiler().SetDir(".")
	mux := devutil.NewMux()
	mux.Match(devutil.NoFileExt, devutil.StaticContent(index))
	mux.Exact("/main.wasm", devutil.NewMainWasmHandler(wc))
	mux.Exact("/wasm_exec.js", devutil.NewWasmExecJSHandler(wc))
	mux.Default(devutil.NewFileServer().SetDir("."))

	log.Fatal(http.ListenAndServe(l, mux))
}
