[![Build Status](https://travis-ci.org/MikeLupanov/henhouse.svg?branch=master)](https://travis-ci.org/MikeLupanov/henhouse)
[![Deb Package](https://img.shields.io/badge/deb-packagecloud.io-844fec.svg)](https://packagecloud.io/MikeLupanov/henhouse)
[![GoDoc](https://godoc.org/github.com/MikeLupanov/henhouse?status.svg)](http://godoc.org/github.com/MikeLupanov/henhouse)
[![Coverage Status](https://coveralls.io/repos/MikeLupanov/henhouse/badge.svg?branch=master&service=github)](https://coveralls.io/github/MikeLupanov/henhouse?branch=master)
[![Go Report Card](http://goreportcard.com/badge/MikeLupanov/henhouse)](http://goreportcard.com/report/MikeLupanov/henhouse)

# Henhouse

Scoreboard for jeopardy-style CTFs.

Fundamental principle: if henhouse is not helping you make jeopardy-style CTF easily, then there is a bug in henhouse.

![Imgur image](https://i.imgur.com/uMCFPw7.png)

## Install

### Packagecloud

Built for Ubuntu 16.04.

    $ curl -s https://packagecloud.io/install/repositories/MikeLupanov/henhouse/script.deb.sh | sudo bash
    $ sudo apt install henhouse

### Build deb package from source

    $ apt install golang build-essential binutils upx-ucl
    $ export GOPATH=$(realpath ./) && go get github.com/MikeLupanov/henhouse/...
    $ cd ${GOPATH}/src/github.com/MikeLupanov/henhouse
    $ ./package.sh

## Development

### Depends

#### Gentoo

    $ sudo emerge dev-lang/go dev-db/postgresql

#### Ubuntu

    $ sudo apt install golang postgresql

### Build

First you need set GOPATH environment variable.

    $ export GOPATH=$(realpath ./)

After you need download and build henhouse with depends.

    $ go get github.com/MikeLupanov/henhouse

### Run

    $ sudo psql -U postgres
    postgres=# CREATE DATABASE henhouse;
    postgres=# CREATE USER henhouse WITH password 'STRENGTH_PASSWORD';
    postgres=# GRANT ALL privileges ON DATABASE henhouse TO henhouse;

After that you need to fix 'connection' parameter in configuration file.
(And other parameters, of course)

Now, run it!

    $ ${GOPATH}/bin/henhouse ${GOPATH}/src/github.com/MikeLupanov/henhouse/config/henhouse.toml --reinit
