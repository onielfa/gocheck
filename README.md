# GoCheck

Golang application created to make easier to know if a service is up or ko.

## Idea

The idea is to create a program that will be responsible to check whether a server is up or not. The basic functioning of the application is to ask the server every x units of time and produce an output.

This could be useful to containers which don't have and health endpoint.

## HTTP Checks

Using Barrier pattern, exported functions can be used to know if some endpoints are up and which return code we are receiving.
See Tests for examples.
