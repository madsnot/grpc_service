# grpc_service
Test task for TAGES

```
git clone https://github.com/madsnot/grpc_service.git
```

## Сборка проекта и запуск
В проекте используются сторонние библиотеки поэтому необходимо их загрузить:
```
go mod download
```
Сборка проекта:
```
go build
```
Запуск проекта:
```
grpc_service
```

## Маршрутизация

### UploadImage (stream rpc)
От клиента получает пакеты байт изображения. Загружает файл на сервер.

Запрос в Postman:
```json
{
    "image":{
        "info":{
            "name":"test4",
            "format":".jpg",
            "createDate":"",
            "updateDate":""
            },
        "data":"/9j/4AAQSkZJRgABAQEASABIAAD/2wBDAAYEBQYFBAYGBQYHBwYIChAKCgkJChQODwwQFxQYGBcUFhYaHSUfGhsjHBYWICwgIyYnKSopGR8tMC0oMCUoKSj/2wBDAQcHBwoIChMKChMoGhYaKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCj/wgARCAMAAwADASIAAhEBAxEB/..."
    }
}
```

Ответ: статус-код `0 OK`

### DownloadImage (stream rpc)
Скачивает с сервера изображение. Отдаёт клиенту пакеты байт запрошенного файла.

Запрос в Postman:
```json
{
    "format": ".jpg",
    "name": "test4"
}
```

Ответ: статус-код `0 OK` и изображение
```json
{
    "image":{
        "info":{
            "name":"test4",
            "format":".jpg",
            "createDate":"2023-02-09",
            "updateDate":"2023-02-09"
            },
        "data":"/9j/4AAQSkZJRgABAQEASABIAAD/2wBDAAYEBQYFBAYGBQYHBwYIChAKCgkJChQODwwQFxQYGBcUFhYaHSUfGhsjHBYWICwgIyYnKSopGR8tMC0oMCUoKSj/2wBDAQcHBwoIChMKChMoGhYaKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCj/wgARCAMAAwADASIAAhEBAxEB/..."
    }
}
```

### GetImagesList (unary rpc)
Выводит список изображений на сервере в формате "Имя файла | Дата создания | Дата обновления".

Запрос в Postman:
```json
{}
```

Ответ: статус-код `0 OK` и список
```json
{
    "list": [
        "test4.jpg | 2023-02-09 | 2023-02-09"
    ]
}
```