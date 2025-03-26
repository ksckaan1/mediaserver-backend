package service

import (
	"context"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
)

func (s *Service[CFG]) initCouchbaseBucket(ctx context.Context) error {
	if s.ServiceCfg.CouchbaseHost == "" || s.ServiceCfg.CouchbaseUser == "" || s.ServiceCfg.CouchbasePassword == "" || s.ServiceCfg.CouchbaseBucket == "" {
		return nil
	}
	var err error
	s.cluster, err = gocb.Connect(fmt.Sprintf("couchbase://%s", s.ServiceCfg.CouchbaseHost), gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: s.ServiceCfg.CouchbaseUser,
			Password: s.ServiceCfg.CouchbasePassword,
		},
	})
	if err != nil {
		return fmt.Errorf("gocb.Connect: %w", err)
	}
	s.Logger.Info(ctx, "trying to connect to couchbase cluster")
	err = s.cluster.WaitUntilReady(10*time.Second, &gocb.WaitUntilReadyOptions{
		Context: ctx,
	})
	if err != nil {
		return fmt.Errorf("cluster.WaitUntilReady: %w", err)
	}
	s.Logger.Info(ctx, "successfully connected to couchbase cluster")
	s.CBBucket = s.cluster.Bucket(s.ServiceCfg.CouchbaseBucket)
	s.Logger.Info(ctx, "trying to connect to couchbase bucket")
	err = s.CBBucket.WaitUntilReady(10*time.Second, &gocb.WaitUntilReadyOptions{
		Context: ctx,
	})
	if err != nil {
		return fmt.Errorf("bucket.WaitUntilReady: %w", err)
	}
	s.Logger.Info(ctx, "successfully connected to couchbase bucket")
	return nil
}
