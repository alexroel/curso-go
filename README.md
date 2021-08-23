# Curso de Go (Golang)
Este es el repositorio de curso de Go

## Instalar Go
Para instalar Go descargar el archivo de instalación de la página [https://golang.org/dl/](https://golang.org/dl/). 

## Configuración de Variables de Entorno en Windows 


- GOROOT: En este variable de entorno va la ruta de la instalación `C:\Program Files\Go`

- PATH: En este variable de entorno va la ruta de la instalación pero la ruta de la carpeta bin `C:\Program Files\Go\bin`

- GOPATH: En este variable de entorno va la ruta de tu espacio de trabajo por Ejemplo `C:\User\roel1\go`

- GOBIN: En este variable de entorno va la ruta de la carpeta bin que tienes que tener o crear dentro de tu espacio de trabajo, por ejemplo. `C:\User\roel1\go\bin` 

## Estructura de espacio de trabajo de GO

- BIN: Guarda todo los ejecutables que utilicemos o creemos.

- PKG: Guarda paquetes o librerias que vas a utilizar en tu proyecto 

- SRC: Aqui vas crear todo tu codigo o tus proyectos y tambien la librerias que vas a utilizar de terceros. 

```
C:\User\roel1\go\
    .bin
    .pkg
    .src
```

## Configuración de Variables de Entorno en Linux

Descargar el archivo para linux luego realiza la instalación con indica en la pagina, para darle permiso como administrado usa `sudo`. 

```
sudo tar -C /usr/local -xzf go1.16.7.linux-amd64.tar.gz
```

### Configuiración de Variables de entorno 
Abre el archivo oculto en la dirección de usuario que es `.bashrc` abre con algun editor y configura los siguientes variables de entorno al final de archivo. 

```
export GOPATH=/home/alexroel/workspace/go
export GOBIN=$GOPATH/bin
export GOROOT=/usr/local/go

export PATH=$PATH:$GOBIN:$GOROOT/bin
```