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

### Aufgabe1 : Wir schreiben einen HTTP-Server!
* [Schreiben wir einen Fileserver](01_http/exercise/aufgabe1.md)

## Erweitertes HTTP-Handling
### Selber Anfragen stellen
* [Einfache Wege eigene Anfragen zu stellen](02_http_extended/client/simple_client.go)
* [Aber man kann auch Requests zusammenbauen](02_http_extended/client/client_example.go)

### Plugins schreiben: Logging-Middleware als Beispiel
* [Plugins zu schreiben ist ganz einfach](02_http_extended/chaining/log_middleware.go)

### Aufgabe2 : Wir schreiben eine Asset-Middleware!
* [Schreiben wir ein Plugin!](02_http_extended/exercise/aufgabe2.md)

## JSON mit Go
### Json lesen
* [Einfaches lesen](03_json/json_marshal.go)
* [Lesen größerer Mengen](03_json/json_decoder.go)

### Objekte serialisieren
* [Einfaches schreiben](03_json/json_unmarshal.go)

### Aufgabe3 : JSON-Middleware
* [Wir wollen JSON als Daten-Format für unseren Handler nutzen](03_json/exercise/aufgabe3.md)

## Templating in go
### Das Package text/template
* [Beispiel für text Templates](04_template/template_example.go)

### Das Package html/template
* [Beispiel für html Templates](04_template/html_template_example.go)

### Templating-Server
* [Templating in einem HTTP-Server](04_template/templating_server.go)

### Aufgabe4 : User-Middleware
* [Unser JSON-Handler soll jetzt auch eine schöne Liste liefern](04_template/exercise/aufgabe4.md)

## Das database/sql Package
### Exec & Query
* [SQL-Anweisungen ausführen und Queries gegen die Datenbank absetzen](05_sql/simple_example.go)

### Statements
* [SQL-Anweisungen vorbereiten um Overhead zu vermeiden](05_sql/statement_example.go)

### Transaktionen
* [Transaktionen nutzen um atomare Änderungen zu vollziehen](05_sql/transaction_example.go)

### Aufgabe4 : User-Middleware
* [Unser User-Handler soll jetzt eine echte Datenbank nutzen](05_sql/exercise/aufgabe5.md)

### Ausflug: gorm - ORM mit go!
* [OR-Mapping mit Go (leider nicht mit Windows heute)](05_sql/gorm_example.md)

## go-bindata - Assets ausliefern leicht gemacht
Optional, bei Bedarf
