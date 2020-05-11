# Temperature Humidity Station REST API

REST API written in Go that stores and provide temperature and humidity values from IoT sensor. It store the data in an SQLite database.

## REST endpoints

-   Get data

```
curl --location --request GET 'http://localhost:8080/api/v1' \
--header 'Content-Type: application/x-www-form-urlencoded'
```

-   Store new data

```
curl --location --request POST 'http://localhost:8080/api/v1' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'date=1589187325' \
--data-urlencode 'temperature=33.7' \
--data-urlencode 'humidity=55.0' \
--data-urlencode 'token=secret_token!'
```

## Libraries Used

-   [Gin Web Framework](https://github.com/gin-gonic/gin)
-   [GORM](https://gorm.io/)
-   [Java properties scanner](https://github.com/magiconair/properties)
