# pongo-blender
Renders pongo2 templates from environmental variables.
pongo2 is the successor of pongo, a Django-syntax like templating-language.

### READ: pongo-blender lets me use ~~jinja2~~ pongo2 templates inside docker containers and populate the values(secrets and configs) at container run time. One image, many containers, every container can have its own configs.

##pongo2 examples
https://github.com/flosch/pongo2/blob/master/template_tests/filters.tpl

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
For OSX you will also have to do
```
mkdir $GOPATH/bin
export GOBIN=$GOPATH/bin
```

## Dependencies
* gopkg.in/alecthomas/kingpin.v2
* github.com/flosch/pongo2

## Install pongo-blender
Install pongo-blender and dependencies.
```
cd $GOPATH/src && git clone https://github.com/madedotcom/pongo-blender && cd pongo-blender && go get && go install .
go build
```
You now have a static binary called `pongo-blender`, you can run it anywhere.
You can even copy it to `/bin/` .
You can also add this binary to your docker image.

## Using pongo-blender

To use pongo blender you need to have a template, pongo-blender will collect the environmental variables and will render the template using those variables. pongo-blender will output the rendered template to stdout. You can try pongo-blender like this:

* Create a `vars` file
```
vim vars
```
with something like this
```
export SHOES="green"
export CATS="cute"
export STATE="busy"
```

* Create a `template` file with something like this
 ```
my shoes are {{ SHOES }}
my cats are {{ CATS }}
my state is {{ STATE }}
```


* Run pongo-blender passing it values for variables you want set
```
. vars && pongo-blender template > output
```

* The output file is a complete with your variables and template stuff
```
egidijus@ub-sol:/tmp/pongos$ cat output 
my shoes are green
my cats are cute
my state is busy
```


