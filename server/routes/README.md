# routes
here you can find a description of existing the api routes and the expected arguments

# /API node

## GET /api/ws
```http request
GET http://dummy-host.com/api/ws
```
### Description
It's the entrypoint of the websocket
connect to receive events from the api.

### Headers
**`Authorization`:** Bearer [auth token]







## GET /api/app
```http request
GET http://dummy-host.com/api/app?app=xxxxxx
```
### Description
List all Haddock's images/get a specific app

### Args
**`app`:** (optional) get a specific app (e.g. `?app=busybox`)




## POST /api/app
```http request
POST http://dummy-host.com/api/app
```
### Description
Create a new Haddock app

### Headers
**`Authorization`:** Bearer [auth token]

### Body
A [JSON representation](../../api/database/apps.go) of the app, if a UUID is provided it will be overwritten.

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




## PATCH /api/app
```http request
PATCH http://dummy-host.com/api/app
```
### Description
Edit an existing Haddock app

### Body
A [JSON representation](../../api/database/apps.go) of the app, if an empty field is provided, the existing field will be removed.

### Headers
**`Authorization`:** Bearer [auth token]

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

### Headers
**`Authorization`:** Bearer [auth token]

### Args
**`app`:** the dockerhub ID of the app you want to remove (e.g. `?app=busybox`)








## GET /api/user
```http request
GET http://dummy-host.com/api/user
```
### Description
returns a user object associated with the token

### Headers
**`Authorization`:** Bearer [auth token]




## POST /api/user/auth
```http request
POST http://dummy-host.com/api/user/auth
```

### Description
returns a Haddock user token matching with the provided credentials

### Forms values
**`username`:** the username of the user <br/>
**`password`:** the password of the user <br/>
**`remember_me`:** (optional) if set to true, the token will be valid for the amount of time specified in the configuration file



## POST /api/user
```http request
POST http://dummy-host.com/api/user
```
### Description
Create a new Haddock user

require an authenticated user when the server does not allow account creation by the user itself.

### Headers
**`Authorization` (optional):** Bearer [auth token] 

### Body
A [JSON representation](../../api/database/users.go) of the new user, if a UUID is provided it will be overwritten.



## PATCH /api/user
```http request
PATCH http://dummy-host.com/api/user?token=xxxxx.yyyyy.zzzzz
```
### Description
Edit an existing Haddock user

### Body
A [JSON representation](../../api/database/users.go) of the user, if an empty field is provided, the existing field will be removed.




## DELETE /api/user
```http request
DELETE http://dummy-host.com/api/user
```
### Description
Delete an existing Haddock user

### Headers
**`Authorization`:** Bearer [auth token]

### Args
**`user`:** the user to delete uuid. <br/>
