package main

import (
    "fmt"
    "os"

    "github.com/sirupsen/logrus"
    "github.com/urfave/cli/v2"
    "github.com/voje/stayinshape/golang/scribe/internal/scribe"
)

func main() {
    log := logrus.New()

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

        Commands: []*cli.Command{
            {
                Name: "sync",
                Action: func(ctx *cli.Context) error {
                    fmt.Println("Scribe starting")
                    b, err := scribe.DecodeSvcaccJSON(
                        ctx.String("svcacc-json-gpg"),
                        ctx.String("svcacc-json-gpg-pass"),
                        )
                    if err != nil {
                        return err
                    }
                    s := scribe.NewScribe(log, b, ctx.String("smtp-pass"))
                    s.SyncCalendars()
                    return nil
                },
            }, 
            {
                Name: "list",
                Action: func(ctx *cli.Context) error {
                    fmt.Println("List events")
                    b, err := scribe.DecodeSvcaccJSON(
                        ctx.String("svcacc-json-gpg"),
                        ctx.String("svcacc-json-gpg-pass"),
                        )
                    if err != nil {
                        return err
                    }
                    s := scribe.NewScribe(log, b, ctx.String("smtp-pass"))
                    s.ListEvents()
                    return nil
                },
            }, 
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
