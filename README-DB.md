# Working with Database 
* MySQL

## Step 1 :: Start database
* Initial tables and data
```
$docker compose up -d db
$docker compose ps
```

## Step 2 :: Building item service with database
```
$docker compose build item-service
$docker compose up -d item-service
$docker compose ps
```

Testing again !!
* http://localhost:8080/items

## Step 3 :: Working with Distributed tracing
* [OpenTelemetry](https://opentelemetry.io/)
* [Jaeger](https://www.jaegertracing.io/)

Start Jaeger server
```
$docker compose up -d jaeger
$docker compose ps
```

Access to Jaeger
* http://localhost:16686

Start OTEL Collector
```
$docker compose up -d otel-collector
$docker compose ps
```

## Step 4 :: Add tracing in item service
* [GO](https://opentelemetry.io/docs/languages/go/)
* [Tracing with Echo](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho)
* [Tracing with MYSQL](https://github.com/XSAM/otelsql)
* [OpenTelemetry instrumentations for Go](https://github.com/uptrace/opentelemetry-go-extra)

```
$docker compose build item-service
$docker compose up -d item-service
$docker compose ps
```

Testing again !!
* http://localhost:8080/items

Access to Jaeger
* http://localhost:16686

## Step 5 :: Add metrics in item service
* [Prometheus](https://prometheus.io/)

```
$docker compose up -d prometheus
$docker compose ps
```

Build item service
```
$docker compose build item-service
$docker compose up -d item-service
$docker compose ps
```

Testing again !!
* http://localhost:8080/items

Access to Prometheus
* http://localhost:9090/
  * Search metric name = get_items_count_total


