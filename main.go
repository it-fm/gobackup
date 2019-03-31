package main

import (
    "github.com/it-fm/gozip"
    "github.com/robfig/cron"
    "fmt"
    "math/rand"
    "runtime"
    "strconv"
)

func main() {

    go func() {
        c := cron.New()
        //start backup 30 in 30 minutes
        c.AddFunc("30 * * * *", func() { backup()})
        c.Start()
    }()

    runtime.Goexit()
}

func backup() {
    number := rand.Intn(500 - 1) + 1
    fmt.Println("start backup", number)
    gozip.Zip("./folderBackup", strconv.Itoa(number) + "backupResult.zip")
}
