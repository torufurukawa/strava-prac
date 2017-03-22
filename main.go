package main

import "github.com/strava/go.strava"
import "os"

func main() {
	client := strava.NewClient(os.Getenv("STRAVA_ACCESS_TOKEN"))
	_ = client
}
