# Workshop :: Develop services
* REST APIs
* Documentations
  * Service catalog with [Backstage](https://backstage.io/)
  * [Swagger or OpenAPI](https://swagger.io/)
    * Design-first vs Code-first  
* Technology stack
  * Go + [Echo framework](https://github.com/labstack/echo)
  * Database :: MySQL

## Step 1 :: Design and Develop REST API with [OpenAPI/Swagger](https://swagger.io/)
* Design or API First
* Generate code from OpenAPI Specification
* [Swagger Editor](https://editor.swagger.io/)

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

## Step 3 :: Testing your APIs
* External testing
  * Postman and [newman](https://www.npmjs.com/package/newman)
* Internal testing
  * [net/httptest](https://pkg.go.dev/net/http/httptest)
  * [testify](https://github.com/stretchr/testify)


### Step 3.1 :: External testing
```
$cd postman

$npm install -g newman
$newman run item-service.postman_collection.json
```

### Step 3.2 :: Internal testing
Create file `impl_test.go`
```
package api_test

import (
	"api"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetItems(t *testing.T) {
	// Setup
	e := echo.New()
	server := api.NewServer()
	req := httptest.NewRequest(http.MethodGet, "/api/items", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, server.GetItems(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var items []api.Item
		err := json.Unmarshal(rec.Body.Bytes(), &items)
		assert.NoError(t, err)
		assert.Len(t, items, 2)
		assert.Equal(t, 1, items[0].Id)
		assert.Equal(t, "Item 1", items[0].Name)
		assert.Equal(t, 2, items[1].Id)
		assert.Equal(t, "Item 2", items[1].Name)
	}
}
```

Run test
```
$go mod tidy
$go test ./... -cover -v
```

## Step 4 :: Registering the API service with Backstage
* [Backstage](https://backstage.io/)
  * NodeJs 20
* [Reference](https://backstage.io/docs/features/software-catalog/descriptor-format)

Install Backstage
```
$npx @backstage/create-app@latest
$cd demo
$yarn install
$yarn start
```

Config file `app-config.yaml`
```
baseUrl: http://localhost:7007
reading:
    allow:
        - host: '*.githubusercontent.com'
```

Config API/Service/System in file `entities.yaml`
```
apiVersion: backstage.io/v1alpha1
kind: System
metadata:
  name: demo-system
  description: "Demo system for workshop"
spec:
  owner: team-a
---
apiVersion: backstage.io/v1alpha1
kind: Component
metadata:
  name: items-service
  description: "Go Echo service for Items API"
spec:
  type: service
  owner: team-a
  system: demo-system
  lifecycle: production
  providesApis: [items-api]
---
apiVersion: backstage.io/v1alpha1
kind: API
metadata:
  name: items-api
  description: "The Items service API"
spec:
  type: openapi
  lifecycle: production
  owner: team-a
  system: demo-system
  definition:
    $text: https://raw.githubusercontent.com/up1/workshop-develop-services-2025/main/openapi/openapi.yaml
```

Config User/Group in file `org.yaml`
```
# https://backstage.io/docs/features/software-catalog/descriptor-format#kind-user
apiVersion: backstage.io/v1alpha1
kind: User
metadata:
  name: user01
spec:
  memberOf: [team-a]
---
# https://backstage.io/docs/features/software-catalog/descriptor-format#kind-group
apiVersion: backstage.io/v1alpha1
kind: Group
metadata:
  name: team-a
spec:
  type: team
  children: []
```

Edit file `app-config.yaml`
* Allow to read data from external domain
```
backend:
  baseUrl: http://localhost:7007
  reading:
    allow:
      - host: '*.githubusercontent.com'
```

Start server again
```
$yarn start
```

Access to server
* http://localhost:3000

## Step 5 :: Building software with Docker
```
$docker compose build item-service
$docker compose up -d item-service
$docker compose ps
```

Testing again !!
* http://localhost:8080/items
