# Aufgabe 2: Asset-Middleware
## Kriterien
Es soll ein Plugin geschrieben werden welches Bilder ausliefern kann.

1. Wenn die Anfrage keine GET-Anfrage ist, dann soll delegiert werden
2. Wenn der Pfad mit ".jpg" endet, dann soll ein Bild ausgeliefert werden.

Der URL-Pfad soll hierbei als Dateiname genutzt werden

## Tipps
* Datei lesen `ioutil.ReadFile(...)`