package logger

import (
"github.com/sirupsen/logrus"
"os"
)

var Log = *logrus.New()

func init() {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "production" {
			Log.SetFormatter(&logrus.JSONFormatter{})
			Log.SetLevel(logrus.DebugLevel)
		} else {
			// The TextFormatter is default, you don't actually have to do this.
			Log.SetFormatter(&logrus.TextFormatter{})
			Log.SetLevel(logrus.DebugLevel)
		}
	}

	Log.SetOutput(os.Stdout)
	Log.Out = os.Stdout

	contextLogger := Log.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other": "I also should be logged always",
	})

	Log.WithFields(logrus.Fields{
		"animal": "🐧",
	}).Info("Nobody here but us penguins")
	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	//  log.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }
	contextLogger.Traceln("This I guess is the common field?")
}

