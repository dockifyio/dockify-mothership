package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/dockifyio/dockify-mothership/api"
	"github.com/dockifyio/dockify-mothership/pkg/Utilities"
	"os"
)

func main() {
	var app api.App
	var vaultToken = os.Getenv("VAULT_TOKEN")
	var vaultAddr = os.Getenv("VAULT_ADDR")
	fireBaseApiVaultPath := "firebase/data/keys"
	fireBaseApiKeyName := "api_key"
	myFigure := figure.NewFigure("Dockify Mothership", "", true)
	myFigure.Print()
	// setup vault
	vaultClient, err := Utilities.InitVault(vaultAddr)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Couldn't initialize vault client")
		os.Exit(1)
	}
	fireBaseApiKey, err := Utilities.GetValuesFromVaultV2Api(vaultClient, vaultToken, fireBaseApiVaultPath, fireBaseApiKeyName)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Couldn't get values from vault v2 api for firebase api key")
		os.Exit(1)
	}
	app.Initialize(fireBaseApiKey)
	app.Run()
}
