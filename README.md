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

    $ sudo -u postgres psql -U postgres
    postgres=# CREATE DATABASE henhouse;
    postgres=# CREATE USER henhouse WITH password 'STRENGTH_PASSWORD';
    postgres=# GRANT ALL privileges ON DATABASE henhouse TO henhouse;
    postgres=# \c henhouse
    postgres=# GRANT CREATE ON SCHEMA public TO henhouse;


After that you need to fix 'connection' parameter in configuration file.
(And other parameters, of course)

Now, run it!

    $ ${GOPATH}/bin/henhouse ${GOPATH}/src/github.com/MikeLupanov/henhouse/config/henhouse.toml --reinit
