package main

import (
    "os/exec"
)

func main() {
    redisCmd := exec.Command("redis-server", "--daemonize", "yes")
    if err := redisCmd.Start(); err != nil {
        panic(err)
    }

    flaskCmd := exec.Command("flask", "run", "--host=0.0.0.0")
    if err := flaskCmd.Run(); err != nil {
        panic(err)
    }

    if err := redisCmd.Wait(); err != nil {
        panic(err)
    }
}
