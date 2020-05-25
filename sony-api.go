package main

import (
    "bytes"
    "encoding/json"
		"flag"
    "fmt"
    "io"
		"io/ioutil"
    "net/http"
    "net/url"
    "os"
    "path"
		"strings"
)

func main() {
		var endpoint string
		var action string
    var output_picture string
    var output_picture_url string
    var help bool

		actions := []string{"actTakePicture"}

		flag.StringVar(&endpoint, "endpoint", "", "The endpoint")
		flag.StringVar(&action, "action", "", "What action to perform. Supported actions are: " + strings.Join(actions, ", "))
		flag.StringVar(&output_picture, "output_picture", "", "If present, after taking a picture it will be saved as the given file name.")
		flag.StringVar(&output_picture_url, "output_picture_url", "", "If present, after taking a picture it will save the picture url at given file name.")
		flag.BoolVar(&help, "help", false, "Show the usage instructions")
		flag.Parse()

		if help {
				PrintHelp("")
		} else if action == "" {
				PrintHelp("'action' is required.")
		} else if endpoint == "" {
				PrintHelp("'endpoint' is required.")
		} else if IsActionValid(action, actions) {
				switch action {
				case "actTakePicture":
						res := API_actTakePicture(endpoint)
            if output_picture != "" && res.Status == "200 OK" {
                SavePicture(res.Result[0][0], output_picture)
            }
						if output_picture_url != "" && res.Status == "200 OK" {
                SavePictureURL(res.Result[0][0], output_picture_url)
						}
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

// From: https://golangcode.com/download-a-file-from-a-url/
func DownloadFile(url string, filepath string) error {

		// Get the data
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Create the file
		out, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		return err
}


func SavePicture(image_url string, output_picture string) error {
		return DownloadFile(image_url, output_picture)
}

func SavePictureURL(image_url string, output_file string) error {
    image_url_data := []byte(image_url)
    return ioutil.WriteFile(output_file, image_url_data, 0644)
}

type API_actTakePicture_Result struct {
    Status string
		Result [][]string
    Id int
}

func API_actTakePicture(endpoint string) API_actTakePicture_Result {
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
    body_data, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body_data))

		var result API_actTakePicture_Result
		result.Status = resp.Status
		if resp.Status == "200 OK" {
				json.Unmarshal(body_data, &result)
		}

		return result
}
