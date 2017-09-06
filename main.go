package weeb

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
)

var (
	userAgent = "go-weeb - (https://github.com/KurozeroPB/go-weeb)"
	baseURL   = "https://rra.ram.moe"
	typePath  = "/i/r?type="
	typeList  = "cry, cuddle, hug, kiss, lewd, lick, nom, nyan, owo, pat, pout, rem, slap, smug, stare, tickle, triggered, nsfw-gtn, potato, kermit"
)

func executeRequest(request *http.Request, expectedStatus int) []byte {
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode != expectedStatus {
		panic(fmt.Errorf(
			"Expected status %d; Got %d \nResponse: %#v",
			expectedStatus,
			response.StatusCode,
			buf.String(),
		))
	}
	return buf.Bytes()
}

func newRequest(method string, url string) *http.Request {
	return newUARequest(method, url, userAgent)
}

func newUARequest(method string, url string, ua string) *http.Request {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Set("User-Agent", ua)

	return request
}

func safeGET(url string, expectedStatus int) []byte {
	return executeRequest(
		newRequest("GET", url),
		expectedStatus,
	)
}

func get(url string) []byte {
	return safeGET(url, 200)
}

// GetImage gets image from the given type
func GetImage(Type string) (string, error) {
	newType := strings.ToLower(Type)

	TypeBool := strings.Contains(typeList, newType)
	if TypeBool == false {
		err := fmt.Errorf("type does not exist")
		return "", err
	}
	json, e := gabs.ParseJSON(get(baseURL + typePath + newType))
	img := baseURL + json.Path("path").Data().(string)
	return img, e
}
