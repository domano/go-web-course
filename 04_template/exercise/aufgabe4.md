# Aufgabe 3: Nutzer-Tabelle
## Kriterien
Erweitern sie ihre JSON-Middleware, sodass diese einen gesendetes Objekt als Teil einer Liste anzeigt.

1. Wenn die Anfrage eine POST-Anfrage ist, dann speichern sie das gesendete Objekt in einer Liste
2. Das erwartete JSON entspricht dem Schema `{name:"Dino", age:"27", isAdmin:true}`
3. Wenn die Anfrage eine GET-Anfrage ist, dann geben sie eine Tabelle mit allen Nutzern aus

## Hilfestellung

Sie können sich an folgendem HTML orientieren um ihr Template zu erstellen
```
<html>
<body>
    <h1>
        1 User(s)
    </h1>
    <table>
        <tr>
            <th>Name</th>
            <th>Age</th>
            <th>Status</th>
        </tr>
        <tr>
            <td>Dino Omanovic</td>
            <td>27</td>
            <td>Admin</td>
        </tr>
    </table>
<body>
</html>
```