# golang-env
golang environment

## clone project

```
git clone -b main https://github.com/wachira90/golang-env.git

cd golang-env/

go mod init golang-env
```


## To set the environment variable

```
os.Setenv(key, value)
```

## To get the environment variable

```
value := os.Getenv(key)
```

## code 

```go
package main

import (
  "fmt"
  "os"
)

// use os package to get the env variable which is already set
func envVariable(key string) string {

  // set env variable using os package
  os.Setenv(key, "gopher")

  // return the env variable using os package
  return os.Getenv(key)
}

func main() {
  // os package
  value := envVariable("name")

  fmt.Printf("os package: name = %s \n", value)
  fmt.Printf("environment = %s \n", os.Getenv("APP_ENV"))
}
```

## Run the below command to check.

linux

```sh
APP_ENV=prod go run main.go

// Output
os package: name = gopher
environment = prod
```

windows 

```bat
SETX APP_ENV prod

D:\golang-env>echo %APP_ENV%
prod

D:\golang-env>go run main.go
os package: name = gopher
environment = prod
```

## Install

```
go get github.com/joho/godotenv
```

## godotenv provides a Load method to load the env files.

```
// Load the .env file in the current directory
godotenv.Load()

// or

godotenv.Load(".env")
```


## Create a new .env file in the project root directory

```
STRONGEST_AVENGER=Thor
```

## Update the main.go

```go
package main

import (

    ...
    // Import godotenv
  "github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func main() {
    // os package
    ... 

  // godotenv package
  dotenv := goDotEnvVariable("STRONGEST_AVENGER")

  fmt.Printf("godotenv : %s = %s \n", "STRONGEST_AVENGER", dotenv)
}
```


## Open the terminal and run the main.go.


```
go run main.go

// Output
os package: name = gopher

godotenv : STRONGEST_AVENGER = Thor
```
