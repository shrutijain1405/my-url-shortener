package helper

import (
    "math/rand"
    "time"
)

func GenerateShortUrl(length int) string {
    rand.Seed(time.Now().Unix())

    randomStr := make([]byte, length)

    // Generating Random string
    for i := 0; i < length; i++ {
        randomStr[i] = byte((97 + rand.Intn(25)))
    }

    // Displaying the random string
    str := string(randomStr)
    return (str)
}
