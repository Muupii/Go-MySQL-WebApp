<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>Edit</title>
</head>
<body>

    <form action="/user/update/{{.User.ID}}">
    
        <label >

            Name

            <input type="text" name="Name" value="{{.User.Name}}">

        </label>

        <input type="submit" name="Submit" value="更新">

        <p><a href="/user/delete/{{.User.ID}}">delete</a></p>
        
    
    </form>
    {{.Mess}}
    <p><a href="/user/index">Back</a></p>
</body>
</html>