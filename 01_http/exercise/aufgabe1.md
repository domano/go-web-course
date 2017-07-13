# Aufgabe 1: Fileserver
## Kriterien
Es soll ein Fileserver entstehen welcher

1. bei einem HTTP-POST den Inhalt des HTTP-Requests in eine Datei speichert
2. bei einem HTTP-GET eine Datei ausliest und die den Body der HTTP-Response schreibt

Der URL-Pfad soll hierbei als Dateiname genutzt werden

## Tipps
* Datei lesen `ioutil.ReadFile(...)`
* Ordner erstellen `os.MkdirAll(...)`
* Datei schreiben `ioutil.WriteFile(..)`
* Post Request absetzen `http.Post("http://localhost:10000", "application/json", bytes.NewReader(body))`
* oder Curl: `curl -X POST --data "some-text" http//localhost:port`