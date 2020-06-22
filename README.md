# Guardian Admin

[![Build Status](https://drone.monkiato.com/api/badges/monkiato/guardian-admin/status.svg?ref=refs/heads/master)](https://drone.monkiato.com/monkiato/guardian)
[![codecov](https://codecov.io/gh/monkiato/guardian-admin/branch/master/graph/badge.svg)](https://codecov.io/gh/monkiato/guardian)
[![Go Report Card](https://goreportcard.com/badge/github.com/monkiato/guardian)](https://goreportcard.com/report/github.com/monkiato/guardian)

Admin API for Guardian

## Features

 - fetch data required for the admin UI
 
## Endpoints

**GET /users**   get user data users
 
## Build Docker Image

No extra parameters are required for the docker image, so just run:

`docker build . -t guardian-auth`

## Run Docker Container

`docker-compose.yml` is available as a sample to deploy a stack
 with Guardian and a Postgres DB
  
 Run `docker-compose -f docker-compose.yml up -d`

