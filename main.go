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
	userAgent = "go-weeb/v0.0.3 - (https://github.com/KurozeroPB/go-weeb)"
	baseURL   = "https://rra.ram.moe"
	typePath  = "/i/r?type="
	typeList  = "cry, cuddle, hug, kiss, lewd, lick, nom, nyan, owo, pat, pout, rem, slap, smug, stare, tickle, triggered, nsfw-gtn, potato, kermit"
)

func get(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Accept", "application/vnd.api+json")
	request.Header.Set("Content-Type", "application/vnd.api+json")

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

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Expected status %d; Got %d\nResponse: %#v", 200, response.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
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
