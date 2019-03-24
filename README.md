# urlcat
Command line URL processing tool


## Usage

```
$ urlcat --help
NAME:
   urlcat - url processing tool. Reads line separated

USAGE:
    [global options] command [command options] [arguments...]

VERSION:
   1.0.0

DESCRIPTION:
   Reads Line separated URL's from stdin and returns the requested segment

COMMANDS:
     filename, f   returns the filename (default)
     extension, e  returns the file extension
     host, h       Returns the host
     scheme, s     Returns the URL scheme
     query, q      Returns the query string
     help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```