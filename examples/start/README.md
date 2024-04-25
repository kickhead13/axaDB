# axa start utility

## Scope
Boots a DataBase Dispetcher server that hosts a buffer of user commands to be executed safely on the DataBase.

## Usage
```sh
 $ axa start --at /my/db/directory --as adminUserName:adminPassword --on dispetcherIP:discpetcherPort
(axa server): listening on 127.0.0.1:13131...
(axa server): new connection: (127.0.0.1:56584)
(axa server) received [ feed in users {"kickhead13":{"password":"kickheadsVERYsecurepassword","email":"kickheadsVERYsecretemail@email.com"}} ] from 127.0.0.1:56584 
(axa server): connection to 127.0.0.1:56584 closed...
```

