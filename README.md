# olaf

2020 monorepo for convention over configuration microservices

# goals

- simple conventions for microservices
- development deployments
- prod deployments

# layout

/frontend - hosts all next.js frontend apps
/backend - backend go apps / services / functions

# conventions for backend services

## layout of a backend service group

```
backend/cmd/<name> this is where the program is
backend/service/<name> this is where the main service package should be

```



