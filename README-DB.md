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