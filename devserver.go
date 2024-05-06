//go:build ignore
// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/vugu/vugu/devutil"
)

func main() {
	l := "127.0.0.1:8844"
	log.Printf("Starting HTTP Server at %q", l)

	wc := devutil.NewWasmCompiler().SetDir(".")
	mux := devutil.NewMux()
	mux.Match(devutil.NoFileExt, devutil.DefaultAutoReloadIndex.Replace(
		`<!-- styles -->`,
		`
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-material-design/4.0.2/bootstrap-material-design.css">
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-material-design/4.0.2/bootstrap-material-design.min.css">
		<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@24,400,0,0" />
		`).Replace(
		`<title>Vugu App</title>`,
		`<title>Vugu Sample Project</title>`))
	mux.Match(devutil.NoFileExt, devutil.DefaultAutoReloadIndex.Replace(
		`<!-- scripts -->`,
		`
		<script defer src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-material-design/4.0.2/bootstrap-material-design.umd.min.js"></script>
		<script defer src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-material-design/4.0.2/bootstrap-material-design.iife.js"></script>
		<script defer src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-material-design/4.0.2/bootstrap-material-design.iife.js.map"></script>
		<script defer src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-material-design/4.0.2/bootstrap-material-design.iife.min.js"></script>
		<script defer src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-material-design/4.0.2/bootstrap-material-design.umd.js"></script>
		<script defer src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-material-design/4.0.2/bootstrap-material-design.umd.js.map"></script>
		<script defer src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-material-design/4.0.2/bootstrap-material-design.umd.min.js.map"></script>
		`))
	mux.Exact("/main.wasm", devutil.NewMainWasmHandler(wc))
	mux.Exact("/wasm_exec.js", devutil.NewWasmExecJSHandler(wc))
	mux.Default(devutil.NewFileServer().SetDir("."))

	log.Fatal(http.ListenAndServe(l, mux))
}
