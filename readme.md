# Bookings and Reservations

The repository for [Building Modern Web Applications with Go](https://www.udemy.com/course/building-modern-web-applications-with-go/?referralCode=0415FB906223F10C6800).

Start Server

```
go run cmd/web/*.go
```

- Built in Go version 1.15
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards scs session management](github.com/alexedwards/scs)
- Uses [nosurf](github.com/justinas/nosurf)

Session

- Put

```
remoteIp := r.RemoteAddr
m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

```

- Get

```
remoteIP = m.App.Session.GetString(r.Context(), "remote_ip")
```

Add Execute Permission
chmod +x run.sh

so that we can run directly as follows

./run.sh

DB Migration using soda

https://gobuffalo.io/documentation/database/configuration/

Best practice to only include database.yml.example in git (ignore database.yml)

Create Migration Files

```
soda generate fizz <migration_name>

soda generate fizz CreateUserTable
```

Migrations

```
soda migrate up

soda migrate down
```
