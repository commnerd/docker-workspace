package main

import (
    "github.com/gorilla/mux"
    "encoding/json"
    "text/template"
    "io/ioutil"
    "net/http"
    // "os/exec"
    // "strings"
    "bytes"
    // "fmt"
    "log"
    // "os"
)

const ProxyTemplate = `{{ .EnvBasePath }} {{ .Port }} {{ .SiteBasePath }}`

type Proxy struct{
	Name string         `json:"name"`
	Port string         `json:"port"`
	EnvBasePath string  `json:"env_base_path"`
	SiteBasePath string `json:"site_base_path"`
}

var Proxies []Proxy

func main() {
    Proxies = make([]Proxy, 0)

    r := mux.NewRouter()
    r.HandleFunc("/", Hello).Methods("GET").Name("index")
    r.HandleFunc("/proxies", ProxyList).Methods("GET").Name("proxies.list")
    r.HandleFunc("/proxies", ProxyStore).Methods("POST").Name("proxies.store")
    r.HandleFunc("/proxies/{name}", ProxyDelete).Methods("DELETE").Name("proxies.delete")

    log.Fatalln(http.ListenAndServe(":8000", r))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func ProxyList(w http.ResponseWriter, r *http.Request) {

}

func ProxyStore(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var proxy Proxy
	err = json.Unmarshal(b, &proxy)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

/*
    tmpl, err := template.New("proxy").Parse(ProxyTemplate)
    http.Error(w, err.Error(), 500)
    tmpl.Execute(tmpl, &proxy)
*/
    Proxies = append(Proxies, proxy)

    /*
	f, err := os.Create("/etc/nginx/sites-enabled/"+proxy.Name+".conf")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
    */

    entries := writeEntries()

    /*
	cmd := exec.Command("service", "nginx", "restart")
	err = cmd.Run()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
    */

	w.WriteHeader(http.StatusOK)
    w.Write([]byte(entries))
	// fmt.Fprintf(w, "{ \"status\": \"success\" }\n")
}

func ProxyDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func writeEntries() string {
    var buf bytes.Buffer
    for proxy := range(Proxies) {
        tpl := template.Must(template.New("ProxyTemplate").Parse(ProxyTemplate))
        tpl.Execute(&buf, proxy)
        buf.Write([]byte("\n\t\t\t"))
    }
    return string(string(buf.Bytes()))
}
