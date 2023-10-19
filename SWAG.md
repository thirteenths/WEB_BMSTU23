### /auth/singIn

#### POST
##### Summary:

Sing in person.

##### Description:

post person

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| data | body | singInInput | Yes | [handler.singInInput](#handler.singInInput) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | string |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

### /auth/singUp

#### POST
##### Summary:

Create new person.

##### Description:

post person

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| data | body | modelUI.Person | Yes | [modelUI.Person](#modelUI.Person) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | integer |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

### /comics

#### GET
##### Summary:

Show All info about comics.

##### Description:

get comics.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.getAllComicsResponse](#handler.getAllComicsResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) & object |

#### POST
##### Summary:

Create new comic.

##### Description:

post comic

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| data | body | modelUI.Comic | Yes | [modelUI.Comic](#modelUI.Comic) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [handler.statusResponse](#handler.statusResponse) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 403 | Forbidden | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

### /comics/{id}

#### DELETE
##### Summary:

Delete a comic item by ID.

##### Description:

delete comic by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | comic ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.statusResponse](#handler.statusResponse) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 403 | Forbidden | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### GET
##### Summary:

Show info about comic.

##### Description:

get comic by ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | comic ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [modelUI.Comic](#modelUI.Comic) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### PATCH
##### Summary:

Update count of kek comic item by ID.

##### Description:

update comic by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | comic ID | Yes | string |
| data | body | string | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.statusResponse](#handler.statusResponse) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 401 | Unauthorized | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### PUT
##### Summary:

Update info about comic item by ID.

##### Description:

update comic by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | comic ID | Yes | string |
| data | body | modelUI.comic | Yes | [modelUI.Comic](#modelUI.Comic) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.statusResponse](#handler.statusResponse) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 403 | Forbidden | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

### /events

#### GET
##### Summary:

Show All info about events.

##### Description:

get events.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.getAllEventsResponse](#handler.getAllEventsResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### POST
##### Summary:

Create new event.

##### Description:

post event

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| data | body | modelUI.event | Yes | [modelUI.Event](#modelUI.Event) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | integer |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 403 | Forbidden | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

### /events/{id}

#### DELETE
##### Summary:

Delete a event item by ID.

##### Description:

delete event by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | event ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.statusResponse](#handler.statusResponse) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 403 | Forbidden | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### GET
##### Summary:

Show info about event.

##### Description:

get event by ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | event ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [modelUI.Event](#modelUI.Event) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### PUT
##### Summary:

Update info about event item by ID.

##### Description:

update poster by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | event ID | Yes | string |
| data | body | modelUI.event | Yes | [modelUI.Event](#modelUI.Event) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.statusResponse](#handler.statusResponse) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 403 | Forbidden | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

### /places

#### GET
##### Summary:

Show All info about places.

##### Description:

get places.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.getAllEventsResponse](#handler.getAllEventsResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### POST
##### Summary:

Create new place.

##### Description:

post place

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| data | body | modelUI.place | Yes | [modelUI.Place](#modelUI.Place) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [handler.statusResponse](#handler.statusResponse) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 403 | Forbidden | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

### /places/{id}

#### DELETE
##### Summary:

Delete a place item by ID.

##### Description:

delete place by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | place ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.statusResponse](#handler.statusResponse) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 403 | Forbidden | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### GET
##### Summary:

Show info about place.

##### Description:

get place by ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | place ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [modelUI.Place](#modelUI.Place) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### PUT
##### Summary:

Update info about place item by ID.

##### Description:

update place by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | place ID | Yes | string |
| data | body | modelUI.place | Yes | [modelUI.Place](#modelUI.Place) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.statusResponse](#handler.statusResponse) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 403 | Forbidden | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

### /posters

#### GET
##### Summary:

Show All info about posters.

##### Description:

get posters.

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.getAllPostersResponse](#handler.getAllPostersResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### POST
##### Summary:

Create new poster.

##### Description:

post poster

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| data | body | modelUI.Poster | Yes | [modelUI.Poster](#modelUI.Poster) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | integer |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 403 | Forbidden | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

### /posters/{id}

#### DELETE
##### Summary:

Delete a poster item by ID.

##### Description:

delete poster by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | poster ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.statusResponse](#handler.statusResponse) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 403 | Forbidden | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### GET
##### Summary:

Show info about poster.

##### Description:

get poster by ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | poster ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [modelUI.Poster](#modelUI.Poster) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### POST
##### Summary:

CheckIn poster item by ID.

##### Description:

check in comic by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | poster ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.statusResponse](#handler.statusResponse) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 401 | Unauthorized | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

#### PUT
##### Summary:

Update info about poster item by ID.

##### Description:

update poster by ID

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | poster ID | Yes | string |
| data | body | modelUI.poster | Yes | [modelUI.Poster](#modelUI.Poster) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [handler.statusResponse](#handler.statusResponse) |
| 400 | Bad Request | [handler.errorResponse](#handler.errorResponse) |
| 500 | Internal Server Error | [handler.errorResponse](#handler.errorResponse) |

### Models


#### handler.errorResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |

#### handler.getAllComicsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| comics | [ [modelUI.Comic](#modelUI.Comic) ] |  | No |

#### handler.getAllEventsResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| events | [ [modelUI.Event](#modelUI.Event) ] |  | No |

#### handler.getAllPostersResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| posters | [ [modelUI.Poster](#modelUI.Poster) ] |  | No |

#### handler.singInInput

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| login | string |  | Yes |
| password | string |  | Yes |

#### handler.statusResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| status | string |  | No |

#### modelUI.Comic

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| _ | integer |  | No |
| city | string |  | No |
| countOfKek | integer |  | No |
| image | string |  | No |
| name | string |  | No |
| sentence | string |  | No |

#### modelUI.Event

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| modelUI.Event | object |  |  |

#### modelUI.Person

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| _ | integer |  | No |
| email | string |  | No |
| hash | string |  | No |
| login | string |  | No |
| role | integer |  | No |
| salt | string |  | No |

#### modelUI.Place

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| modelUI.Place | object |  |  |

#### modelUI.Poster

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| _ | integer |  | No |
| about | string |  | No |
| date | string |  | No |
| image | string |  | No |
| name | string |  | No |
| place | string |  | No |