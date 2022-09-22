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

| Name                  | Description                                                                                                                   | Sample Value                           |
|-----------------------|-------------------------------------------------------------------------------------------------------------------------------|----------------------------------------|
| ldap_user_name        | User who the LDAP bind operation will run as for the default directory                                                        | ```jrolstad@internal.salesforce.com``` |
| ldap_user_password    | Password for the user being used for the LDAP bind operation.  If running as yourself, this is your Active Directory password | ```some-supersecret-value!```          |
| ldap_host             | Host Name of the LDAP server connecting to                                                                                    | ```internal.mydomain.com:636```        |
| directoryobject_queue | Name of the queue to publish directory objects to                                                                             | ```my-ingestion-queue```               |
| aws_region            | Region where AWS resources are hosted                                                                                         | ```us-west-2```                        |
| AWS_ACCESS_KEY_ID     | AWS Client Id to authenticate as                                                                                              | ```some value```                       |
| AWS_SECRET_ACCESS_KEY | AWS Client Secret to use when authenticating                                                                                  | ```some secret value```                |
| AWS_SESSION_TOKEN     | AWS Secret Token to use when authenticating                                                                                   | ```some token value```                 |
If you are using [Z Shell](https://en.wikipedia.org/wiki/Z_shell) as your CLI of choice, an example setup for ~/.zshrc is:
```shell
export ldap_host=internal.mydomain.com:636
export ldap_user_name=jrolstad@mydomain.com
export ldap_user_password=some-supersecret-value! 

export directoryobject_queue=identityobject-ingest

export aws_region=us-west-2
export AWS_ACCESS_KEY_ID=value-here
export AWS_SECRET_ACCESS_KEY=secret-here
export AWS_SESSION_TOKEN=token-here
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