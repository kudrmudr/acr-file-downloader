# Acronis Downloader

## Introduction


the structure is:

`client` and `server` - code for client and server

`fixtures` - folder for uploaded files (for server)

`downloads` - folder for downloaded files

## Requirements

* Docker CE 17.06 or later
* docker-composer 1.24.1 or later

## Provisioning (dev)

First time run:

```
$ make start
```

Run the client (Downloader)

```
make download
```

Bring the containers up/down:

```
$ make start

$ make stop
```

Run tests:
```
$ make test
```

## TODO

- Imporve error handling and logging
- add more unit tests


## Assumptions 

 - Symbol should be 1 byte size