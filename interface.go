package cc

import (
	. "github.com/robsix/json"
)

type CCClient interface {
	RegisterSheet(b64UrnAndManifestPath string) (*Json, error)
	GetSheetRegistration(registrationId string) (*Json, error)
	RegisterClashTest(leftRegId string, rightRegId string) (*Json, error)
	GetClashTest(clashTestRegId string) (*Json, error)
	RerunClashTest(clashTestRegId string) (*Json, error)
}
