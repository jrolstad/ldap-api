# LDAP Identity Api
A simple API wrapper around the Lightweight Directory Access Protocol (LDAP).  Meant to showcase how data from Active Directory or any other LDAP instance can be exposed via API and not direct LDAP searches.

## Requirements
* golang 1.18 or higher

## Components
### Executables
|Name| Description                                                                               |
|---|-------------------------------------------------------------------------------------------|
|cmd/api-server| Main API endpoints for the service.  Hosted using [Gin](https://github.com/gin-gonic/gin) |

### Packages
|Name| Description                                                        |
|---|--------------------------------------------------------------------|
|internal/pkg/configuration| Functions to read application configurations and secrets           |
|internal/pkg/core| Package agnostic functionality that runs on native types only      |
|internal/pkg/directory| Functions and services that interact with LDAP directory instances |
|internal/pkg/models|Structs that represent data used in the service|
|internal/pkg/orchestration|Functions that orchestrate the features and processes enabled by the service|

## API Endpoints
The following endpoints are exposed on port 8080 when the api-server code is ran

|Address|Description|Example|
|---|---|---|
| /{directory}/user/{name}|Gets a specific user by name| /mydomain/user/jrolstad|
|/{directory}/group/{name}|Gets details about a specific group|/mydomain/group/all_users|
|/{directory}/group/{name}/member|Gets all members in the specified group|/mydomain/group/all_users/member|

## Development Environment Setup
### Configuration Values
The following configuration values need to be added to your environment variables.

|Name|Description| Sample Value                     |
|---|---|----------------------------------|
|ldap_user_name|User who the LDAP bind operation will run as for the default directory| ```jrolstad@internal.salesforce.com``` |
|ldap_user_password|Password for the user being used for the LDAP bind operation| ```some-supersecret-value!```          |

Example setup for ~/.zshrc
```shell
export ldap_user_name=jrolstad@internal.salesforce.com
export ldap_user_password=some-supersecret-value! 
```

### How to Run
Once values are configured, the service can be executed from the _/cmd/api-server_ directory by running:
```shell
go run *.go
```

Once the gin server is up and running, GET calls can be made via cURL or in a browser
* http://localhost:8080/internal.salesforce.com/user/jrolstad
* http://localhost:8080/internal.salesforce.com/group/mygroupname
* http://localhost:8080/internal.salesforce.com/group/mygroupname/member