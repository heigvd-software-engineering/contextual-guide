# Contributing to Contextual Guide

Start by installing [Go](https://golang.org/doc/install), [Docker](https://docs.docker.com/get-docker/), and [Docker Compose](https://docs.docker.com/compose/install/).

Then, clone the git repository.

```
git clone git@github.com:heigvd-software-engineering/contextual-guide.git
```

Copy the .env.example file to .env and modify it according to your setup. 
You will need a mail account and an SMTP server to enable the [GoTrue](https://github.com/netlify/gotrue) authentication server.

```
cp .app.env.example .app.env
vim .app.env
cp .gotrue.env.example .gotrue.env
vim .gotrue.env
```

You can now start the services with Docker Compose.

```
docker-compose up
```
