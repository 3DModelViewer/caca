package cc

import (
	"errors"
	"fmt"
	"github.com/robsix/golog"
	. "github.com/robsix/json"
	"io"
	"net/http"
)

func NewCCClient(ccHost string, log golog.Log) CCClient {
	return &ccClient{
		host:         ccHost,
		log:          log,
	}
}

type ccClient struct {
	host               string
	log                golog.Log
}

func (c *ccClient) RegisterSheet(base6fUrnAndManifestPath string) (*Json, error) {
	res, err := regsiterSheet(c.host, base6fUrnAndManifestPath)
	if err != nil {
		c.log.Error("CC.RegisterSheet Error: %v", err)
	} else {
		c.log.Info("CC.RegisterSheet Success")
	}
	return res, err
}

func (c *ccClient) GetSheetRegistration(regId string) (*Json, error) {
	res, err := getSheetRegistration(c.host, regId)
	if err != nil {
		c.log.Error("CC.GetSheetRegistration Error: %v", err)
	} else {
		c.log.Info("CC.GetSheetRegistration Success")
	}
	return res, err
}

func (c *ccClient) RegisterClashTest(leftRegId string, rightRegId string) (*Json, error) {
	res, err := registerClashTest(c.host, leftRegId, rightRegId)
	if err != nil {
		c.log.Error("CC.RegisterClashTest Error: %v", err)
	} else {
		c.log.Info("CC.RegisterClashTest Success")
	}
	return res, err
}

func (c *ccClient) GetClashTest(clashTestId string) (*Json, error) {
	res, err := getClashTest(c.host, clashTestId)
	if err != nil {
		c.log.Error("CC.GetClashTest Error: %v", err)
	} else {
		c.log.Info("CC.GetClashTest Success")
	}
	return res, err
}


/**
 * helpers
 */

func checkResponse(resp *http.Response, err error) error {
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		body, _ := FromReadCloser(resp.Body)
		bodyStr, _ := body.ToString()
		return errors.New(fmt.Sprintf("statusCode: %d, status: %v, body: %v", resp.StatusCode, resp.Status, bodyStr))
	}
	return nil
}

func newRequest(method string, urlStr string, body io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	return req, nil
}

func doAdhocJsonRequest(req *http.Request) (ret *Json, err error) {
	resp, err := getStandardHttpClient().Do(req)
	if resp != nil {
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		err = checkResponse(resp, err)
		if err == nil {
			ret, err = FromReadCloser(resp.Body)
		}
	}
	return
}
