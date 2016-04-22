package cc

import (
	sj "github.com/robsix/json"
)

func regsiterSheet(host string, base64UrnAndMaifestPath string) (ret *sj.Json, err error) {
	body, err := sj.New()
	if err != nil {
		return nil, err
	}
	body.Set(base64UrnAndMaifestPath, "model", "urn")

	reader, err := body.ToReader()
	if err != nil {
		return nil, err
	}

	req, err := newRequest("POST", host+"/api/v1/models", reader, "application/json")
	if err != nil {
		return nil, err
	}

	ret, err = doAdhocJsonRequest(req)
	return
}
