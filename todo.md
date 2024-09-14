
### Initialize a Go Module: Run the following command to create a new Go module. This will create a go.mod file in the project root:
```bash
go mod init icepeak
```

### Create Subdirectories: You can create the necessary directories as per our planned structure:

```bash
mkdir -p app/{controllers,models,middlewares,views,services}
mkdir -p bootstrap
mkdir -p config
mkdir -p core/{routing,orm,validation,middleware,cache,logging,response,utils}
mkdir -p database/{migrations,seeds}
mkdir -p public/assets
mkdir -p storage/{logs,uploads,cache}
mkdir -p tests/{integration,unit}
```


### Create Essential Files: Add an empty main.go file as the entry point and a go.mod file to handle dependencies:

```bash
touch main.go
touch bootstrap/init.go
touch config/app.yaml
touch config/database.yaml
touch config/routes.yaml
```

### Next Steps
Now that we have a basic structure in place, we can proceed with setting up the routing component. The routing module will handle incoming HTTP requests and direct them to the appropriate controllers, similar to Laravel's routing system.


Routing
```bash
touch core/routing/router.go
touch core/routing/route.go
touch core/routing/middleware.go
```