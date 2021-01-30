package main

import (
    b64 "encoding/base64"
    "fmt"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    dat, err := ioutil.ReadFile("./1610120402_recording.tar.gz")
    check(err)
    sEnc := b64.StdEncoding.EncodeToString([]byte(string(dat)))
    fmt.Print(sEnc)
}
