# go-web-course
## Das net/http Package
* [Ein minimaler Webserver](01_http/minimal_serve_http.go)

### Handler
* [Ein einfacher Handler](01_http/handler_example.go)

### Server
* [Ein Minimaler Server](01_http/server_example.go)
* [Ein Server mit verschiedenen Routen](01_http/serve_mux_example.go)
* [Ein Handler mit verschiedenen Routen](01_http/handler.go)

### HTTPS
* [TLS zu nutzen ist ganz einfach](01_http/tls_serve_http.go)

### Wir schreiben einen HTTP-Server!
* [Schreiben wir einen Fileserver](01_http/exercise/aufgabe1.md)

## Erweitertes HTTP-Handling
### Selber Anfragen stellen
* [Einfache Wege eigene Anfragen zu stellen](02_http_extended/client/simple_client.go)
* [Aber man kann auch Requests zusammenbauen](02_http_extended/client/client_example.go)

### Plugins schreiben: Logging-Middleware als Beispiel
* [Plugins zu schreiben ist ganz einfach](02_http_extended/chaining/log_middleware.go)

### Wir schreiben eine Asset-Middleware!
* [Schreiben wir ein Plugin!](02_http_extended/exercise/aufgabe2.md)

## Templating in go
### Das Package text/template
### Das Package html/template
### Template aus dem Dateisystem laden
### Beispiel: Wert aus dem Request in HTML Seite einfügen

## Das database/sql Package
### Driver laden
### Exec
### Query
### Statements
### Beispiel: Template mit Werten aus der Datenbank befüllen
### Wir benutzen eine Datenbank! Schreibe einen Wert in die Datenbank und zeige eine Liste  der Werte an.
### Ausflug: gorm - ORM mit go!

## go-bindata - Assets ausliefern leicht gemacht
