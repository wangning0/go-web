<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>文件上传</title>
</head>
<body>
    <form enctype="multipart/form-data" action="/upload" method="POST">
        <input type="file" name="uploadfile">
        <input type="hidden" name="token" value="{{.}}"> 
        <input type="submit" value="upload">
    </form>
</body>
</html>