<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>Hello Golang</title>
</head>
<body>

    <p>User Index</p>
    {{.}}
    {{range .}}
    
    <a href="/user/edit/{{.ID}}">{{.Name}}</a>

    {{end}}
    
    <p><a href="/user/new">new</a></p>
    
</body>
</html>