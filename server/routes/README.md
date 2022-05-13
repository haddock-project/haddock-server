# routes
here you can find a description of existing the api routes and the expected arguments

# /API node

## GET /api/ws
```http request
GET http://dummy-host.com/api/ws?token=xxxxxx
```
### Description
It's the entrypoint of the websocket
connect to receive events from the api.

###Args
**`token`:** auth token (not implemented yet)



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
Download a new Haddock image

### Args
**`app`:** the dockerhub ID of the app you want to download (e.g. `?app=busybox`)

### Events
**`APP_DOWNLOAD_ERROR`:** tells that an error has occurred while downloading the app<br/>
- `image_name`: the concerned image

<br/>

**`APP_DOWNLOAD_COMPLETE`:** tells that an app have been successfully downloaded <br/>
- `image_name`: the concerned image

<br/>

**`APP_DOWNLOAD_PROGRESS`:** 
- `image_name`: the concerned image
- `progress`:
  - `current`: the number of bytes downloaded
  - `total`: the total number of bytes to download

<br/>

**`APP_EXTRACT_PROGRESS`:** <br/>
- `image_name`: the concerned image
- `progress`:
    - `current`: the number of bytes downloaded
    - `total`: the total number of bytes to download

## DELETE /api/app
```http request
DELETE http://dummy-host.com/api/app?app=xxxxxx
```
### Description
remove a Haddock image

### Args
**`app`:** the dockerhub ID of the app you want to remove (e.g. `?app=busybox`)
