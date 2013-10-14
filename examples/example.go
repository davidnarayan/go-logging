package main

import (
    "github.com/davidnarayan/go-logging"
)

func main() {
    logging.SetLevel(logging.DEBUG)

    logging.Debug("test message")
    logging.Info("test message")
    logging.Warn("test message")
    logging.Error("test message")
    logging.Fatal("test message")
    logging.Stats("test message")
}
