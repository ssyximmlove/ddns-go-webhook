# DDNS-Go-Webhook

## Description

[DDNS-Go-Webhook](https://github.com/ssyximmlove/ddns-go-webhook) is a simple middleware that connect [DDNS-Go](https://github.com/jeessy2/ddns-go) and [Everypush](https://github.com/PeanutMelonSeedBigAlmond/EveryPush), based on the Golang standard library "net/http"

## Usage

```docker compose
version: '3.7'

services:
  ddns-go-webhook:
    image: ssyximmlove/ddns-go-webhook:latest
    ports:
      - 8080:8080
    volumes:
      - "./webhook:/app/config"
```
Write your configuration file in the ./webhook directory, and the configuration file name is config.toml

## Configuration

```toml
[app]
addr = ":8080"
endpoint = "PushEndpoint"
```
