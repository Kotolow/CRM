# CRMGo

build project:
```shell
go build -o crm ./cmd/crm
```
```shell
./crm
```

Generate swagger docs:
```shell
swag init -g ./cmd/crm/main.go -o ./docs --parseDependency
```

http://localhost:8080/docs/index.html#/

Date format:
"2024-09-26T10:00:00Z"

http://localhost:8025/ - проверка почты

## Поля для new_branch:
 - branch_name
 - source_branch

Нейминг веток: Название-задачи-название_ветки

Пример: H&h-1-temp

## Поля для new_branch:
- branch_name
- source_branch




## Создание проекта
```json
{
"code": "H&h",
"created_by": 2,
"description": "Hackaton project",
"name": "Horns&Hooves"
}
```

## Создание таски
```json
{
  "assigned_to": 1,
  "description": "Нужно сделать авторизацию",
  "due_date": "2024-09-28T00:00:00Z",
  "priority": "blocker",
  "project_id": 1,
  "status": "open",
  "time_spent": 120,
  "title": "Авторизация"
}
```

## Создание ветки
```json
{
  "branch_name": "H&h-1-backend",
  "source_branch": "main"
}
```

## Создание пр-а
```json
{
  "body": "Test body",
  "source_branch": "Backend",
  "target_branch": "main",
  "title": "H&h-1:Test PR"
}
```
