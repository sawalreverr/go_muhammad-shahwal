# Echo Golang

## What is Echo ?

Echo merupakan salah satu framework GoLang yang high performance, extensible, minimalist Go web framework. Framework Echo memiliki sintaks yang minimalis, fitur yang cukup lengkap dan pastinya lebih mudah dipahami oleh pemula. Echo juga memiliki Forum & Chat untuk para developer.

Kelebihan menggunakan Echo framework ialah:

- Optimized Router
- Data Rendering
- Data Binding
- Middleware
- Scalable

## Get started using Echo

Install

```bash
$ mkdir myapp && cd myapp
$ go mod init myapp
$ go get github.com/labstack/echo/v4
```

example use

```go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

example routing

```go
e.GET("/", welcome)
e.GET("/users/:id", getUser)
e.POST("/users/form", saveUserByForm)
e.POST("/users/json", saveUserByJson)
e.PUT("/users/:id", updateUser)
e.DELETE("/users/:id", deleteUser)
e.POST("/upload", fileUpload)
```

example controller from routing above

```go
type User struct {
    Name string `json:"name" xml:"name" form:"name" query:"name"`
    Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func getUser(c echo.Context) error {
    id: = c.Param("id")
    team: = c.QueryParam("team")
    return c.String(http.StatusOK, "id: " + id + ", team: " + team)
}

func saveUserByForm(c echo.Context) error {
    name: = c.FormValue("name")
    return c.String(http.StatusOK, "name: " + name)
}

func saveUserByJson(c echo.Context) error {
    u: = new(User)
    if err: = c.Bind(u);err != nil {
        return err
    }
    return c.JSON(http.StatusCreated, u)
}
```
