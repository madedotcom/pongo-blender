package main

import (
    "os"
    "strings"
    "gopkg.in/alecthomas/kingpin.v2"
    "github.com/flosch/pongo2"
)

func getContext() map[string]interface{}{
    env := os.Environ()
    ctx := make(map[string]interface{})
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
    t := pongo2.Must(pongo2.FromFile(*templatePath))

    err := t.ExecuteWriter(getContext(), os.Stdout)
    if err != nil {
        panic(err)
    }
}
