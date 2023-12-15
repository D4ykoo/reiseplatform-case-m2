# User Management

## Prerequisites
* A postgresql database where the configuration and credentials are known

## How to run - bare metal

### Configuration
All the configuration is done in the .env of the file. See comments for some additional information.

### Backend
#### Development
```bash
cd backend
go run ./main.go
```
#### Production
```bash
cd backend
go build
./usermanagement
```
### Frontend
#### Development
```bash
cd frontend
npm run dev
```
#### Production
```bash
cd frontend
npm run build
```
Now the whole application is located in the dist/ directory.<br>
The application can be served by any desired webserver by coping the whole directory and renaming it to e.g. usermanagement. 

## Technology Stack
* TypeScript + VueJS 3 + TailwindCSS + daisyUI 
* Go + Gin + GORM

## Architecture
![architecture](./resources/usermangement_architecture.png?raw=true "Usermanagement Architecture")

## Project Structure
The project is separated in frontend and backend directories as well as resources for all the documentation resources.<br>
The following represents a quick and small overview about the microservice structure.
```bash
├── backend
│   ├── adapter
│   ├── application
│   ├── domain
│   │   └── model
│   ├── ports
│   ├── tests
│   ├── utils
│ 
├── frontend
│   ├── public
│   ├── src
│   │   └── assets
│   │   └── components
│   │   └── models
│   │   └── router
│   │   └── services
│   │   └── store
│   │   └── views
│ 
├── resources
├── .env
```

## Additional Information
* Project management done with self-hosted [Leantime](https://github.com/Leantime/leantime)