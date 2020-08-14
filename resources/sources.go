package resources

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/droxolite/drx"
	"net/http"
	"strings"
)

type Source struct {
	client *http.Client
	req    *http.Request
}

func APIResource(clnt *http.Client, r *http.Request) *Source {
	return &Source{
		client: clnt,
		req:    r,
	}
}

func (src *Source) get(api, path string, params ...string) (interface{}, error) {
	tkninfo := drx.GetIdentity(src.req)
	url, err := tkninfo.GetResourceURL(api)

	if err != nil {
		return nil, err
	}

	fullURL := fmt.Sprintf("%s/%s/%s", url, path, strings.Trim(strings.Join(params, "/"), "/"))

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	req.Header.Set("Authorization", "Bearer "+drx.GetToken(src.req))

	if err != nil {
		return nil, err
	}

	resp, err := src.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result interface{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, nil
}
