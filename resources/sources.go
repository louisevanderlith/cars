package resources

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/droxolite/context"
	"net/http"
	"strings"
)

type Source struct {
	client *http.Client
	ctx    context.Contexer
}

func APIResource(clnt *http.Client, ctx context.Contexer) *Source {
	return &Source{
		client: clnt,
		ctx:    ctx,
	}
}

func (src *Source) get(api, path string, params ...string) (interface{}, error) {
	tkninfo := src.ctx.GetTokenInfo()
	url, err := tkninfo.GetResourceURL(api)

	if err != nil {
		return nil, err
	}

	fullURL := fmt.Sprintf("%s/%s/%s", url, path, strings.Trim(strings.Join(params, "/"), "/"))

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	req.Header.Set("Authorization", "Bearer "+src.ctx.GetToken())

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
