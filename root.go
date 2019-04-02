package mos

import "os"

func Getenv(key string, def string) string {
    val := os.Getenv(key)
    if val != "" {
       return val
    }
    return def
}
