package chaining

import (
	"github.com/Sirupsen/logrus"
	"net/http"
)

type LoggingMiddleware struct {
	Delegate http.Handler
}

func (p *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logrus.StandardLogger().
		WithField("Method", r.Method).
		WithField("Path", r.URL.Path).
		Info("Access!")
	if p.Delegate != nil {
		p.Delegate.ServeHTTP(w, r)
	}
}

func main() {
	chain := &LoggingMiddleware{Delegate: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Das wurde geloggt!"))
	})}

	panic(http.ListenAndServe(":8080", chain))
}
