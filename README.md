# Experimental foundation for an API server

## Introduction

This project is an experiment to:
a) learn Golang.
b) create the core components of an API server, to be used in several projects I have planned.
c) determine the right choice of technologies for a high performance, reliable service.

## Goals (TBD)

- ~~Add oauth2 support for user authentication~~ JWT is more appropriate for API authentication.
- Use external file (yaml or json) for configuration

# TODO

The current focus of work:
- Get authentication system in place
    - Setup database for users
    - Validate /authenticate against user database
- Figure out how to handle routing better, so heimdall can be included in another project and just provide the routing/etc.
## Resources

Tutorials and articles that I've used/read to build this:
- https://golang.org/doc/code.html
- https://thenewstack.io/make-a-restful-json-api-go/
- https://www.thepolyglotdeveloper.com/2017/03/authenticate-a-golang-api-with-json-web-tokens/
