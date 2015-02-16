## ohgi

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
      checks [check]            Returns the list of checks
      clients [client]          Returns the list of clients
      events [client] [check]   List and resolve current events
      resolve [client] [check]  Resolves an event (delayed action)
      health                    Returns the API info
      info                      Returns the API info
      version                   Print git revision of ohgi
      help [command]            Help about any command
    
    Use "ohgi help [command]" for more information about that command.
