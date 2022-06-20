# Gofiber Backend Boilerplate
Thanks to [Cozy (ItsCosmas)](https://github.com/ItsCosmas) to build gofiber boilerplate. I start learning gofiber because of its performance.
Also, the following gofiber features could make the great projects.
- Express-inspired
- Robust routing
- Serve static files
- Extreme performance
- API ready
- Flexible middleware support
- Low memory footprint
- Template engine
- Websocket support
- Rate limiter

Running and Developing locally
------------------------------

1.  Create `.env` at src, i.e.

```source-shell
cp src/.env.example src/.env
```

2.  Download Swag for generating docs

```source-shell
go get -u github.com/swaggo/swag/cmd/swag
```

3.  Run

-   NOTE: You have to generate swagger docs before running the app.

```source-shell
# Terminal 1
swag init -g src/api/app.go --output ./src/api/docs # Generates Swagger

# Terminal 2
docker-compose --env-file ./src/.env up        # docker-compose up (Run App With AutoReload)
docker-compose --env-file ./src/.env down      # docker-compose down (Shutdown App)
```

-   API `http://localhost:8000/api/v1`
-   Swagger Doc `http://localhost:8000/api/v1/docs`

* * * * *

[](https://github.com/ItsCosmas/gofiber-boilerplate?ref=https://githubhelp.com#packaging-for-production)Packaging For Production
--------------------------------------------------------------------------------------------------------------------------------

1.  Create `.env` at src, i.e.

```source-shell
cp src/.env.example src/.env
```

2.  Update your `.env` variables for production

-   Point to your prod database
-   Update JWT issuer, secret key , blah blah
-   Basically just follow good production practice

3.  Download Swag for generating docs

```source-shell
go get -u github.com/swaggo/swag/cmd/swag
```

-   Generate Swagger Docs. You have to generate swagger docs before packaging the app.

```source-shell
swag init -g src/api/app.go --output ./src/api/docs # Generates Swagger
```

4.  Build Your Image

-   Permission the build script to run.

```
chmod +x docker-build.sh

```

-   You could set the image port on `Dockerfile.prod`
-   Run the build script. You must provide a version tag as shown below.

```
./docker-build.sh -v gofiber:1.0.0
```
