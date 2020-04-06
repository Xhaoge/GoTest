package concurrency

import (
	"context"
	"errors"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

type Election struct {
	electionPath  string
	candidateName string
	election      *concurrency.Election
	session       *concurrency.Session
}

func NewElection(electionPath string, name string, client *clientv3.Client) (*Election, error) {
	s, err := concurrency.NewSession(client, concurrency.WithTTL(1))
	if err != nil {
		return nil, err
	}
	return &Election{
		electionPath:  electionPath,
		candidateName: name,
		session:       s,
		election:      concurrency.NewElection(s, electionPath),
	}, nil
}

// Try to become the leader, return nil if successfully done or some error happened.
func (election *Election) Campaign(ctx context.Context) error {
	if election.session == nil {
		return errors.New("Election session is nil")

	}
	return election.election.Campaign(ctx, election.candidateName)
}

func (election *Election) Close() error {
	if election.session != nil {
		return election.session.Close()
	}

	return errors.New("Closing nil election session")
}

func (election *Election) DoIfLeader(ctx context.Context, handler func()) error {
	defer election.Close()
	if err := election.Campaign(ctx); err != nil {
		return err
	}
	handler()
	return nil
}
