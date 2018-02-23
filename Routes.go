package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "UserIndex",
        "GET",
        "/users",
        UserIndex,
    },
    Route{
        "UserShow",
        "GET",
        "/users/{usreid}",
        UserShow,
    },
    Route{
        "Version",
        "GET",
        "/version",
        Version,
    },
    Route{
        "Status",
        "GET",
        "/status",
        NotImplemented,
    },

    Route{
        "Authenticate",
        "POST",
        "/authenticate",
        CreateTokenEndpoint,
    },
    Route{
        "Test",
        "GET",
        "/test",
        Authenticate(TestEndpoint),
    },

    // router.HandleFunc("/authenticate", CreateTokenEndpoint).Methods("POST")
    // router.HandleFunc("/protected", ProtectedEndpoint).Methods("GET")
    // router.HandleFunc("/test", ValidateMiddleware(TestEndpoint)).Methods("GET")
}
