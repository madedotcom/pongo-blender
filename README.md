# pongo-blender
Renders pongo2 templates from environmental variables.

## Requirements
* golang
* $GOPATH set

## Set your $GOPATH 
```
mkdir ~/.go
echo "GOPATH=$HOME/.go" >> ~/.bashrc
echo "export GOPATH" >> ~/.bashrc
echo "PATH=\$PATH:\$GOPATH/bin # Add GOPATH/bin to PATH for scripting" >> ~/.bashrc
source ~/.bashrc
```

## Dependancies
* gopkg.in/alecthomas/kingpin.v2
* github.com/flosch/pongo2

## Install pongo-blender
Install pongo-blender and dependancies.
```
cd $GOPATH/src && git clone https://github.com/madedotcom/pongo-blender && cd pongo-blender && go install .
```

## Using pongo-blender

To use pongo blender you need to have a template, pongo-blender will collect the environmental variables and will render the template using those variables. pongo-blender will output the rendered template to stdout. You can try pongo-blender like this:

* Create a template file called `test_template` and populate with some templated `{{variables}}`
```
cd ~
echo "{% if PONGO_ENV_VAR %} PONGO_ENV_VAR: {{ PONGO_ENV_VAR }} {% endif %}" > test_template
```

* Run pongo-blender passing it values for variables you want set
```
PONGO_ENV_VAR=testing go run $GOPATH/src/pongo-blender/pongo-blender.go ~/test_template
```

The STDOUT should be your template with populated variable values.
```
 PONGO_ENV_VAR: testing
```


