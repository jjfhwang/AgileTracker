// cmd/agiletracker/main.go
package main

import (
"flag"
"log"
"os"

"agiletracker/internal/agiletracker"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := agiletracker.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
