package utils

import (
    "os"
)

func CheckError(err error, error_message string) {
    if err != nil {
        println(error_message + ": ", err.Error())
        os.Exit(1)
    }
}
