package HTTPHandler

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// HTTPHandlerUtil implement interface
type HTTPHandlerUtil struct {
	populationhandler *PolutionHandler
}

// HTTPHandlerFactory creates a interface for HTTPHandlerUtil
func HTTPHandlerFactory() *HTTPHandlerUtil {
	thisHandler := new(HTTPHandlerUtil)
	thisHandler.populationhandler = PolutionHandlerFactory()
	return thisHandler
}

func (handler *HTTPHandlerUtil) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.Method == "GET" {
		if r.URL.Path == "/Ings" {
			m, err := url.ParseQuery(r.URL.RawQuery)
			if err != nil {
				log.Println("Failed to parse url: ", m["h"][0])
				return
			} else {
				log.Println("Requested day type: ", m["h"][0])
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				if _, ok := m["lat"]; ok {
					handler.populationhandler.ProcessRequest(w, m["h"][0], m["lat"][0], m["lon"][0])
				} else {
					handler.populationhandler.ProcessRequest(w, m["h"][0], "47.3872661", "8.5079584,13")
				}
			}
			return
		}
	}
	path := r.URL.Path
	if path == "/" {
		path = "/StartPage.html"
	}

	modifiedPath, contentType := handler.parseRequest(&path)
	data, err := ioutil.ReadFile(string(modifiedPath))
	log.Println(modifiedPath)
	if err == nil {
		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 not found"))
	}
}

func (handler *HTTPHandlerUtil) parseRequest(path *string) (string, string) {
	var contentType string
	modifiedSourcePath := "Webview" + *(path)
	if strings.HasSuffix(*path, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(*path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(*path, ".js") {
		modifiedSourcePath = "../../ViewModel" + *(path)
		contentType = "application/javascript"
	} else if strings.HasSuffix(*path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}
	return modifiedSourcePath, contentType
}
