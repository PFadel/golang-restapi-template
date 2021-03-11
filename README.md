# Golang REST API Template

[![Go Report Card](https://goreportcard.com/badge/github.com/PFadel/golang-restapi-template)](https://goreportcard.com/report/github.com/PFadel/golang-restapi-template)
[![Build Status](https://travis-ci.com/PFadel/golang-restapi-template.svg?branch=main)](https://travis-ci.com/PFadel/golang-restapi-template)
[![codecov](https://codecov.io/gh/PFadel/golang-restapi-template/branch/main/graph/badge.svg?token=0NX7BIMV7U)](https://codecov.io/gh/PFadel/golang-restapi-template)

This is my try at creating a golang rest api template for anyone to use. It is entirely based on my own techinal opinion and personal taste, but I will try to explain why I did some things the way I did. I also am open to receive suggestions and (constructive) criticism.

## No-Brainers

Things I believe are no-brainers because most of the community use and are considered good practices so far:

- [docker](https://www.docker.com/) & [docker-compose](https://docs.docker.com/compose/install/)
- [go mod](https://golang.org/ref/mod)
- [golangci](https://github.com/golangci/golangci-lint)
- [Logrus](https://github.com/sirupsen/logrus)

That being said, I will try to speak about my other choices below.

## Other Choices

### stdlib net/http

### uber-go/fx
