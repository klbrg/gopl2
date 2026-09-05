// server serves a page and its stylesheet from files embedded in the binary.
package main

import (
	"embed"
	"flag"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
)

var (
	addr = flag.String("addr", "localhost:8000", "HTTP service address")
	dev  = flag.Bool("dev", false, "serve static assets from disk, not from the binary")
)

//go:embed static
var staticFiles embed.FS

//go:embed templates/*.tmpl
var templateFiles embed.FS

var page = template.Must(template.ParseFS(templateFiles, "templates/*.tmpl"))

func main() {
	flag.Parse()

	static, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}
	if *dev {
		static = os.DirFS("static") // read from disk, no rebuild
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServerFS(static)))

	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		names, err := fs.Glob(static, "*")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := page.ExecuteTemplate(w, "index.tmpl", names); err != nil {
			log.Print(err)
		}
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
