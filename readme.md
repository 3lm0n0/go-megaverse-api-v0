# Go Simple Microservice.

## Phase 1 - Do a polyantes Cross.
In this phase i was asked to create a diagonal matrix (cross) of polyanets inferring some parameters.
Matrix parameters:
    Total rows: 11,
    Total columns: 11,
    Offset: 2
```
[0 0 0 0 0 0 0 0 0 0 0]
[0 0 0 0 0 0 0 0 0 0 0]
[0 0 1 0 0 0 0 0 1 0 0]
[0 0 0 1 0 0 0 1 0 0 0]
[0 0 0 0 1 0 1 0 0 0 0]
[0 0 0 0 0 1 0 0 0 0 0]
[0 0 0 0 1 0 1 0 0 0 0]
[0 0 0 1 0 0 0 1 0 0 0]
[0 0 1 0 0 0 0 0 1 0 0]
[0 0 0 0 0 0 0 0 0 0 0]
[0 0 0 0 0 0 0 0 0 0 0]
```
## Solution API usage.
```
URL: localhost:port/megaverse/polyanets
    Methods allowed
        POST: To create a single polyanet.
            Body:{
                "Row": 2,
                "Column": 2
            }
        DELETE: To delete a single polyanet.
            Body:{
                "Row": 2,
                "Column": 2
            }
```
```
URL: localhost:port/megaverse/polyanets/cross
    Methods allowed
        POST: To generate cross.
            Body: {
                "rows": 11,
                "columns": 11,
                "offset": 2
            }
```

## Phase 2 - Logo.
```
URL: localhost:port/megaverse/logo
    Methods allowed
        POST: To create a logo.
            Body:{
                "goal":[
                    ["SPACE","SPACE","SPACE", ... ,"SPACE"]
                    ["SPACE","SPACE","SPACE", ... ,"SPACE"]
                    ["SPACE","SPACE","SPACE", ... ,"SPACE"]
                        .       .       .           .
                        .       .       .           .
                        .       .       .           .
                    ["SPACE","SPACE","SPACE", ... ,"SPACE"]
                ]
                
            }
```

# Commands Usage:
```
$ make          --> compiles all binary.
$ make init     --> creates microservice module.
$ make setup    --> sets up the microservice.
$ make build    --> builds the microservice.
$ make run      --> runs the microservice.
$ make clean    --> removes ALL binaries and objects.
```