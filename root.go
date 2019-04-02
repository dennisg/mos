package mos

import "os"
import "strconv"

func Getenv(key string, def string) string {
    val := os.Getenv(key)
    if val != "" {
       return val
    }
    return def
}

func GetInt(key string, def int) int {
    val, err := strconv.Atoi(Getenv(key, strconv.Itoa(def)))
    if err != nil {
        panic(err)
    }
    return val
}
