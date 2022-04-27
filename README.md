# Redis Operator

Redis Operator is a helper, manage redis connection and provide some useful operations.

## Install

```
go get github.com/KL-Engineering/ro
```

## Usage

Init options before use

```
import "github.com/KL-Engineering/ro"

ro.SetConfig(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Username: "",
	Password: "",
	DB:       0,
})
```

Then use `ro.MustGetRedis` to get `redis.Client`
