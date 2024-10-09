package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SocService/pkg/database"
	model "github.com/SocService/pkg/repositories/model"

	"github.com/SocService/pkg/utils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"log"
)

type SocRepository struct {
}

func (s *SocRepository) MakeRelation(from string, to string) error {
	driver, err := database.NeoDriver()
	if err != nil {
		return err
	}
	ctx := context.Background()
	defer func(driver neo4j.DriverWithContext, ctx context.Context) {
		err := driver.Close(ctx)
		if err != nil {
			log.Printf(err.Error())
		}
	}(driver, ctx)

	if err := driver.VerifyConnectivity(ctx); err != nil {
		return err
	}

	result, err := neo4j.ExecuteQuery(ctx, driver, `
	MATCH (from: User{id: $from})
	MATCH (to: User{id: $to})
	CREATE (from)-[:KNOWS]->(to)
	`, map[string]any{
		"from": from,
		"to":   to,
	}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(utils.GetValue("NEO_DATABASE")))
	if err != nil {
		return err
	}
	if !result.Summary.Counters().ContainsUpdates() {
		return fmt.Errorf("can't create relationship")
	}
	return nil
}

func (s *SocRepository) MakeDetach(from string, to string) error {
	driver, err := database.NeoDriver()
	if err != nil {
		return err
	}
	ctx := context.Background()
	if err := driver.VerifyConnectivity(ctx); err != nil {
		return err
	}

	result, err := neo4j.ExecuteQuery(ctx, driver, `MATCH (from: User{id: $from})-[r:KNOWS]->(to: User{id: $to}) DELETE r`,
		map[string]any{
			"from": from,
			"to":   to,
		}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(utils.GetValue("NEO_DATABASE")))
	if err != nil {
		return err
	}

	if !result.Summary.Counters().ContainsUpdates() {
		return fmt.Errorf("can't detach relationship")
	}
	return nil
}

func (s *SocRepository) CreateNode(person model.Person) error {
	ctx := context.Background()
	driver, err := database.NeoDriver()
	if err != nil {
		return err
	}
	defer func(driver neo4j.DriverWithContext, ctx context.Context) {
		err := driver.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(driver, ctx)

	if err := driver.VerifyConnectivity(ctx); err != nil {
		return err
	}

	result, err := neo4j.ExecuteQuery(ctx, driver,
		"CREATE (user: User{id: $id, name: $name, image: $image, isPrivate: $isPrivate})",
		map[string]any{
			"id":        person.Id,
			"name":      person.Name,
			"image":     person.Image,
			"isPrivate": person.IsPrivate,
		}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(utils.GetValue("NEO_DATABASE")),
	)
	if err != nil {
		return err
	}

	summary := result.Summary
	fmt.Printf("Created %v nodes in %+v.\n", summary.Counters().NodesCreated(), summary.ResultAvailableAfter())

	return nil
}

func (s *SocRepository) GetFollowings(id string, offset int, limit int) ([]*model.Person, error) {
	ctx := context.Background()
	driver, err := database.NeoDriver()
	if err != nil {
		return nil, err
	}

	defer func(driver neo4j.DriverWithContext, ctx context.Context) {
		err := driver.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(driver, ctx)

	if err := driver.VerifyConnectivity(ctx); err != nil {
		return nil, err
	}

	result, err := neo4j.ExecuteQuery(ctx,
		driver,
		"MATCH (f: User{id: $from})-[KNOWS]->(d: User) RETURN d LIMIT $limit",
		map[string]any{"from": id, "limit": limit},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(utils.GetValue("NEO_DATABASE")),
	)
	if err != nil {
		return nil, err
	}

	var res []*model.Person
	for _, record := range result.Records {
		jsonData, err := json.Marshal(record.Values[0].(dbtype.Node).Props)
		if err != nil {
			log.Fatal(err.Error())
		}
		var tmp model.Person

		if err := json.Unmarshal(jsonData, &tmp); err != nil {
			log.Fatal(err.Error())
		} else {
			res = append(res, &tmp)
		}
	}

	fmt.Printf("The query `%v` returned %v records in %+v.\n",
		result.Summary.Query().Text(), len(result.Records),
		result.Summary.ResultAvailableAfter())

	return res, nil
}

func (s *SocRepository) GetFollowers(id string, limit int, offset int) ([]*model.Person, error) {
	ctx := context.Background()
	driver, err := database.NeoDriver()
	if err != nil {
		return nil, err
	}
	defer func(driver neo4j.DriverWithContext, ctx context.Context) {
		err := driver.Close(ctx)
		if err != nil {
			log.Fatal(err.Error())
		}
	}(driver, ctx)

	if err := driver.VerifyConnectivity(ctx); err != nil {
		return nil, err
	}

	result, err := neo4j.ExecuteQuery(ctx, driver, "MATCH (f: User)-[r:KNOWS]->(d: User{id: $id}) return d LIMIT $limit", map[string]any{
		"id":    id,
		"limit": limit,
	}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(utils.GetValue("NEO_DATABASE")))
	if err != nil {
		return nil, err
	}

	var res []*model.Person
	for _, record := range result.Records {
		jsonData, err := json.Marshal(record.Values[0].(dbtype.Node).Props)
		if err != nil {
			log.Fatal(err.Error())
		}
		var tmp model.Person

		if err := json.Unmarshal(jsonData, &tmp); err != nil {
			log.Fatal(err.Error())
		} else {
			res = append(res, &tmp)
		}
	}

	fmt.Printf("The query `%v` returned %v records in %+v.\n",
		result.Summary.Query().Text(), len(result.Records),
		result.Summary.ResultAvailableAfter())

	return res, nil
}
