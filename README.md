# go-heroku

### Example minimal golang project api in `Heroku`

### Clone locally
git clone `https://github.com/FerdinaKusumah/go-heroku`

### Run local with
```go
go run main.go
```

Your app should now be running on [localhost:8080](http://localhost:8080/).

### Deployment
    * You can connect your heroku apps with github source and deploy it from specific branch.
    * You can use heroku cli to push project to heroku apps.

### Caveats
    * Don't forget to rename your golang binary name to spesicic folder project name, in file:
        * Procfile
        * Makefile
        * heroku.yml
        * Dockerfile