package main

import (
    "github.com/gorilla/mux"
    "encoding/json"
    "text/template"
    "io/ioutil"
    "net/http"
    "os/exec"
    "bytes"
    "log"
    "fmt"
    "os"
)

const ProxyTemplate = `location {{ .EnvBasePath }} { http://localhost:{{ .Port }}{{ .SiteBasePath }}; }`

type Proxy struct{
	Port string         `json:"port"`
	EnvBasePath string  `json:"env_base_path"`
	SiteBasePath string `json:"site_base_path"`
}

var Proxies map[string]*Proxy

func main() {
    Proxies = make(map[string]*Proxy)

    r := mux.NewRouter()
    r.HandleFunc("/", Hello).Methods("GET").Name("index")
    r.HandleFunc("/proxies", ProxyList).Methods("GET").Name("proxies.list")
    r.HandleFunc("/proxies", ProxyStore).Methods("POST").Name("proxies.store")
    r.HandleFunc("/proxies", ProxyDelete).Methods("DELETE").Name("proxies.delete")

    log.Fatalln(http.ListenAndServe(":8000", r))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func ProxyList(w http.ResponseWriter, r *http.Request) {
    list, err := json.Marshal(ProxyList)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    fmt.Fprintf(w, string(list))
}

func ProxyStore(w http.ResponseWriter, r *http.Request) {
	proxy, err := getProxyFromRequest(r)
    if(err != nil) {
        http.Error(w, err.Error(), 500)
        return
    }

    Proxies[proxy.EnvBasePath] = &proxy

    tmpl, err := template.ParseFiles("/root/.nginx.default.template")
    if(err != nil) {
        http.Error(w, err.Error(), 500)
        return
    }

    entries := writeEntries()
    replacement := struct{
        ProxyList string
    }{
        ProxyList: entries,
    }

    file, err := os.OpenFile("/etc/nginx/sites-available/default", os.O_RDWR, 0644)
    if(err != nil) {
        http.Error(w, err.Error(), 500)
        return
    }
    defer file.Close()

    err = tmpl.Execute(file, replacement)
    if(err != nil) {
        http.Error(w, err.Error(), 500)
        return
    }

	cmd := exec.Command("service", "nginx", "restart")
	err = cmd.Run()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	fmt.Fprintf(w, "{ \"status\": \"success\" }\n")
}

func ProxyDelete(w http.ResponseWriter, r *http.Request) {
    proxy, err := getProxyFromRequest(r)
    if(err != nil) {
        http.Error(w, err.Error(), 500)
        return
    }

	delete(Proxies, proxy.EnvBasePath)

    fmt.Fprintf(w, "{ \"status\": \"success\" }\n")
}

func getProxyFromRequest(r *http.Request) (Proxy, error) {
    // Read body
    b, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        return Proxy{}, err
    }

    // Unmarshal
    var proxy Proxy
    err = json.Unmarshal(b, &proxy)
    if err != nil {
        return Proxy{}, err
    }

    return proxy, nil
}

func writeEntries() string {
    var buf bytes.Buffer
    for _, proxy := range(Proxies) {
        tpl := template.Must(template.New("ProxyTemplate").Parse(ProxyTemplate))
        tpl.Execute(&buf, proxy)

        buf.Write([]byte("\n\t\t\t"))
    }
    return string(buf.Bytes())
}
