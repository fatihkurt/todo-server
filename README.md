# todo-server

## Installation

Execute command ``go get .`` or ``go mod install`` for install dependencies and ``go run .`` to start api server.

Program needs to run mysql server on localhost with ``root:password@/todo`` dsn string.

To run localhost a docer compose file can run on local terminal with ``docker compose up``.

Execute command ``go test ./test -v`` to run tests.

---

### TODO
* Fix tests
* Get passwords by env
