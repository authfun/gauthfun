# gauthfun
RBAC system with tenants and patterns.


## How to use

- Download package

``` bash
go mod tidy
```

- Modiy `config.toml`

``` ini
[service]
port = 8080

[mysql]
user = "{user}"
password = "{password}"
host = "{host}"
port = 3306
db_name = "{db_name}"
``` 

- Start the process

``` bash
go run main.go
```


## Basic Design

![image](https://user-images.githubusercontent.com/6275608/129182060-b0aa5599-d525-44b3-94d0-e8df922ffb1d.png)

- Policy template

  `p, sub, dom, obj, act`

- Group template

  `g, user, role, dom`
  
  `g, member, group, dom`

- Feature holds obj, it should be a policy, like:

  `p, feature, *, menu, *` # all domains with all acts
  
  `p, feature, dom, menu, *` # specific domain with all acts
  
  `p, feature, dom, /api/item, GET` # specific domain with specific act

- Feature inherits feature, it should be a relation or group, like:

  `g, feature, parent-feature, *` # all domains
  
  `g, feature, parent-feature, dom` # specific domain

- Role inherits role, it should be a relation or group, like:

  `g, role, parent-role, *` # all domains
  
  `g, role, parent-role, dom` # specific domain

- Role has features, it should be a relation or group, like:

  `g, role, feature, *` # all domains
  
  `g, role, feature, dom` # specific domain

- User has roles, it should be a relation or group, like:

  `g, user, role, *` # all domains
  
  `g, user, role, dom` # specific domain

- User in organizations, it should be a relation or group, like:

  `g, user, org, *` # all domains
  
  `g, user, org, dom` # specific domain


- Organization binds roles, it should be a relation or group, like:

  `g, org, role, *` # all domains
  
  `g, org, role, dom` # specific domain

- Permission chain

  - A user has a role, the role has a feature, the feature holds an obj, then the user could access the obj.
  
  ```
  # cross domain
  p, feature, *, menu, *
  p, feature, *, /api/items, GET
  g, role, feature, *
  g, user, role, *
  
  user, dom, menu, * -> true
  user, dom, /api/items, GET -> true  
  user, *, menu, * -> true
  user, *, /api/items, GET -> true  

  # specific domain
  p, feature, dom, menu, *
  p, feature, dom, /api/items, GET
  g, role, feature, dom
  g, user, role, dom
  
  user, dom, menu, * -> true
  user, dom, /api/items, GET -> true  
  user, *, menu, * -> false
  user, *, /api/items, GET -> false
  user, dom1, menu, * -> false
  user, dom1, /api/items, GET -> false
  ```
  
  - A user in an organization, the organization binds a role, the role has a feature, the feature holds an obj, then the user could access the obj.
  
  ```
  # cross domain
  p, feature, *, menu, *
  p, feature, *, /api/items, GET
  g, role, feature, *
  g, org, role, *
  g, user, org, *
  
  user, dom, menu, * -> true
  user, dom, /api/items, GET -> true  
  user, *, menu, * -> true
  user, *, /api/items, GET -> true  
  
  # specific domain
  p, feature, dom, menu, *
  p, feature, dom, /api/items, GET
  g, role, feature, dom
  g, org, role, dom
  g, user, org, dom  
    
  user, dom, menu, * -> true
  user, dom, /api/items, GET -> true  
  user, *, menu, * -> false
  user, *, /api/items, GET -> false
  user, dom1, menu, * -> false
  user, dom1, /api/items, GET -> false
  ```
  
