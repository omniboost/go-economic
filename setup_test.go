package economic_test

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-economic"
)

func client() *economic.Client {
	grantToken := os.Getenv("GRANT_TOKEN")
	secretToken := os.Getenv("SECRET_TOKEN")

	client := economic.NewClient(nil)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)
	client.SetGrantToken(grantToken)
	client.SetSecretToken(secretToken)
	return client
}
