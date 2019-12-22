package cmd

import (
	"fmt"

	"github.com/irvifa/meetup-twitter-api/pkg/generator"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// NewGenerateCommand returns the "generate" command
func NewGenerateCommand() *cobra.Command {
	opts := &generator.Options{}
	cmd := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "Generate a tweet based on the YAML",
		Run:     RunGen(opts),
	}

	addGenFlags(cmd.PersistentFlags(), opts)
	return cmd
}

func addGenFlags(fs *pflag.FlagSet, opts *generator.Options) {
	fs.StringVar(&opts.TweetFile, "tweet-file", "tweet.yaml", "Point to the tweet.yaml file")
	fs.StringVar(&opts.AccessToken, "access-token", "accesstoken", "Twitter Access Token.")
	fs.StringVar(&opts.AccessTokenSecret, "access-token-secret", "accesstokensecret", "Twitter Access Token Secret.")
	fs.StringVar(&opts.ConsumerKey, "consumer-key", "consumerkey", "Twitter Consumer Key.")
	fs.StringVar(&opts.ConsumerSecret, "consumer-secret", "consumersecret", "Twitter Consumer Secret.")
}

func RunGen(opts *generator.Options) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		tweet, client, err := generator.Generate(opts)
		if err != nil {
			log.Fatal(err)
		}

		tweetString := fmt.Sprintf("Dont forget to come to %s on %s, register on %s.", tweet.Meetup.Name, tweet.Meetup.Date, tweet.Meetup.URL)
		print(tweetString)

		_, resp, err := client.Statuses.Update("A Test Tweet from a new Bot I'm building!", nil)
		if err != nil {
			log.Println(err)
		}
		log.Printf("%+v\n", resp)
	}
}
