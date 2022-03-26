<div align="center">

# :pie: Pi-connect

## Connect to your Pi without mDNS

</div>

This module is built for connection to your raspberry pi without using 
mDNS. This is useful in cases where

1. You are setting up your pi in headless mode.
2. The IP address of your raspberry pi changes constantly.
3. Broadcasts do not work reliably on your network for mDNS.
4. Finding your pi's IP is time-consuming (No access to router / large subnet).

## :zap: Quick start

### Install

```terminal
go install github.com/andylolu2/pi-connect
```

### Usage

#### Set environment variables

##### Option 1:

Add `MONGODB_URI=value` to `.env` file.

```bash
echo 'MONGODB_URI=mongodb+srv://<connection_str>' >> .env
```

##### Option 2:

Set `MONGODB_URI` directly.
```bash
export MONGODB_URI='mongodb+srv://<connection_str>'
```

##### Option 3

Pass `MONGODB_URI` to command.

```bash
MONGODB_URI='mongodb+srv://<connection_str>' pi-connect postip
```

#### Help

```bash
pi-connect -h
# Posts and gets the ip address of the hosts to mongodb. Can be used to find the local ip address of raspberry pi in headless setup.

# Usage:
#   pi-connect [flags]
#   pi-connect [command]

# Available Commands:
#   completion  Generate the autocompletion script for the specified shell
#   getip       Get the specified host machine's ip address from a mongodb document.
#   help        Help about any command
#   postip      Post the host machine's ip address to a mongodb document.
#   ssh         Start ssh session to the specified host machine's name.

# Flags:
#   -h, --help   help for pi-connect

# Use "pi-connect [command] --help" for more information about a command.
```
