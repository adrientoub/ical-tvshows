# README

This project is about creating a server generating icalendar files containing
a personalized TV Shows schedule.

## Installation

It is really easy to install this project on you own server.
- Install `go` on your server (tested with go 1.6.2)
- Configure the `GOPATH` environment variable
- `go get github.com/adrientoub/ical-tvshows`
- Create the configuration file
- Go to the `$GOPATH/bin/` folder and launch `./ical-tvshows`

You also need to install Redis if you want to use it as your caching system.

## Configuration

Create a `config.json` file in the folder containing the program. It must
contain the following informations:

```
{
  "uri": "LISTEN_ADDRESS:PORT",
  "api_key": "A Betaseries API_KEY",
  "cache": "redis",
  "redis": {
    "address": "localhost:6379",
    "password": "",
    "db": 0
  }
}
```

In the config, the "cache" key can be any of: "redis", "files" or "none".
If you choose to use Redis as a cache you must set the connection parameters.
