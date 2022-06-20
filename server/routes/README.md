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
POST http://dummy-host.com/api/app
```
### Description
Create a new Haddock app

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








## GET /api/user
```http request
GET http://dummy-host.com/api/user?token=xxxxx.yyyyy.zzzzz
```
### Description
returns a user object associated with the token

### Args
**`token`:** a [jwt](https://jwt.io/introduction) token associated with the user




## POST /api/user/auth
```http request
POST http://dummy-host.com/api/user/auth
```

### Description
Authenticate a Haddock user matching with the provided credentials

### Forms values
**`username`:** the username of the user <br/>
**`password`:** the password of the user



## POST /api/app
```http request
POST http://dummy-host.com/api/user?token=xxxxx.yyyyy.zzzzz
```
### Description
Create a new Haddock user

### Args
**`token` (optional):** a [jwt](https://jwt.io/introduction) token associated with an authorised user, is used when the server does not allow account creation by the user itself.

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
DELETE http://dummy-host.com/api/user?user=xxxxx?token=xxxxx.yyyyy.zzzzz
```
### Description
Delete an existing Haddock user

### Args
**`user`:** the user to delete uuid. <br/>
**`token`:** a [jwt](https://jwt.io/introduction) token associated with an authorised user.
