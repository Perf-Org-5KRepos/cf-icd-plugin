package otc

import (
    "fmt"
    "bytes"
    "net/http"
    "io/ioutil"
)

func Request(url string, method string, buf *bytes.Buffer) (string) {
    client := &http.Client{}
    req, err := http.NewRequest(method, url, buf)
    check(err)
    resp, err := client.Do(req)
    check(err)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    check(err)
    return string(body)
}

func check(e error) {
    if e != nil {
        fmt.Println("Warning: cf icd not successful, message from otc: %s",e)
        os.Exit(0)
//        panic(e)
    }
}
