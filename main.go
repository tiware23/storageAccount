package main

import (
	"fmt"
	"os"

	"flag"

	f "github.com/tiware23/storageAccount/syncfile"
)

func setAccountStorage(fileStorage *f.AssetFile) {

	fileStorage.SetAccountVars()
	fileStorage.GetCrendials()
	fileStorage.ParseContainerURL()

}

func uploadFile(fileUpload f.AssetFile) {
	fmt.Println("Uploading files...")
	fmt.Println("")

	fileName := flag.Args()
	for _, f := range fileName {
		fmt.Println("Uploading the file: ", f, "to Blob: ", fileUpload.ContainerName)
		fmt.Println(f)
		fileUpload.UploadToBlob(f)
	}
}

func main() {

	commFile := flag.Int("ftype", 1, "Choose Option")
	accName := flag.String("stacc", "cards", "Storage Account Name")
	contType := flag.String("ctype", "application/json", "File's content type")
	contName := flag.String("ctame", "static", "Container Name")

	flag.Parse()

	switch *commFile {
	case 1:
		jsonFile := f.AssetFile{ContentType: *contType, ContainerName: *contName, AccountName: *accName}
		setAccountStorage(&jsonFile)
		uploadFile(jsonFile)
		os.Exit(0)

	case 2:
		textFile := f.AssetFile{ContentType: *contType, ContainerName: *contName, AccountName: *accName}
		setAccountStorage(&textFile)
		uploadFile(textFile)
		os.Exit(0)

	default:
		fmt.Println("No option available")
		os.Exit(-1)

	}
}
