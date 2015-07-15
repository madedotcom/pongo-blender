# pongo-blender
Renders pongo2 templates from environmental variables.

## How do I use this thing?

To use pongo blender you need to have a template, pongo blender will collect the environmental variables and will render the template using those variables. Pongo blender will output the rendered template to stdout. You can try pongo blender like this:

```
$ echo "{% if PONGO_ENV_VAR %} PONGO_ENV_VAR: {{ PONGO_ENV_VAR }} {% endif %}" > test_template
$ PONGO_ENV_VAR=testing go run p2.go test_template
 PONGO_ENV_VAR: testing
```

