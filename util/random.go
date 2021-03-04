package util

import (
    "math/rand"
    "time"
    "strings"
    "fmt"
)

const alphabet  =  "abcdefghijklmnopqrstuvwxyz"

func init() {
    seed := time.Now().UnixNano()
    rand.Seed(seed)
}

func RandomInt(min, max int64) int64 {
    return min + rand.Int63n(max-min+1)
}

func RandomString(n int)string {
    var sb strings.Builder
    k:=len(alphabet)
    for i:=0;i<n;i++ {
        c := alphabet[rand.Intn(k)]
        sb.WriteByte(c)
    }
    return sb.String()
}

func RandomOwner() string {
    r :=  RandomString(6)
    return r
}

func RandomMoney() int64 {
    return RandomInt(0, 1000)
}

func RandomCurrency() string {
    currencies := []string{USD, EUR, CAD}
    n := len(currencies)
    return currencies[rand.Intn(n)]
}

func RandomEmail() string {
    return fmt.Sprintf("%s@email.com", RandomString(6))
}

