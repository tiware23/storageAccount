package main

import (
	"fmt"
	"os"

	f "github.com/tiware23/storageAccount/syncfile"
)

func main() {

	jsonFile := f.AssetFile{ContentType: "application/json", ContainerName: "static",
		AccountName: "accountName", AccountKey: "accountKey"}

	jsonFile.SetAccountVars()
	jsonFile.GetCrendials()
	jsonFile.ParseContainerURL()

	fileName := os.Args[1:]
	for _, f := range fileName {
		fmt.Println("Uploading the file: ", f)
		jsonFile.UploadToBlob(f)
	}
}
