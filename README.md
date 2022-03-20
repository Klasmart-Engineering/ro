# Redis Operator

Redis Operator is a helper, manage redis connection and provide some useful operations.


## Install

```
go get gitlab.badanamu.com.cn/calmisland/ro
```

## Usage

Init options before use
```
import "gitlab.badanamu.com.cn/calmisland/ro"

ro.SetConfig(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Username: "",
	Password: "",
	DB:       0,
})
```

Then use `ro.MustGetRedis` to get `redis.Client`