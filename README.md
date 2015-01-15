Negroni Boilerplate (MySQL version)
===============

This is a forked Negroni web app with basic auth and user registrations via MySQL database

Note
----------
We're working on Bcrypt passwords...

Features
----------
* Signup/Login
* TOML support for configuration
* Pages rendered from templates
* Very Simple API call via Javascript on homepage

Requirements
-----------
* go get github.com/BurntSushi/toml
* go get github.com/codegangsta/negroni
* go get github.com/goincremental/negroni-sessions
* go get github.com/goincremental/negroni-sessions/cookiestore
* go get github.com/go-sql-driver/mysql
* go get github.com/unrolled/render

Configuration
--------------
Import users-mysql.sql in your MySQL database, then edit config.toml

To Do
-----------

* Fix login with Bcrypt
* More extensive api examples
* Add a minimal CRUD


