# swapi-BE-template 

## A docker mongo container with the raw Star Wars API data preloaded

Ever needed a fresh database filled with almost universally meaningful dummy data for getting your head round a new stack or technology?

https://swapi.dev/ and before that https://swapi.co are great APIs for this and are a nice way to get dummy data for a frontend app...

...but if you are interested in building a backend project and want something universally meaningful as dummy data then a prepopulated mongo docker might seem interesting.

## How to use

Look at `database/makefile` and `database/main.go` for a full understanding of what is going on.

To get a container running: - `make fresh-run` - the first time you run it.

The `makefile` comments should help you with all the context needed to overcome any other issues!

If you are running local the container should be accessible with something like:

```
mongodb://127.0.0.1:27017/swapi
mongodb://localhost:27017/swapi
```

Obviously this is just the *raw data* and not a full replication of the API responses you'll find in the swapi app thats where you come in.

Enjoy - May the force be with you (cringe)...
