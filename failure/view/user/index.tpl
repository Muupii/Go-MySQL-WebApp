<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>Hello Golang</title>
</head>
<body>

    <p>User Index</p>
    <p><a href="/user/new">new</a></p>
    {{.User}}
    {{range .}}
    
    <a href="/user/edit/{{.ID}}">{{.Name}}</a>

    {{end}}
    
    
    
</body>
</html>