//Base64 is an encoding scheme, NOT encryption.
// It converts binary data → text

// Output uses only 64 safe characters
//A–Z  a–z  0–9  +  /


// Encoding = reversible

// Encryption = secret + secure

// Hashing = irreversible

package main

import (
    b64 "encoding/base64"
    "fmt"
)

func main() {

    data := "abc123!?$*&()'-=@~"

    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}