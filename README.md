# Guardian Admin

[![Build Status](https://drone.monkiato.com/api/badges/monkiato/guardian-admin/status.svg?ref=refs/heads/master)](https://drone.monkiato.com/monkiato/guardian-admin)
[![codecov](https://codecov.io/gh/monkiato/guardian-admin/branch/master/graph/badge.svg)](https://codecov.io/gh/monkiato/guardian-admin)
[![Go Report Card](https://goreportcard.com/badge/github.com/monkiato/guardian-admin)](https://goreportcard.com/report/github.com/monkiato/guardian-admin)

Admin API for Guardian.

It's used to fetch data required for the Guardian Admin UI. The reason why this is a different service and not part of the main Guardian API is because this is information not exposed publicly, the Admin API is an internal bridge between the Admin UI and the DB used to obtain sensitive information.
 
## Endpoints

**GET /users**   get user data users
 
## Build Docker Image

No extra parameters are required for the docker image, so just run:

`docker build . -t guardian-auth`

## Run Docker Container

`docker-compose.yml` is available as a sample to deploy a stack with Guardian and a Postgres DB

Run `docker-compose -f docker-compose.yml up -d`
