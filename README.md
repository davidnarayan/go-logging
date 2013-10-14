go-logging
==========

Package logging provides a simple leveled logging API


Installation
============

```bash
$ go get github.com/davidnarayan/go-logging
```

Quick Start
===========

```go
package main

import (
    "os"
    "github.com/davidnarayan/go-logging"
)

func main() {
    // Basic usage
    logging.Debug("This is a debug message")
    logging.Info("This is an info message")
    logging.Warn("This is a warning message")
    logging.Error("This is an error message")

    // Note: Fatal() calls os.Exit(1)
    // logging.Fatal("This is a fatal message")

    // ... with format strings
    foo := "bar"
    logging.Info("This is an info message with a parameter foo=%s", foo)

    // Change the output destination
    logging.Info("This message goes to STDERR")
    logging.SetWriter(os.Stdout)
    logging.Info("This message goes to STDOUT")

    // Change the logging level
    logging.SetLevel(logging.ERROR)
    logging.Info("This message won't be displayed")
    logging.Error("This message will be displayed")
}
```
