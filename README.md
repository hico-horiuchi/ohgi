## ohgi

![image.png](https://raw.githubusercontent.com/hico-horiuchi/ohgi/master/image.png)

#### Installation

    $ git clone git://github.com/hico-horiuchi/ohgi.git
    $ cd ohgi
    $ go get ./...
    $ sudo make install

#### Configuration

    /* ~/.sensu.json */
    
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
    events                    List and resolve current events
    version                   Print git revision of ohgi
    help [command]            Help about any command
    
    Use "ohgi help [command]" for more information about that command.
