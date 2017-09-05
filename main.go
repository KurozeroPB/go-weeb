package weeb

import (
	"bytes"
	"fmt"
	"go-weeb/version"
	"io"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
)

var (
	// USERAGENT Set UA
	USERAGENT = "go-weeb/" + version.BOT_VERSION + " (https://github.com/KurozeroPB/go-weeb)"
	baseURL   = "rra.ram.moe"
	typePath  = "/i/r?type="
	typeList  []string
)

// executeRequest Executes a http request
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

// newRequest Creates a new request
func newRequest(method string, url string) *http.Request {
	return newUARequest(method, url, USERAGENT)
}

// newUARequest Adds a custom user agent
func newUARequest(method string, url string, ua string) *http.Request {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Set("User-Agent", ua)

	return request
}

// SafeGET Makes GET request
func SafeGET(url string, expectedStatus int) []byte {
	return executeRequest(
		newRequest("GET", url),
		expectedStatus,
	)
}

// GET calls the SafeGET func to make a GET request
func GET(url string) []byte {
	return SafeGET(url, 200)
}

// TypeInList check if type is in the list
func TypeInList(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// GetImage gets image
func GetImage(Type string) (string, error) {
	typeList[0] = "cry"
	typeList[1] = "cuddle"
	typeList[2] = "hug"
	typeList[3] = "kiss"
	typeList[4] = "lewd"
	typeList[5] = "lick"
	typeList[6] = "nom"
	typeList[7] = "nyan"
	typeList[8] = "owo"
	typeList[9] = "pat"
	typeList[10] = "pout"
	typeList[11] = "rem"
	typeList[12] = "slap"
	typeList[13] = "smug"
	typeList[14] = "stare"
	typeList[15] = "tickle"
	typeList[16] = "triggered"
	typeList[17] = "nsfw-gtn"
	typeList[18] = "potato"
	typeList[19] = "kermit"

	newType := strings.ToLower(Type)
	TypeBool := TypeInList(newType, typeList)
	if TypeBool == false {
		err := fmt.Errorf("type does not exist")
		return "", err
	}
	json, e := gabs.ParseJSON(GET(baseURL + typePath + newType))
	img := baseURL + json.Path("path").Data().(string)
	return img, e
}
