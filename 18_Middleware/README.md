# Middleware

## What is Middleware

Middleware adalah komponen software yang menjembatani aplikasi, data, dan pengguna.
Middleware membantu menyederhanakan pengembangan aplikasi dengan menangani tugas-tugas umum seperti transformasi data, komunikasi, dan keamanan.

Dengan middleware, developer dapat bekerja secara lebih efisien karena mereka dapat berfokus pada pengembangan fitur-fitur software, dan bukannya konektivitas antar komponen yang berbeda.

Contoh middleware yang sering dipakai:

- Authentication
- Authorization
- Validasi input
- Sanitasi input
- Response handler
- Data logger

### Why Middleware Important ?

- Akses informasi real-time
- Peningkatan efisiensi
- Mempertahankan integritas data
- Mempertahankan integritas data
- Pemberdayaan developer

## Echo Middleware

Echo sudah menyediakan middleware bawaan mereka secara default, berikut merupakan beberapa tipe echo middleware:

### Dieksekusi sebelum router memproses request (Echo#Pre)

- HTTPSRedirect
- HTTPSWWWRedirect
- WWWRedirect
- AddTrailingSlash
- RemoveTrailingSlash
- MethodOverride
- Rewrite

### Dieksekusi sesudah router memproses request dan mempunyai akses penuh ke echo.Context API (Echo#Use)

- BodyLimit
- Logger
- Gzip
- Recover
- BasicAuth
- JWTAuth
- Secure
- CORS
- Static

More information [echo middleware](https://echo.labstack.com/docs/category/middleware)

## Implementation

### Log Middleware

Logger middleware berfungsi mencatat informasi tentang setiap request HTTP yang diterima.

#### Old Logger middleware

Usage

```go
e.Use(middleware.Logger())
```

Sample output

```bash
{"time":"2017-01-12T08:58:07.372015644-08:00","remote_ip":"::1","host":"localhost:1323","method":"GET","uri":"/","status":200,"error":"","latency":14743,"latency_human":"14.743µs","bytes_in":0,"bytes_out":2}
```

#### Custom Configuration

```go
e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
  Format: "method=${method}, uri=${uri}, status=${status}\n",
}))

// Example above uses a Format which logs request method and request URI.
```

Sample output

```bash
method=GET, uri=/, status=200
```

### Auth Middleware

Salah satu Auth middleware yang ada di echo adalah Basic Auth

Usage

```go
e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// Be careful to use constant time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
		return true, nil
	}
	return false, nil
}))
```

Custom Configuration

```go
BasicAuthConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Validator is a function to validate BasicAuth credentials.
  // Required.
  Validator BasicAuthValidator

  // Realm is a string to define realm attribute of BasicAuth.
  // Default value "Restricted".
  Realm string
}

e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{}))
```

### JWT Middleware

JWT menyediakan middleware otentikasi JSON Web Token (JWT). Middleware Echo JWT terletak di https://github.com/labstack/echo-jwt/v4

Usage

```go
e.Use(echojwt.JWT([]byte("secret")))
```

Custom Configuration

```go
e.Use(echojwt.WithConfig(echojwt.Config{
  // ...
  SigningKey:             []byte("secret"),
  // ...
}))
```

More implementation [docs jwt echo](https://echo.labstack.com/docs/cookbook/jwt)
