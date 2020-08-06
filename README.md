# kanka 

[![PkgGoDev](https://pkg.go.dev/badge/Henry-Sarabia/kanka)](https://pkg.go.dev/Henry-Sarabia/kanka) [![Build Status](https://travis-ci.com/Henry-Sarabia/kanka.svg?branch=master)](https://travis-ci.com/Henry-Sarabia/kanka) [![Coverage Status](https://coveralls.io/repos/github/Henry-Sarabia/kanka/badge.svg?branch=master)](https://coveralls.io/github/Henry-Sarabia/kanka?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/Henry-Sarabia/kanka)](https://goreportcard.com/report/github.com/Henry-Sarabia/kanka)

Manage your [Kanka](https://kanka.io/en-US) campaign or build tools for other
creators with the thoroughly tested and documented **kanka** package.

The **kanka** package provides a client to handle all communication with the
[Kanka API](https://kanka.io/en-US/docs/1.0). 

The package is structured into convenient and discoverable services for 
managing characters, locations, organizations, events, and 
[much more](https://kanka.io/en-US/features).


## Installation

If you do not have Go installed yet, you can find installation instructions 
[here](https://golang.org/doc/install).

To pull the most recent version of **kanka**, use `go get`.

```
go get github.com/Henry-Sarabia/kanka
```

Then import the package into your project as you normally would.

```go
import "github.com/Henry-Sarabia/kanka"
```

## Usage

### Creating A Client

To use the **kanka** package, you need a Kanka API key. If you do not have a key
yet, you can follow the instructions [here](https://kanka.io/en-US/docs/1.0/setup)
to get one.

Create a client with your API key to start communicating with the Kanka API.

```go
c, err := kanka.NewClient("YOUR_API_KEY", nil)
```

If you need to use a preconfigured HTTP client, simply pass its address to the
`NewClient` function.

```go
c, err := kanka.NewClient("YOUR_API_KEY", &custom)
```

### Services

The client contains a separate service for working with each of the Kanka API
endpoints. Each service has a set of functions to retrieve, list, create, update,
or delete campaign data.

To start communicating with the Kanka API, choose a service and call one of its
functions. 

Take the `Campaigns` service for this example. 

To retrieve a list of the current user's campaigns, use the `Index` function.

```go
cmps, err := c.Campaigns.Index()
```
You now have access to a list of the user's campaigns via `cmps`.

### Retrieving An Entity

To retrieve a specific entity from a campaign, use the `Get` function.

Take the `Quests` service for example. 

For this service, `Get` requires a campaign ID and quest ID.

```go
qst, err := c.Quests.Get(cmpID, qstID)
```

The result is stored in `qst` of type `Quest`. 

### Retrieving A List Of Entities

To retrieve a list of a campaign's entities of a certain type, use the `Index` function.

Take `Locations` for example.

For this service, `Index` requires only a campaign ID.

```go
locs, err := c.Locations.Index(cmpID, nil)
```

If you want to limit the results to only the locations that have been updated
since a specific time, provide that time to the `Index` function.

```go
t := time.Date(2019, time.November, 4, 11, 0, 0, 0, time.UTC)

locs, err := c.Locations.Index(cmpID, t)
```
The result is stored in `locs` of type `[]Location`.


### Creating An Entity

To create a new entity, use the `Create` function.

Take `Characters` for example.

For this service, `Create` requires a campaign ID and the data for the new 
character in the form of type `SimpleCharacter`.

```go
ch := SimpleCharacter{
    Name: "Daenerys Targaryen",
    Sex: "Female",
    Title: "Mother of Dragons",
}

_, err := c.Characters.Create(cmpID, ch)
```
The `Create` functions return the newly created entity back to the caller.

This example simply discards the value.

### Updating An Entity

To update an existing entity, use the `Update` function.

Take `Items` for example.

For this service, `Update` requires a campaign ID, an item ID, and the data for
the new item in the form of type `SimpleItem`.

```go
item := SimpleItem{
    Name: "Bag of Holding",
    Size: "15 pounds",
    Price: "300 gold",
}

_, err := c.Items.Update(cmpID, item)
```
The `Update` functions return the updated entity back to the caller.

This example simply discards the value.


### Deleting An Entity

To delete an entity, use the `Delete` function.

Take `Journals` for example.

For this service, `Delete` requires a campaign ID and a journal ID.

```go
err := c.Journals.Delete(cmpID, jrnID)
```

### Rate Limits, Errors, And You

The Kanka API is rate limited. For the most accurate and updated information,
please visit the Kanka [documentation](https://kanka.io/en-US/docs/1.0/setup#endpoints). 

If one of your requests to the Kanka API fails due to the rate limit or other temporary reason,
the error returned can be asserted for the `Temporary` behavior.

For more information about temporary errors, please visit Dave Cheney's
[blog](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully).

## Contributions

If you would like to contribute to this project, please adhere to the following
guidelines.

* Submit an issue describing the problem.
* Fork the repo and add your contribution.
* Add appropriate tests.
* Run go fmt, go vet, and golint.
* Prefer idiomatic Go over non-idiomatic code.
* Follow the basic Go conventions found [here](https://github.com/golang/go/wiki/CodeReviewComments).
* If in doubt, try to match your code to the current codebase.
* Create a pull request with a description of your changes.


