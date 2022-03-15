# routes
here you can find a description of existing the api routes and the expected arguments

## GET /ws
```http request
GET http://dummy-host.com/ws?token=xxxxxx
```
### Description
It's the entrypoint of the websocket
connect to receive events from the api.

###Args
**`token`:** (optional) if accounts are enabled it is needed to authenticate



## GET /api/app
```http request
GET http://dummy-host.com/api/app?app=xxxxxx
```
### Description
List all Haddock's images

### Args
**`app`:** (optional) filters the results (e.g. `?app=busybox`)




## POST /api/app
```http request
POST http://dummy-host.com/api/app?app=xxxxxx
```
### Description
List all Haddock's images

### Args
**`app`:** (optional) filters the results (e.g. `?app=busybox`)
