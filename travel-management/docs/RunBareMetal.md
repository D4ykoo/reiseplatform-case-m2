# How to run - Bare Metal
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
./travelmanagement
```
### Frontend
#### Development
```bash
cd frontend
ng serve --port 8085
```
Use the port 8087 to avoid CORS erros.
#### Production
```bash
cd frontend
npm run build
```
Now the whole application is located in the dist/ directory.<br>
The application can be served by any desired webserver after coping the whole directory to e. g. `var/www/html/`.