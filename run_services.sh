#!/bin/bash

#Internal Service
docker build -t internalservice -f Dockerfile_internalapi.yml
docker run -p 8080:8080 internalservice

#Device Service
docker build -t deviceapi -f Dockerfil_deviceapi.yml
docker run -p 8081:8081 deviceapi
