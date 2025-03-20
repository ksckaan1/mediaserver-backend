package main

import (
	"testing"
)

func TestABC(t *testing.T) {
	// cbURL := "couchbase://localhost"
	// cluster, err := gocb.Connect(cbURL, gocb.ClusterOptions{
	// 	Authenticator: gocb.PasswordAuthenticator{
	// 		Username: "cbadmin",
	// 		Password: "cbpass",
	// 	},
	// })
	// require.NoError(t, err)
	// err = cluster.WaitUntilReady(10*time.Second, nil)
	// require.NoError(t, err)

	// bucketName := "media_server"
	// bucketManager := cluster.Buckets()
	// buckets, err := bucketManager.GetAllBuckets(nil)
	// require.NoError(t, err)

	// _, ok := buckets[bucketName]
	// if !ok {
	// 	err := bucketManager.CreateBucket(gocb.CreateBucketSettings{
	// 		BucketSettings: gocb.BucketSettings{
	// 			Name: bucketName,
	// 		},
	// 	}, nil)
	// 	require.NoError(t, err)
	// }

	// bucket := cluster.Bucket(bucketName)
	// bucket.CollectionsV2().CreateScope("example_scope", nil)
	// t.Log(buckets)
	// cluster.buc
	// bucket := cluster.Bucket(bucketName)

	// // _, err = bucket.DefaultScope().
	// // 	Query(`INSERT INTO tmdb_infos (KEY, VALUE) VALUES ('movie:123123', {"title": "John Wick 3", "description": "Keanu Reis"})`, nil)
	// // require.NoError(t, err)

	// _, err = bucket.Collection("tmdb_infos").Insert("movie:6666666", ExampleData{
	// 	Title:       "John Wick 4",
	// 	Description: "Keanu Reeves",
	// }, nil)
	// require.NoError(t, err)

}

type ExampleData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
