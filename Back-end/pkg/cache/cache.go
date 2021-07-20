package cache

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var m map[string]string
var ctx = context.Background()

func init() {
	m = make(map[string]string)
}

func ServeJQueryWithCache(w http.ResponseWriter, key string) string {
	// start := time.Now()
	if v, ok := m[key]; ok {
		// fmt.Fprintln(w, "There are data in local cache")
		// fmt.Fprintln(w, fmt.Sprintf("Cache hit \n %v \n %v", time.Since(start), v))
		// fmt.Fprintln(w, v)
		return v
	} else {
		return ServeJQueryWithRemoteCache(w, key)
	}
	// fmt.Fprintln(w, ServeJQueryWithRemoteCache(w, key, data))
}
func ServeJQueryWithRemoteCache(w http.ResponseWriter, key string) string {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	val2, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		// fmt.Fprintln(w, "No data in remote cache")
		// fmt.Fprintln(w, "Data in database: ", data)
		return "No data in remote cache"
	} else if err != nil {
		// Neu Docker chua run thi goi den database
		return "No data in remote cache"
	} else {
		// fmt.Fprintln(w, "There are data in remote cache")
		m[key] = string(val2)
		return val2
	}
}
func InsertData(key, data string) string {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	m[key] = string(data)
	err1 := rdb.Set(ctx, key, data, 0).Err()
	if err1 != nil {
		return data
	}
	return data
}
func DeleteLocalCache(w http.ResponseWriter, key string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	delete(m, key)
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	} else {
		fmt.Fprintln(w, "Delete Local Cache seccessfully!")
	}
}
func DeleteRemoteCache(w http.ResponseWriter, key string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	} else {
		fmt.Fprintln(w, "Delete Remote Cache seccessfully!")
	}
}
