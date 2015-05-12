## ohgi v0.1.7

![image.png](https://raw.githubusercontent.com/hico-horiuchi/ohgi/master/image.png)

#### Requirements

  - [Golang](https://golang.org/) >= 1
  - [Sensu](http://sensuapp.org/) >= 0.13
  - [spf13/cobra](https://github.com/spf13/cobra)
  - [tcnksm/go-latest](https://github.com/tcnksm/go-latest)

#### Installation

    $ git clone git://github.com/hico-horiuchi/ohgi.git
    $ cd ohgi
    $ go get ./...
    $ sudo make install

#### Configuration

`~/.ohgi.json`

    {
      "host": "127.0.0.1",  // Required
      "port": 4567,         // Required
      "user": "",           // Optional
      "password": "",       // Optional
      "timeout": 3          // Optional
    }

#### Usage

    Sensu command-line tool by golang
    https://github.com/hico-horiuchi/ohgi
    
    Usage:
      ohgi [command]
    
    Available Commands:
      checks      List locally defined checks and request executions
      request     Issues a check execution request
      clients     List and delete client(s) information
      jit         Dynamically created clients, added to the client registry
      history     Returns the history for a client
      events      List and resolve current events
      resolve     Resolves an event
      health      Check the status of the API's transport & Redis connections, and query the transport's status
      info        List the Sensu version and the transport and Redis connection information
      silence     Create, list, and delete silences
      version     Print and check version of ohgi
      help        Help about any command
    
    Flags:
      -h, --help=false: help for ohgi
    
    Use "ohgi help [command]" for more information about a command.

#### License

ohgi is released under the [MIT license](https://raw.githubusercontent.com/hico-horiuchi/ohgi/master/LICENSE).
