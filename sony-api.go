package main

import (
    "bytes"
		"flag"
    "fmt"
		"io/ioutil"
    "net/http"
    "net/url"
    "path"
		"strings"
)

func main() {
		var endpoint string
		var action string
    var help bool

		actions := []string{"actTakePicture"}

		flag.StringVar(&endpoint, "endpoint", "", "The endpoint")
		flag.StringVar(&action, "action", "", "What action to perform. Supported actions are: " + strings.Join(actions, ", "))
		flag.BoolVar(&help, "help", false, "Show the usage instructions")
		flag.Parse()

		if help {
				PrintHelp("")
		} else if endpoint == "" {
				PrintHelp("'endpoint' is required.")
		} else if action == "" {
				PrintHelp("'action' is required.")
		} else if IsActionValid(action, actions) {
				switch action {
				case "actTakePicture":
						actTakePicture(endpoint)
				}
		} else {
				PrintHelp("Action: '" + action + "' is not yet supported. Supported actions are: " + strings.Join(actions, ", "))
		}
}

func PrintHelp(err string) {
		if err != "" {
        fmt.Println("Error: ", err)
		    fmt.Println()
    }
		flag.PrintDefaults()
}

func IsActionValid(action string, valid_actions []string) bool {
		for _, valid_action := range valid_actions {
			if action == valid_action { return true }
    }
    return false
}

func actTakePicture(endpoint string) {
    action_url, err := url.Parse(endpoint)
    if err != nil {
        panic(err)
    }
    action_url.Path = path.Join(action_url.Path, "/sony/camera")

    jsonStr := []byte("{ 'method': 'actTakePicture', 'params': [], 'id': 1, 'version': '1.0' }")
    req, err := http.NewRequest("POST", action_url.String(), bytes.NewBuffer(jsonStr))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
