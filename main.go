package main

import (
	"flag"
	"fmt"
	"log"

	"path/filepath"

	"github.com/deckarep/gosx-notifier"
	"go.roman.zone/passport/government"
)

const (
	assetsDir = "assets"
)

var (
	applicationID = flag.String("id", "", "Application ID")
)

func main() {
	flag.Parse()

	application, err := government.GetApplication(*applicationID)
	if err != nil {
		log.Fatalf("Failed to get info about application: %s", err)
	}
	fmt.Println(application.Info.Status.PassportStatus.StatusDescEng)

	note := gosxnotifier.NewNotification("🛂 " + application.Info.Status.PassportStatus.StatusDescEng)
	if application.PassportReady {
		note.Title = "Your passport is READY"
		note.Sound = gosxnotifier.Default
	} else {
		note.Title = "Your passport is not ready yet"
		note.Sound = gosxnotifier.Basso
	}
	note.Group = "zone.roman.go.passport"
	note.Link = government.Scheme + "://" + government.Host
	note.AppIcon = filepath.Join(assetsDir, "russia.jpg")
	err = note.Push()
	if err != nil {
		log.Fatalf("Uh oh! %s", err)
	}
}
