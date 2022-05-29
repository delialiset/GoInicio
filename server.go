package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var prod []productos

type productos struct {
	Producto    string
	Descripcion string
}

func main() {

	http.HandleFunc("/", home)

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Host:", r.Host)
		fmt.Fprintln(w, "URI:", r.RequestURI)
		fmt.Fprintln(w, "Method:", r.Method)
		fmt.Fprintln(w, "RemoteAddr:", r.RemoteAddr)
	})

	tmpl := template.Must(template.ParseFiles("pro.html"))

	http.HandleFunc("/producto", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		produc := productos{
			Producto:    r.FormValue("producto"),
			Descripcion: r.FormValue("descrip"),
		}

		_ = produc

		tmpl.Execute(w, struct{ success bool }{true})

		prod = append(prod, produc)

		html := "<html>"
		html += "<body>"
		html += "<hr>"
		html += "<h3><center>Total Productos: " + strconv.Itoa(len(prod)) + "</h3>"

		if len(prod) == 0 {
			html += "No hay productos guardados"
		} else {
			html += "<h2>Listado de productos</h2>"

			for _, v := range prod {
				html += "<br>Producto: " + v.Producto
				html += "<br>Descripcion: " + v.Descripcion
				html += "<hr>"
			}
		}
		html += "<a href=\"/\">Home</a>"
		html += "</body>"
		html += "</html>"

		w.Write([]byte(html))
	})

	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {

	html := "<html>"
	html += "<body>"
	html += "<h1>Primera Pagina</h1>"
	html += "<br>"
	html += "<h3>Menu</h3>"
	html += "<br>"
	html += "<a href=\"/producto\">Productos</a>"
	html += "<br>"
	html += "<a href=\"/info\">Info</a>"
	html += "</body>"
	html += "</html>"

	w.Write([]byte(html))
}
