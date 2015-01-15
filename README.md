Negroni Boilerplate (MySQL version)
===============

This is a forked Negroni web app with basic auth and user registrations via MySQL database

Note
----------
We're working on Bcrypt passwords...

Features
----------
* Signup/Login
* Pages rendered from templates
* Very Simple API call via Javascript on homepage

Requirements
-----------

* go get github.com/codegangsta/negroni
* go get github.com/goincremental/negroni-sessions
* go get github.com/goincremental/negroni-sessions/cookiestore
* go get github.com/go-sql-driver/mysql
* go get github.com/unrolled/render

Configuration
--------------
Import users-mysql.sql in your MySQL database, then edit database url in main.go .

Soon I'll provide a config file for settings.

To Do
-----------

* Fix login with Bcrypt
* More extensive api examples
* Add a minimal CRUD


