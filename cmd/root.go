package cmd

import (
	"io"

	"github.com/cloud-native-nordics/meetup-kit/pkg/logs"
	logflag "github.com/cloud-native-nordics/meetup-kit/pkg/logs/flag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var logLevel = logrus.InfoLevel

// NewMeetupApiTwitterCommand returns the root command for ignite
func NewMeetupApiTwitterCommand(in io.Reader, out, err io.Writer) *cobra.Command {

	root := &cobra.Command{
		Use:   "meetup-api-twitter",
		Short: "meetup-api-twitter: Create a Tweet for Meetup",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Set the desired logging level, now that the flags are parsed
			logs.Logger.SetLevel(logLevel)
		},
	}

	addGlobalFlags(root.PersistentFlags())

	root.AddCommand(NewGenerateCommand())
	return root
}

func addGlobalFlags(fs *pflag.FlagSet) {
	logflag.LogLevelFlagVar(fs, &logLevel)
}
