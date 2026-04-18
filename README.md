# ProjectPDF

ProjectPDF is a API service that accepts images and converts them into a combined pdf.

---

# Features

- User can authenticate.
- Authenticated user can save upto 10pdfs.
- User upload images (each image upto 5mb, at most 10 images), then generate a combined pdf.
- Download the generated pdf

---

# Architecture

    [Client] -> [Frontend] -> [Backend] -> [Client]
                                  |
                                  V
    ----------------------------------------------------------------------------------------------------
    |                                             [BACKEND]                                            |
    |                                                                                                  |
    |               [ROutes]   ->    [           Middlewares        ]                                  |
    |                               /                 |             \                                  |
    |                              V                  V              V                                 |
    |                        [Auth Service]  [Convert Service]  [User Service]                         |
    |                                                |                                                 |
    |                                                V                                                 |
    |                                           [Job Queue]                                            |
    |                                                |                                                 |
    |                                                V                                                 |
    |                   [Object storage]    <-   [Worker pool] (pdf convertion)    ->   [Database]     |
    |                                                                                                  |
    ----------------------------------------------------------------------------------------------------

- Each convertion is executed in isolated environment, worker pool selects container
- PDFs are stored in object storage
- Graceful shutdown
- Image verification(corrupt and malicious files, wrong MIME)

---

# API LIST
## User APIs
```
GET /users/profile
POST /users/profile
POST /users/change-password
```

## PDFs APIs
```
GET /pdfs
GET /pdfs/{id}/download
DELETE /pdfs/{id}
```

## Auth APIs
```
POST /auth/login
POST /auth/logout
POST /auth/signup
POST /auth/reset-password
GET /auth/verify/{token}
```

## Convert APIs
```
POST /convert
GET /convert/status/{jobId}
```

## Admin APIs
```
GET /admin/users
DELETE /admin/users/{id}
POST /container/{id}/down
POST /container/{id}/up
GET /jobs
POST /jobs/{id}/status
POST /jobs/{id}/down
```

```
cp .env.example .env
```