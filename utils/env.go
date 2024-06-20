package utils



import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "strconv"
)


func EnvVariableStr(key string) (string) {

    var envPath string
    envPath = ".env"


    err := godotenv.Load(envPath)

    if err != nil {
        log.Fatalf("Error loading .env file", err)
    }

    v := os.Getenv(key)

    return v

}


func EnvVariableBool(key string) (bool) {
    s := EnvVariableStr(key)

    v, err := strconv.ParseBool(s)
    if err != nil {
        return false
    }
    return v
}


func EnvVariableInt64(key string) (int64) {
    s := EnvVariableStr(key)

    v, err := strconv.ParseInt(s, 10, 64)
    if err != nil {
        return -1
    }
    return v
}



