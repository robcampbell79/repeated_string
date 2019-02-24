package main

import(
    "fmt"
    "bufio"
    "os"
)

func repeatedString(s string, n int64) int64 {
    
    count := int64(0)

    for i := 0; i < len(s); i++ {
        if s[i] == 'a' {
            count++
        }
    }

    count = count * (n / int64(len(s)))

    for i := int64(0); i < n % int64(len(s)); i++ {
        if s[i] == 'a' {
            count++
        }
    }

    return count
}

func main() {

    var n int64
    fmt.Println("Enter string length and string with a space seprating them: ")

    fmt.Scan(&n)

    scanner := bufio.NewReader(os.Stdin)
    s, _ := scanner.ReadString('\n')

    result := repeatedString(s, n)

    fmt.Println(result)


}
