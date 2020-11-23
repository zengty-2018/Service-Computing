package service

import (
	"html/template"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}

	mx.HandleFunc("/test", testHandler(formatter)).Methods("GET")
	mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")
	mx.HandleFunc("/user", userHandler).Methods("POST")
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
}

func testHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			S_ID string `json:"s_id"`
			Name string `json:"name"`
		}{S_ID: "18342007", Name: "zengty"})
	}
}

func homeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "page", struct {
			Content string `json:"content"`
		}{Content: "Input your name and S_ID."})
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := template.HTMLEscapeString(r.Form.Get("name"))
	s_id := template.HTMLEscapeString(r.Form.Get("s_id"))
	t := template.Must(template.New("user.html").ParseFiles("./templates/user.html"))
	err := t.Execute(w, struct {
		Name string
		S_ID string
	}{Name: name, S_ID: s_id})
	if err != nil {
		panic(err)
	}
}
