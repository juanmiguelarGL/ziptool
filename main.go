package main

import (
    "fmt"
    "os"

    "github.com/juanmiguelarGL/ziptool/cmd" // Adjust this path according to your go.mod
)

// version, author, etc. (optional)
var (
    version = "v0.1.0"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
