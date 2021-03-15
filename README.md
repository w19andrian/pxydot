# PXYDOT
# Table of Contents

  
- [PXYDOT](#pxydot)
- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
- [Implementation](#implementation)
  - [Quickstart](#quickstart)

# Introduction
`PXYDOT` *(pixie.dot)* is a lightweight DNS proxy. `PXYDOT` accept both TCP & UDP connections on standard DNS port(53) and forward the requests to the secured DNS-over-TLS(DoT) upstream servers on TCP/853. `PXYDOT` already implemented `round robin` load balancing if supplied with more than one upstream servers in the configuration file. Because of it's insecure nature, it is not recommended for a public facing DNS server deployment.
# Implementation
## Quickstart
To start using the app with default configuration , you can run by the following command:
```shell
make build
make run
```
OR

you can also run it with docker(Docker required):
```shell
make package
make docker-run
```
Below are all of the options you can run with make
| Command | Description|
|:--------|:-----------|
|make all |equivalent of "make build" and "make package"|
| make build| build binary file inside ./bin directory |
|make package | build docker image |
|make run | run the app with default configuration |
|make docker-run | run an interactive containerized app with default configuration |
|make clean | remove the binary file and docker image |

