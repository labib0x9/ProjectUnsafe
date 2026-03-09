# ProjectUnsafe

A simple playground for Linux Playground..

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