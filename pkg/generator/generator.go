package generator

import (
	"io/ioutil"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/irvifa/meetup-twitter-api/api"
	"github.com/irvifa/meetup-twitter-api/pkg/types"
	"sigs.k8s.io/yaml"
)

type Options struct {
	TweetFile         string
	AccessToken       string
	AccessTokenSecret string
	ConsumerKey       string
	ConsumerSecret    string
}

var unmarshal = yaml.UnmarshalStrict

func Generate(opts *Options) (*types.Tweet, *twitter.Client, *types.error) {
	tweet := &types.Tweet{}
	tweetContent, err := ioutil.ReadFile(opts.TweetFile)

	if err != nil {
		return nil, nil, err
	}

	if err := unmarshal(tweetContent, &tweet); err != nil {
		return nil, nil, err
	}

	credentials := &api.Credentials{
		ConsumerSecret:    *opts.ConsumerSecret,
		ConsumerKey:       *opts.ConsumerKey,
		AccessToken:       *opts.AccessToken,
		AccessTokenSecret: *opts.AccessTokenSecret,
	}

	client, err := *api.getClient(&credentials)

	if err != nil {
		return nil, nil, err
	}

	return tweet, client, nil
}
