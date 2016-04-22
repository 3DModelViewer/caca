package cc

import (
	sj "github.com/robsix/json"
)

func registerClashTest(host string, leftRegId string, rightRegId string) (ret *sj.Json, err error) {
	body, err := sj.New()
	if err != nil {
		return nil, err
	}
	body.Set(leftRegId, "clash_test", "left_id")
	body.Set(rightRegId, "clash_test", "right_id")

	reader, err := body.ToReader()
	if err != nil {
		return nil, err
	}

	req, err := newRequest("POST", host+"/api/v1/clash_tests", reader, "application/json")
	if err != nil {
		return nil, err
	}

	ret, err = doAdhocJsonRequest(req)
	return
}
