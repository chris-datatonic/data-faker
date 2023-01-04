# data-faker
Creates fake data and sends to pubsub

## Usage

`go run main.go --input "input.json" --num_values 200`

## Input Example

```json
[   
    {
        "name": "Name",
        "fakerType": "faker:\"name\"",
        "jsonType": "string"
    },
    {
        "name": "IsMarried",
        "jsonType": "bool"
    },
    {
        "name": "Email",
        "fakerType": "faker:\"email\"",
        "jsonType": "string"
    },
    {
        "name": "Age",
        "fakerType": "faker:\"oneof: 10, 20\"",
        "jsonType": "int"
    }
]
```

## Supported Tags

https://github.com/go-faker/faker/blob/v4.0.0-beta.4/faker.go#L29
