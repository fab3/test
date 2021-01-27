// package main

// func serveForm(w http.ResponseWriter, r *http.Request) {
// 	// we can use a walruss as a string!
// 	// r.ParseForm()
// 	uri := fmt.Sprintf(" RequestURI: %v\n  Hosts: %v\n Form:%v\n ", r.RequestURI, r.Host, r.Form)
// 	// log.Printf("I See u using Method %v\n", r.Method)
// 	log.Printf("RequestURI %v\n", r.RequestURI)
// 	// log.Printf(uri)
// 	fmt.Fprintf(w, uri)
// 	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

// }

// func serveFile(w http.ResponseWriter, r *http.Request) {
// 	var err error
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	http.ServeFile(w, r, filepath.Join(wd, r.URL.Path))
// 	// log.Println("WD:", wd)
// }

// func main() {
// 	http.HandleFunc("/form", serveForm)
// 	http.HandleFunc("/", serveFile)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
// starts here!!!!
// func serve(w http.ResponseWriter, r *http.Request) {
// 	// log.Print(w, r)
// 	fmt.Fprint(w, "<h1>Hello!</h1>") //server write sth
// 	uri := fmt.Sprintf(" RequestURI: %v\n  Hosts: %v\n Form:%v\n ", r.RequestURI, r.Host, r.Form)
// 	log.Printf("I See u using Method %v\n", r.Method)
// 	log.Printf("RequestURI %v\n", r.RequestURI)
// 	log.Printf(uri)
// }
// func main() {
// 	// http.HandleFunc("/", serve)
// 	//handle using a static file
// 	http.Handle("/form", http.FileServer(http.Dir("./GO/repos/testdata")))
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
// next
package main

import (
	"log"
	"net/http"
	"os"
)

// func serveForm(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	jsn, _ := json.MarshalIndent(r.Form, "", "  ")
// 	fmt.Fprintf(w, string(jsn))
// }

func serve(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "POST" {
	// 	serveForm(w, r)
	// 	return
	// }
	var err error
	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		return
	}
	http.ServeFile(w, r, wd+r.URL.Path)
}

func main() {
	http.HandleFunc("/", serve)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
