# Endpoints API

## 1. Create To-Do
**POST** `/todos`

### Request Body:
```
{
  "title": "Turu",
  "note": "Mengarungi pulau kapuk",
  "deadline": "2024-11-30T23:59:59Z",
  "completed": false
}
```
### Response
```
{
    "ID": 26,
    "Title": "Turu",
    "Note": "Mengarungi pulau kapuk",
    "Completed": false,
    "Deadline": "2024-11-30T23:59:59Z",
    "CreatedAt": "2024-11-23T00:34:39.016+07:00",
    "UpdatedAt": "2024-11-23T00:34:39.016+07:00"
}
```

## 2. Get All To-Do
**GET** `/todos`
### Response
```
[
    {
        "ID": 22,
        "Title": "Belajar ORM di golang pakai GORM",
        "Note": "Pelajari migrasi, seeder, dan CRUD",
        "Completed": false,
        "Deadline": "2024-12-03T00:10:21.202+07:00",
        "CreatedAt": "2024-11-23T00:10:21.207+07:00",
        "UpdatedAt": "2024-11-23T00:10:21.207+07:00"
    },
    ...
    {
        "ID": 26,
        "Title": "Turu",
        "Note": "Mengarungi pulau kapuk",
        "Completed": false,
        "Deadline": "2024-12-01T06:59:59+07:00",
        "CreatedAt": "2024-11-23T00:34:39.016+07:00",
        "UpdatedAt": "2024-11-23T00:34:39.016+07:00"
    }
]
```

## 3. Update To-Do
**PUT** `/todos/:id`
### Request Body
```
{
  "title": "Turu Edit",
  "note": "Update dari note sebelumnya",
  "deadline": "2024-12-01T23:59:59Z",
  "completed": true
}
```
### Response
```
{
    "ID": 26,
    "Title": "Turu Edit",
    "Note": "Update dari note sebelumnya",
    "Completed": true,
    "Deadline": "2024-12-01T23:59:59Z",
    "CreatedAt": "2024-11-23T00:40:35.205+07:00",
    "UpdatedAt": "2024-11-23T00:40:35.205+07:00"
}
```

## 4. Delete To-Do
**DELETE** `/todos/:id`
### Response
```
{
    "message": "Todo deleted successfully"
}
```
