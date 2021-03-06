## <img src="https://raw.githubusercontent.com/hico-horiuchi/ohgi/master/logo.png" alt="logo.png" width="60" height="29"> ohgi v0.5.2

[![GoDoc](https://godoc.org/github.com/hico-horiuchi/ohgi/ohgi?status.svg)](https://godoc.org/github.com/hico-horiuchi/ohgi/ohgi) [![Circle CI](https://circleci.com/gh/hico-horiuchi/ohgi.svg?style=shield)](https://circleci.com/gh/hico-horiuchi/ohgi)

![screen.gif](https://raw.githubusercontent.com/hico-horiuchi/ohgi/master/screen.gif)

#### Requirements

  - [Golang](https://golang.org/) >= 1
  - [Sensu](http://sensuapp.org/) >= 0.19

#### Documents

  - [github.com/hico-horiuchi/ohgi/ohgi](http://godoc.org/github.com/hico-horiuchi/ohgi/ohgi) 

#### Installation

    $ git clone git://github.com/hico-horiuchi/ohgi.git
    $ cd ohgi
    $ make gom link
    $ sudo make install

#### Configuration

For example:

    {
      "datacenters": [{
        "name": "server-1",       // Required
        "host": "192.168.11.10",  // Required
        "port": 4567,             // Required
        "user": "sensu-1",        // Optional
        "password": "password"    // Optional
      }, {
        "name": "server-2",
        "host": "192.168.11.20",
        "port": 4567
      }]
    }

Specify a configuration file by `-C` ( `--config` ) option as below.  
If a configuration file is not specified, use `~/.ohgi.json` .

    $ ohgi -C ohgi.json events

Specify a datacenter by `-x` ( `--datacenter` ) option as below.  
If a datacenter is not specified, use first of `datacenters` .

    $ ohgi -x server-1 events

#### Usage

    Sensu command-line tool by Golang
    https://github.com/hico-horiuchi/ohgi
    
    Usage:
      ohgi [command]
    
    Available Commands:
      clients     List and delete client(s) information
      jit         Dynamically created clients, added to the client registry
      history     Returns the history for a client
      checks      List locally defined checks and request executions
      request     Issues a check execution request
      events      List and resolve current events
      results     List current check results
      aggregates  List and delete check aggregates
      resolve     Resolves an event
      silence     Create, list, and delete silence stashes
      health      Check the status of the API's transport & Redis connections, and query the transport's status
      info        List the Sensu version and the transport and Redis connection information
      version     Print and check version of ohgi
      help        Help about any command
    
    Flags:
      -x, --datacenter="": Specify a datacenter
      -h, --help[=false]: help for ohgi
    
    Use "ohgi [command] --help" for more information about a command.

#### License

ohgi is released under the [MIT license](https://raw.githubusercontent.com/hico-horiuchi/ohgi/master/LICENSE).
