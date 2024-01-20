# Travelplatform CASE-M2
Project for the lecture Cloud Native Development.
As the name suggests it is a travel management platform for creating hotels, travels and book them. Everything packed with a usermanagement!

![overview_architecture.png](ressources/overview_architecture.png)
## Overview
### User Management
This hexagonal service takes care of login, register, user operations (create, read, update, delete) as well as JWT operations.

**Technology Stack:**  
Frontend: VueJS3 + Typescript + Vite + TailwindCSS + DaisyUI  
Backend: Go + Gin + GORM

Author: Dario Köllner

### Travel Management
The travel management is a hexagonal service, which has CRUD operations for hotels and corresponding travels. Each hotel can have several travels.
They can be stored inside a cart provided by the checkout service.  

**Technology Stack:**  
Frontend: Angular + Typescript  
Backend: Go + Gin + GORM

Author: Michael Gürtner

### Checkout
The checkout service has, depending on the user, all cart information is capable of storing, creating and deleting it.  
The payment process is mocked.

**Technology Stack:**  
Frontend: Angular + Typescript + TailwindCSS + DaisyUI  
Backend: Rust + Actix Web + Diesel

Author: Dario Köllner

### Monitoring
This service receives logs provided from the services mentioned above and displays them.

**Technology Stack:**  
Frontend: Angular + Typescript  
Backend: Rust + Diesel

Author: Michael Gürtner

## Project Structure
Every service is seperated in frontend and backend directories. Each microservice has its own documentation as well as some docker-compose files.
Besides, there is in the root of this repository a [docker](docker) and [kubernetes](kubernetes) directory. Both dirs contain the corresponding deployments.

Each service can be run isolated - though that does not mean they will function properly.

For further information about the services have a look at the READMEs located there.

## Quick Start
#### Docker
```bash
cd docker
```

Create docker networks:
```bash
chmod +x network_management.sh && ./network_management.sh -g
```

Start all containers:
```bash
chmod +x start-container.sh && ./start-container.sh
```
Further information and instructions can be found in the [README](docker/README.md).

#### Kubernetes
TODO

#### Run in VM
Is documented for every service. 
## Default Ports when deployed
| **Service**                | **Exposed Port** | **Internal Port** |
|----------------------------|------------------|-------------------|
| Usermanagement Frontend    | 8081             | 80                |
| Usermanagement Backend     | 8082             | 8082              |
| Postgres Usermanagement    | 8092             | 5432              |
|                            |                  |                   |
| Checkout Frontend          | 8083             | 80                |
| Checkout Backend           | 8084             | 8084              |
| Postgres Checkout          | 8094             | 5432              |
|                            |                  |                   |
| Travelmanagement Frontend  | 8085             | 8085              |
| Travelmanagement Backend   | 8086             | 8086              |
| Postgres Travelmanagement  | 8096             | 5432              |
|                            |                  |                   |
| Monitoring Frontend        | 8087             | 8087              |
| Monitoring Backend         | 8088             | 8088              |
| Postgres Monitoring        | 8098             | 5432              |


## Project Insights 
```bash
❯ cloc --exclude-dir=node_modules,target -exclude-ext=json .

github.com/AlDanial/cloc v 1.98  T=0.20 s (1854.1 files/s, 95300.9 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
YAML                           118            342           1825           4294
Go                              59            720            159           2661
TypeScript                      72            264             94           2223
Rust                            18            312             56           1498
HTML                            19             10              8           1246
Markdown                        30            289              0           1198
Vuejs Component                 12             60             14            923
CSS                              6             40             10            245
Bourne Shell                     5             30              3            153
JavaScript                      11             10             18            132
TOML                            10             17              9            102
Dockerfile                       7             33              7             98
SQL                              6             13             26             66
INI                              1              3              0             13
-------------------------------------------------------------------------------
SUM:                           374           2143           2229          14852
-------------------------------------------------------------------------------

```