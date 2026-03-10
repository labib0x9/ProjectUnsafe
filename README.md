# ProjectUnsafe

A simple playground for Linux Playground, each lab is sandboxed using a container.

---

# Features

- Anonymous login, a temporary account for only 30minutes lifecycle.
- Realtime lab communication via websocket and xterm.js.
- Lab starts in container. Reset the container and terminate the lab.

---

# Architecture

    [Client] -> [Frontend] -> [Backend] -> [Client]
                                    |
                                    V
    --------------------------------------------------------
    |                  [BACKEND]                           |
    |                                                      |
    |                                                      |
    |                                                      |
    |                                                      |
    |                                                      |
    |                                                      |
    |                                                      |
    |                                                      |
    --------------------------------------------------------

---

# API LIST
## Lab APIs
```
GET /labs
GET /lab/${id}
POST /lab/start
POST /lab/reset
POST /lab/terminate
GET /lab/hints?labId=${labId}
```

## Auth APIs
```
POST /auth/login
POST /auth/signup
POST /auth/anonymous
POST /auth/reset-password
POST /auth/logout
```

## Playground APIs
```
GET /problems
GET /problem/${id}
POST /code/run-custom
POST /code/run
```

## Admin APIs
```
GET /admin/containers
GET /admin/users
POST /admin/terminate
```