package main

import (
	root_http "github.com/SocService/pkg/handlers/http"
	root_rpc "github.com/SocService/pkg/handlers/rpc"
)

func main() {
	go root_rpc.NewServerHandler()
	root_http.InitRoutes()
}

//func main() {
//	ctx := context.Background()
//	dbUri := "neo4j://localhost"
//	dbUser := "neo4j"
//	dbPass := "12345678"
//	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth(dbUser, dbPass, ""))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer func(driver neo4j.DriverWithContext, ctx context.Context) {
//		err := driver.Close(ctx)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}(driver, ctx)
//
//	err = driver.VerifyConnectivity(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Connected to neo4j cluster")
//
//	//createDB("Alice1", driver, ctx)
//	//createDB("Alice2", driver, ctx)
//	//createDB("Alice3", driver, ctx)
//	//createDB("Bob", driver, ctx)
//	//queryDB(driver, ctx)
//
//	//updateDB(driver, ctx)
//
//	//createRelationship("Alice", "Alice1", driver, ctx)
//	//createRelationship("Alice", "Alice2", driver, ctx)
//	//createRelationship("Alice", "Alice3", driver, ctx)
//	//createRelationship("Alice", "Bob", driver, ctx)
//	//detachRelationship(driver, ctx)
//
//	queryListDB(driver, ctx)
//}
//
//func createDB(name string, driver neo4j.DriverWithContext, ctx context.Context) {
//	result, err := neo4j.ExecuteQuery(ctx, driver,
//		"CREATE (p:Person {name: $name}) RETURN p",
//		map[string]any{
//			"name": name,
//		}, neo4j.EagerResultTransformer,
//		neo4j.ExecuteQueryWithDatabase("neo4j"))
//	if err != nil {
//		panic(err)
//	}
//	summary := result.Summary
//	fmt.Printf("Created %v nodes in %+v.\n",
//		summary.Counters().NodesCreated(),
//		summary.ResultAvailableAfter())
//}
//
//func queryDB(driver neo4j.DriverWithContext, ctx context.Context) {
//	result, err := neo4j.ExecuteQuery(ctx, driver,
//		"MATCH (p:Person) return p.name as name",
//		nil,
//		neo4j.EagerResultTransformer,
//		neo4j.ExecuteQueryWithDatabase("neo4j"),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, record := range result.Records {
//		//name, _ := record.Get("name")
//		name := record.AsMap()["name"].(string)
//		fmt.Println(name)
//	}
//
//	fmt.Printf("The query `%v` returned %v records in %+v.\n",
//		result.Summary.Query().Text(), len(result.Records),
//		result.Summary.ResultAvailableAfter())
//}
//
//func queryListDB(driver neo4j.DriverWithContext, ctx context.Context) {
//	result, err := neo4j.ExecuteQuery(ctx, driver,
//		"MATCH (f: Person{name: $from})-[r:KNOWS]->(d: Person) return d LIMIT $limit ",
//		map[string]any{"from": "Alice"},
//		neo4j.EagerResultTransformer,
//		neo4j.ExecuteQueryWithDatabase("neo4j"),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, record := range result.Records {
//		name := record.Values[0].(dbtype.Node).Props
//		fmt.Println(name)
//	}
//
//	fmt.Printf("The query `%v` returned %v records in %+v.\n",
//		result.Summary.Query().Text(), len(result.Records),
//		result.Summary.ResultAvailableAfter())
//}
//
//func updateDB(driver neo4j.DriverWithContext, ctx context.Context) {
//	result, err := neo4j.ExecuteQuery(ctx, driver, `
//    MATCH (p:Person {name: $name})
//    SET p.age = $age
//    `, map[string]any{
//		"name": "Alice",
//		"age":  42,
//	}, neo4j.EagerResultTransformer,
//		neo4j.ExecuteQueryWithDatabase("neo4j"))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Query updated the database?",
//		result.Summary.Counters().ContainsUpdates())
//}
//
//func createRelationship(from string, to string, driver neo4j.DriverWithContext, ctx context.Context) {
//	result, err := neo4j.ExecuteQuery(ctx, driver, `
//    MATCH (alice:Person {name: $name})
//    MATCH (bob:Person {name: $friend})
//    CREATE (alice)-[:KNOWS]->(bob)
//    `, map[string]any{
//		"name":   from,
//		"friend": to,
//	}, neo4j.EagerResultTransformer,
//		neo4j.ExecuteQueryWithDatabase("neo4j"))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Query updated the database?",
//		result.Summary.Counters().ContainsUpdates())
//}
//
//func detachRelationship(from string, to string, driver neo4j.DriverWithContext, ctx context.Context) {
//	result, err := neo4j.ExecuteQuery(ctx, driver, `
//    MATCH (alice:Person {name: $name})-[r:KNOWS]->(bob:Person {name: $friend})
//    DELETE r
//    `, map[string]any{
//		"name":   from,
//		"friend": to,
//	}, neo4j.EagerResultTransformer,
//		neo4j.ExecuteQueryWithDatabase("neo4j"))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Query updated the database?",
//		result.Summary.Counters().ContainsUpdates())
//}
