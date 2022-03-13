#!/bin/bash

docker run -v ~/haddock/server/persistent/:/data/ -v /var/run/docker.sock:/var/run/docker.sock haddock-server