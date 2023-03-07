package main

import (
    "flag"
    "fmt"
    "os"
    "os/exec"
    "strings"
    "time"
)

func main() {
    cmd := flag.String("cmd", "", "Command to execute")
    interval := flag.Int("interval", 10, "Time interval between command executions in seconds")
    maxAttempts := flag.Int("max-attempts", 0, "Maximum number of command execution attempts (0 for infinite)")
    searchString := flag.String("search-string", "", "String to search for in the command output")
    flag.Parse()

    attempt := 0
    for *maxAttempts == 0 || attempt < *maxAttempts {
        attempt++
        command := exec.Command("bash", "-c", *cmd)
        output, err := command.Output()
        if err != nil {
            fmt.Println(err)
            os.Exit(1) // Exit with error code 1
        }
        outputStr := string(output)
        if strings.Contains(outputStr, *searchString) {
            fmt.Printf("Found %s\n", *searchString)
            os.Exit(0) // Exit with success code 0
        }
        fmt.Printf("%s not found. Waiting %d seconds before trying again...\n", *searchString, *interval)
        time.Sleep(time.Duration(*interval) * time.Second)
    }
    fmt.Printf("Maximum number of attempts reached (%d attempts)\n", *maxAttempts)
    os.Exit(1) // Exit with error code 1
}
