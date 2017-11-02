package weeb

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
)

const (
	userAgent = "go-weeb/v0.0.2 - (https://github.com/KurozeroPB/go-weeb)"
	baseURL   = "https://rra.ram.moe"
	typePath  = "/i/r?type="
	typeList  = "cry, cuddle, hug, kiss, lewd, lick, nom, nyan, owo, pat, pout, rem, slap, smug, stare, tickle, triggered, nsfw-gtn, potato, kermit"
)

func executeRequest(request *http.Request, expectedStatus int) ([]byte, error) {
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != expectedStatus {
		return nil, fmt.Errorf(
			"Expected status %d; Got %d \nResponse: %#v",
			expectedStatus,
			response.StatusCode,
			buf.String(),
		)
	}
	return buf.Bytes(), nil
}

func newRequest(method string, url string) (*http.Request, error) {
	req, err := newUARequest(method, url, userAgent)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func newUARequest(method string, url string, ua string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", ua)
	request.Header.Set("Accept", "application/vnd.api+json")
	request.Header.Set("Content-Type", "application/vnd.api+json")
	return request, nil
}

func safeGET(url string, expectedStatus int) ([]byte, error) {
	req, e := newRequest("GET", url)
	if e != nil {
		return nil, e
	}
	byt, err := executeRequest(req, expectedStatus)
	if err != nil {
		return nil, err
	}
	return byt, nil
}

func get(url string) ([]byte, error) {
	byt, err := safeGET(url, 200)
	if err != nil {
		return nil, err
	}
	return byt, nil
}

// GetImage gets image from the given type
func GetImage(Type string) (string, error) {
	newType := strings.ToLower(Type)

	TypeBool := strings.Contains(typeList, newType)
	if TypeBool == false {
		err := fmt.Errorf("type %s is not a valid option", newType)
		return "", err
	}
	res, err := get(baseURL + typePath + newType)
	if err != nil {
		return "", err
	}
	json, e := gabs.ParseJSON(res)
	if e != nil {
		return "", e
	}
	img := baseURL + json.Path("path").Data().(string)
	return img, nil
}
