## ohgi v0.1.2

![image.png](https://raw.githubusercontent.com/hico-horiuchi/ohgi/master/image.png)

#### Requirements

  - [Golang](https://golang.org/) >= 1
  - [Sensu](http://sensuapp.org/) >= 0.13
  - [spf13/cobra](https://github.com/spf13/cobra)

#### Installation

    $ git clone git://github.com/hico-horiuchi/ohgi.git
    $ cd ohgi
    $ go get ./...
    $ sudo make install

#### Configuration

`~/.sensu.json`

    {
      "host": "127.0.0.1",  // Required
      "port": 4567,         // Required
      "user": "",           // Optional
      "password": ""        // Optional
    }

#### Usage

    Sensu command-line tool by golang
    https://github.com/hico-horiuchi/ohgi
    
    Usage:
      ohgi [command]
    
    Available Commands:
      checks [check]               Returns the list of checks
      request [check] [subscriber] Issues a check execution request
      clients [client]             Returns the list of clients
      history [client]             Returns the client history
      events [client] [check]      List and resolve current events
      resolve [client] [check]     Resolves an event (delayed action)
      health                       Returns the API info
      info                         Returns the API info
      silence [client] [check]     Returns a list of silences
      version                      Print git revision of ohgi
      help [command]               Help about any command
    
    Use "ohgi help [command]" for more information about that command.

#### License

Ohgi is released under the [MIT license](https://raw.githubusercontent.com/hico-horiuchi/ohgi/master/LICENSE).
