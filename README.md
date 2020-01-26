# King's Thursday Bootcamp on GoLang

The subject matter of this bootcamp is the Go Programming Language. We will look at some basic principles of Go and some popular frameworks that are frequently use to build REST Api's, and then we will build from scratch one simple Web REST Api for an todo application.

## Agenda

- Theoretical overview of GoLang (1 hour or so)

- Coding session (rest of the time, up to two hours, or longer if desired)

## Prerequisites

For the coding session, depending on your favorite operating system you'll need to setup your environment and workspace(install Go on your local machine, setup necesary encironment variables, etc) before the bootcamp.We don't want to spend time during the workshop installing Go.

Some installation guides for your speific operation system:

* Windows: https://www.freecodecamp.org/news/setting-up-go-programming-language-on-windows-f02c8c14e2f/
* Macbook: https://tecadmin.net/install-go-on-macos/
* Linux: https://www.linode.com/docs/development/go/install-go-on-ubuntu/

## Coding Session

Develop an REST Api which will be used to manage an todo applciation , every todo has some basic data that you'll 
have to manage like: name, description project, date, time, priority. No database is needed to store the data since 
we're not gonna work with ORM in Go for this session. 

Write a very simple client that calls your REST API and retrieves the information in json format.

You have some simple data in the `examples/samples` directory.

### Todo REST Api Endpoints

#### Create a new todo:  `POST /todos`

Create a new todo by posting its details to `/todos`. You receive back its details
in `JSON` format.

```
POST /todos
Content-Type: application/json
{   
    "name": "Attend Artificial Inteligence Laboratory", 
    "description": "Last attendence for in year", 
    "project" : "University",
    "date":  "1/9/2019"
    "time": "13:00 PM"
    "priority": 3
}
```

```
HTTP/1.0 200 OK
Content-Type: application/json
{
    "name": "Attend Artificial Inteligence Laboratory",
    "description": "",
    "project": "University",
    "id": "693cc8a2-75f0-4c09-9b42-c5a3f6a2c9ab",
    "date": "1/9/2019",
    "time": "13:00 PM",
    "priority": "3"
},
```

#### Retrieve all todos: `GET /todos`

You will receive back a list with all existing todos.

```
GET /todos
```

```
HTTP/1.0 200 OK
Content-Type: application/json
{
    "todos" : [
        {
            "name": "Attend Artificial Inteligence Laboratory",
            "description": "",
            "project": "University",
            "id": "693cc8a2-75f0-4c09-9b42-c5a3f6a2c9ab",
            "date": "1/9/2019",
            "time": "13:00 PM",
            "priority": "3"
        },
        {
            "name": "Send Reports to Project Manager",
            "description": "",
            "project": "Work",
            "id": "eef9c28f-ebe5-4a39-8fa7-78c7d917a1b2",
            "date": "1/2/2019",
            "time": "10:00 AM",
            "priority": "4"
        },
    ]
}
```

#### Update a Todo: `POST /todo/<uuid>`

Update your todo with new data:

```
POST /todo/eef9c28f-ebe5-4a39-8fa7-78c7d917a1b2
{
    "name": "Send Reports to Project Manager",
    "description": "Monthly billing",
    "project": "Work",
    "id": "eef9c28f-ebe5-4a39-8fa7-78c7d917a1b2",
    "date": "1/3/2019",
    "time": "15:00 PM",
    "priority": "4"
}
```

```
HTTP/1.0 200 OK
```

#### Retrieve a Todos by Project: `GET /todo/<project>`

```
GET /todo/work
```

```
HTTP/1.0 200 OK
Content-Type: application/json

[
    {
        "name": "Send Reports to Project Manager",
        "description": "",
        "project": "Work",
        "id": "eef9c28f-ebe5-4a39-8fa7-78c7d917a1b2",
        "date": "1/2/2019",
        "time": "10:00 AM",
        "priority": "4"
    },
    {
        "name": "Commit Changes from last workload",
        "description": "",
        "project": "Work",
        "id": "58c82be6-7469-4a34-a8c3-8d0a923e9f46",
        "date": "1/10/2019",
        "time": "12:00 AM",
        "priority": "4"
    }
]
```