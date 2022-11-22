
---------------------------------------<b>Event Management</b>-------------------------------------

#### Using Tech:
<b> Programming Language:</b> Golang </br>
<b> Framework:</b> Gin </br>
<b>DB: </b> MySQL </br>
<b>ORM: </b> GORM </br>

####  DB table schema:
<b> DB name: </b> `event_management`
<br>
<b> Tables: </b>

```
events (id, title, start_at, end_at)

workshops (id, event_id, start_at, end_at, title, description)

reservations (id, name, email, workshop_id)

```
#### REST API LIST:

<b>Notes: </b> Project runs on localhost server and 8081 port.

1. <b> Event List API, where we can get all active events with pagination: </b>

Request:
localhost:8081/event-list/:limit/:offset

Example:
localhost:8081/event-list/2/1

Response:
```
{
    "events": [
        {
            "id": 1,
            "title": "event1",
            "start_at": "2022-12-19 03:14:07",
            "end_at": "2022-12-31 03:14:07"
        },
        {
            "id": 2,
            "title": "event2",
            "start_at": "2022-12-19 03:14:07",
            "end_at": "2022-12-31 03:14:07"
        }
    ],
    "pagination": {
        "total": 5,
        "per_page": 2,
        "total_pages": 3,
        "current_page": 1
    }
}
```
2. <b> Event Details API, where we can get single event information: </b>

Request:
localhost:8081/get-event-details/:eventId

Example:
localhost:8081/get-event-details/2

Response:
```
{
    "id": 2,
    "title": "event2",
    "start_at": "2022-12-19 03:14:07",
    "end_at": "2022-12-31 03:14:07",
    "total_workshops": 2
}
```


3. <b> Workshop List API, where we can get all the active workshops of a single event: </b>

Request:
localhost:8081/get-workshop-details-by-event/:eventId

Example:
localhost:8081/get-workshop-details-by-event/1

Response:
```
{
    "id": 1,
    "title": "event1",
    "start_at": "2022-12-19 03:14:07",
    "end_at": "2022-12-31 03:14:07",
    "workshops": [
        {
            "id": 1,
            "title": "workshop1",
            "description": "this is event1 workshop",
            "start_at": "2022-12-19 03:14:07",
            "end_at": "2022-12-31 03:14:07"
        }
    ]
}
```
4. <b> Workshop Details API, where we can get single workshop information: </b>

Request:
localhost:8081/get-workshop-details/:workshopId

Example:
localhost:8081/get-workshop-details/1

Response:
```
{
    "id": 1,
    "title": "workshop1",
    "description": "this is event1 workshop",
    "start_at": "2022-12-19 03:14:07",
    "end_at": "2022-12-31 03:14:07",
    "total_reservations": 3
}
```
5.<b> Make a workshop reservation API </b> </br>
If user make a workshop reservation then they need to post with below post body
</br>
Request: (Post Api)

localhost:8081/workshop-reservationV2

Post Body:
```
{
    "name":"mohaiminur",
    "email":"mohaiminur404@gmail.com",
    "workshop_id":"1"
}
```
Response:
```
{
    "reservation": {
        "id": 1,
        "name": "mohaiminur",
        "email": "mohaiminur404@gmail.com"
    },
    "event": {
        "id": 1,
        "title": "event1",
        "start_at": "2022-12-19 03:14:07",
        "end_at": "2022-12-31 03:14:07"
    },
    "workshop": {
        "id": 1,
        "title": "workshop1",
        "description": "this is event1 workshop",
        "start_at": "2022-12-19 03:14:07",
        "end_at": "2022-12-31 03:14:07"
    }
}
```
if user already reserved workshop then if they only give existing username and email then they can also get workshop reservation
```
{
"name":"mohaiminur",
"email":"mohaiminur404@gmail.com"
}
```
Response:
```
{
    "reservation": {
        "id": 1,
        "name": "mohaiminur",
        "email": "mohaiminur404@gmail.com"
    },
    "event": {
        "id": 1,
        "title": "event1",
        "start_at": "2022-12-19 03:14:07",
        "end_at": "2022-12-31 03:14:07"
    },
    "workshop": {
        "id": 1,
        "title": "workshop1",
        "description": "this is event1 workshop",
        "start_at": "2022-12-19 03:14:07",
        "end_at": "2022-12-31 03:14:07"
    }
}
```

---------------------------------------------Thanks------------------------------------------------