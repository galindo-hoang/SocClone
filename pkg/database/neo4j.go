package database

import (
	"fmt"
	"github.com/SocService/pkg/utils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func NeoDriver() (neo4j.DriverWithContext, error) {
	var (
		dbUri  = fmt.Sprintf("%v://%v", utils.GetValue("NEO_DATABASE"), utils.GetValue("NEO_HOST"))
		dbUser = utils.GetValue("NEO_USERNAME")
		dbPwd  = utils.GetValue("NEO_PASSWORD")
	)
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth(dbUser, dbPwd, ""))
	if err != nil {
		return nil, err
	}
	return driver, nil
}
