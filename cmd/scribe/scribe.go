package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
        Name:  "scribe",
        Usage: "Sync google calendars, send some e-mails",

        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:    "svcacc-json-gpg",
                Usage:   "Location of the encrypted service account credentials",
                EnvVars: []string{"SCRIBE_SVCACC_JSON_GPG"},
                Required: true,
            },
            &cli.StringFlag{
                Name:    "svcacc-json-gpg-pass",
                Usage:   "Password for decrypting svc-acc.json.gpg",
                EnvVars: []string{"SCRIBE_SVCACC_JSON_GPG_PASS"},
                Required: true,
            },
            &cli.StringFlag{
                Name:    "smtp-pass",
                Usage:   "Password for using Google's SMTP server",
                EnvVars: []string{"SCRIBE_SMTP_PASS"},
                Required: true,
            },
        },

        Action: func(ctx *cli.Context) error {
       		fmt.Println("Scribe starting")
            s := scribe.NewScribe(
                scribe.decodeSvcaccJSON(
                    ctx.String("svcacc-json-gpg"),
                    ctx.String("svcacc-json-gpg-pass"),
                )
            )
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
