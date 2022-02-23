package main

import (
  "flag"
  "fmt"
  "os"
  "net/http"
  "io/ioutil"
  "log"
  "encoding/json"
  "bytes"
)

/*
type Flag struct {
  name string
  default_val string
  message string
}

type Command struct {
  name string
  flags []Flag
}
*/

type loginRequest struct {
  Email string
  Password string
}

func main() {
  os.Setenv("URL", "localhost:3000")
  /*
  commands := map[string]Command{
    "help": {
      name: "help",
    },
    "login": {
      name: "login",
      flags: []Flag{
        {
          name: "Email",
          message: "Email used to sign in",
        },
        {
          name: "Password",
          message: "Password used to sign in",
        },
      },
    },
  }
  */


  if len(os.Args) < 2 {
    fmt.Println("Expected a valid subcommand. Do takoyaki-cli help")
    os.Exit(1)
  }

  switch os.Args[1] {
  case "help":
    helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
    helpCmd.Parse(os.Args[2:])
    fmt.Println("help bro.")
  case "login":
    loginCmd := flag.NewFlagSet("login", flag.ExitOnError)
    loginEmail := loginCmd.String("email", "", "Email used to sign in")
    passwordEmail := loginCmd.String("password", "", "Password used to sign in")
    loginCmd.Parse(os.Args[2:])

    // Login logic

    postBody, _ := json.Marshal(loginRequest{
      Email: *loginEmail,
      Password: *passwordEmail,
    })

    responseBody := bytes.NewBuffer(postBody)

    response, err := http.Post(os.Getenv("URL"), "application/json", responseBody)

    if err != nil {
      log.Fatalf("An error occurred %v", err)
    }

    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)

    if err != nil {
      log.Fatalf("An error occurred when parsing: %v", err)
    }

    sb := string(body)
    log.Printf(sb)

  default:
    fmt.Println("Expected a valid subcommand. Do takoyaki-cli help")
    os.Exit(1)

  }
}

