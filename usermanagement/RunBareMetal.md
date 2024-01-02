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
The application can be served by any desired webserver after coping the whole directory to e. g. `var/www/html/`.