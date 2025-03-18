# Workshop :: Develop services
* REST APIs
* Technology stack
  * Go + [Echo framework](https://github.com/labstack/echo)
  * Database :: MySQL

## Step 1 :: Design and Develop REST API with [OpenAPI/Swagger](https://swagger.io/)
* Design or API First
* Generate code from OpenAPI Specification

Use [OAPI-CodeGen](https://github.com/oapi-codegen/oapi-codegen)
```
$go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
$ oapi-codegen -version
```

Generate code 
```
$cd item-service
$go mod init api
$oapi-codegen --config=../openapi/config.yaml ../openapi/openapi.yaml
$go mod tidy
```

Run service
```
$go run cmd/main.go
```

Testing APIs
```
$curl http://localhost:8080/items | jq
```

## Step 2 :: Generate OpenAPI Documentation
* [Redoc](https://github.com/Redocly/redoc)
* [Reference](https://github.com/up1/workshop-api-first/tree/main/workshop/swagger)

```
$npx @redocly/cli build-docs ./openapi/openapi.yaml

// Open file
* redoc-static.html
```

Use [redocly cli](https://redocly.com/docs/cli)
```
$npm i -g @redocly/cli 
$redocly lint ./openapi/openapi.yaml
$redocly build-docs ./openapi/openapi.yaml

// Open file
* redoc-static.html
```
