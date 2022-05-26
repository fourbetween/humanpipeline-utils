package lambdaappsync

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

type Params struct {
	Region  string
	URL     string
	Method  string
	ReqBody interface{}
}

func Fetch(p Params) ([]byte, error) {
	reqBodyJson, err := json.MarshalIndent(p.ReqBody, "", "  ")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(p.Method, p.URL, bytes.NewReader(reqBodyJson))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	config := aws.Config{Region: aws.String(os.Getenv("REGION"))}
	sess := session.Must(session.NewSession(&config))
	signer := v4.NewSigner(sess.Config.Credentials)
	signer.Sign(req, bytes.NewReader(reqBodyJson), "appsync", os.Getenv("REGION"), time.Now())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(res.Body)
}
