# Aufgabe 3: Nutzer-Tabelle - mit SQL!
## Kriterien
Erweitern sie ihre User-Middleware, sodass diese einen gesendetes Objekt in einer Datenbank persistiert

1. Wenn die Anfrage eine POST-Anfrage ist, dann speichern sie das gesendete Objekt in der Tabelle "Users"
2. Das erwartete JSON entspricht dem Schema `{name:"Dino", age:"27", isAdmin:true}`
3. Wenn die Anfrage eine GET-Anfrage ist, dann geben sie eine Tabelle mit allen Nutzern aus die in der Tabelle "Users" vorkommen

## Hilfestellung

Sie können sich an folgendem HTML orientieren um ihr Template zu erstellen (diesmal mit DB-ID)
```
<html>
<body>
    <h1>
        1 User(s)
    </h1>
    <table>
        <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Age</th>
            <th>Status</th>
        </tr>
        <tr>
            <td>1</td>
            <td>Dino Omanovic</td>
            <td>27</td>
            <td>Admin</td>
        </tr>
    </table>
<body>
</html>
```