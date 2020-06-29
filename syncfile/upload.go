package syncfile

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type AssetFile struct {
	ContentType, ContainerName, AccountName, accountKey string
}

func (a *AssetFile) SetAccountVars() (string, string) {
	a.AccountName, a.accountKey = os.Getenv("AZURE_STORAGE_ACCOUNT"), os.Getenv("AZURE_STORAGE_ACCESS_KEY")

	if len(a.AccountName) == 0 || len(a.accountKey) == 0 {
		log.Fatal("Either the AZURE_STORAGE_ACCOUNT or AZURE_STORAGE_ACCESS_KEY environment variable is not set")
	}
	return a.AccountName, a.accountKey
}

func (a *AssetFile) GetCrendials() *azblob.SharedKeyCredential {
	credential, err := azblob.NewSharedKeyCredential(a.AccountName, a.accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}

	return credential
}

func (a *AssetFile) ParseContainerURL() azblob.ContainerURL {
	credential := a.GetCrendials()
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", a.AccountName, a.ContainerName))

	containerURL := azblob.NewContainerURL(*URL, p)
	return containerURL
}

func (a *AssetFile) UploadToBlob(fileName string) {
	urlContainer := a.ParseContainerURL()
	urlBlob := urlContainer.NewBlockBlobURL(fileName)

	f, err := os.Open(fileName)

	ctx := context.Background()

	if err != nil {
		log.Fatal(err)
	}

	urlProperties := azblob.BlobHTTPHeaders{ContentType: a.ContentType}

	_, err = azblob.UploadFileToBlockBlob(ctx, f, urlBlob, azblob.UploadToBlockBlobOptions{
		BlockSize:       4 * 1024 * 1024,
		BlobHTTPHeaders: urlProperties,
		Parallelism:     16})

	if err != nil {
		log.Fatal(err)
	}

	f.Close()
}
