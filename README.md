# A very Simple GO Lang REST API

> TODO : use database instead of slice

## Quick Start


``` bash
# Install mux router
go get -u github.com/gorilla/mux
```

``` bash
go build
./Main
```

## Endpoints

### Get All Movies
``` bash
GET movies
```
### Get Single Movie

``` bash
GET movies/{id}
```

### Delete Movie

``` bash
DELETE movies/{id}
```

### Create Movie

``` bash
POST movies/

# Request sample
{
	"title": "Mari Pulang",
	"rating": "5.6",
	"year": 2019,
	"actor": {
		"name": "Rahmatullah",
		"age": 22
	}
}
```

### Update Movie

``` bash
PUT movies/{id}

# Request sample
{
	"title": "Mari Pulang Bersama",
	"rating": "5.6",
	"year": 2019,
	"actor": {
		"name": "Rahmatullah aka matx",
		"age": 22
	}
}

```
### Author

Muh Rahmatullah
