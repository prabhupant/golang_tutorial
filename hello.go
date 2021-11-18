package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
    if name == "" {
        return englishHelloPrefix + "world"
    }

    return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return
}

func main() {
    fmt.Println(Hello("Prabhu", "Spanish"))
    fmt.Println(Hello("Prabhu", "English"))
    fmt.Println(Hello("", ""))
}

