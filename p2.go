package main

import (
    "os"
    "strings"
    "text/template"
    "gopkg.in/alecthomas/kingpin.v2"
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


var (
    templatePath = kingpin.Arg("template", "Template.").Required().String()
)

func main() {
    kingpin.Parse()

    t, err := template.ParseFiles(*templatePath)
    if err != nil { panic(err) }
    m := getContext()
    t.Execute(os.Stdout, m)
}
