package sol

import (
"net/http"
"strings"
	"io/ioutil"
	"fmt"
)
type AssetProxy struct {
	Delegate              http.Handler
	Suffixes              []string
}

func (p *AssetProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p.isAssetPath(r.URL.Path) && r.Method == http.MethodGet{
		pathElements := strings.Split(r.URL.Path, "/")
		fileName := pathElements[len(pathElements)-1]
		fileBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err.Error())
		}
		w.Write(fileBytes)
	} else if p.Delegate != nil {
		p.Delegate.ServeHTTP(w, r)
	}
}

func (p *AssetProxy) isAssetPath(path string) bool {
	path = strings.ToLower(path)
	for _, ending := range p.Suffixes {
		if strings.Contains(path, ending) {
			return true
		}
	}
	return false
}