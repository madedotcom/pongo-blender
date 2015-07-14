package main

import (
    "os"
    "strings"
    "text/template"
)

func getContext() map[string]string{
    env := os.Environ()
    ctx := make(map[string]string)
    for _, el := range env {
        split := strings.SplitN(el, "=", 2)
        ctx[split[0]] = split[1]
    }
    return ctx
}

func main() {
    t,err := template.New("conf").Parse("{{.LANGUAGE}}")
    if err != nil { panic(err) }
    m := getContext()
    t.Execute(os.Stdout, m)
}
