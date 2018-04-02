beehive
------------

A cli tool to watch log information

# Installation

```shell
$ go get github.com/wonsikin/beehive/cmd/beehive
```

# Usage

### Help

```shell
$ beehive --help

  NAME:
   beehive - eat log messages and display it

  USAGE:
   beehive [global options] command [command options] [arguments...]

  VERSION:
   v0.1.0

  COMMANDS:
     init     create a configuration file
     help, h  Shows a list of commands or help for one command

  GLOBAL OPTIONS:
   --role value, -r value  running as 'queen' or 'worker' (default: "queen")
   --port value, -p value  listen port of beehive (default: 13000)
   --config FILE, -c FILE  load configuration from FILE, default value is ./beehive-queen.conf.yaml when running as the queen role and is ./beehive-worker.conf.yaml when running as worker role
   --help, -h              show help
   --version, -v           print the version
```

### Init configuration file

```
$ beehive init

  Use the arrow keys to navigate: ↓ ↑ → ←
  ? Generate configuration for which role?:
  ▸ queen
    worker
```

### Running as the queen server

```
$ beehive -r queen

  2018-04-02 18:53:13 info src/db/mongo/connect.go:22 connect to mongodb successfully
  2018-04-02 18:53:13 info cmd/beehive/app.go:120 beehive is served at :13000
```

### Running as the worker

```
$ beehive -r worker

  Worker arturo-n1.local is running
  2018/04/02 18:55:58 Seeked ./log.log - &{Offset:0 Whence:2}
```


# TODO

* [x] merge two cli into one
* [ ] supported multi log sources
* [ ] management platform
* [ ] workers' health report
* [ ] set log level and log path parameter
