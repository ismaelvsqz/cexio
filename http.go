package cexio

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func readBody(body io.Reader) ([]byte, error) {
	result, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf(errorReadingResponse, err)
		return nil, err
	}
	return result, nil
}

func getMethod(u string) ([]byte, error) {
	res, err := http.Get(u)
	if err != nil {
		log.Printf(errorHTTPComunication, err)
		return nil, err
	}
	defer res.Body.Close()
	return readBody(res.Body)
}

func postMethod(u string, v []byte) ([]byte, error) {
	res, err := http.Post(u, postContentType, bytes.NewBuffer(v))
	if err != nil {
		log.Printf(errorHTTPComunication, err)
		return nil, err
	}
	defer res.Body.Close()
	return readBody(res.Body)
}
