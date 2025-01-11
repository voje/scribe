package scribe

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type Scribe struct {
    log *logrus.Logger
    svcAccJSON []byte
    smtpPass string

    timeInterval time.Duration
    calDecet string
    calStern string

    calSvc *calendar.Service
}

func NewScribe(log *logrus.Logger, svcAccJSON []byte, smtpPass string) *Scribe {
    s := Scribe{
        log: log,
        svcAccJSON: svcAccJSON,
        smtpPass: smtpPass,
        timeInterval: time.Duration(24 * time.Hour * 365),
        calDecet: "garaznidecet@gmail.com",
        calStern: "reteskesternce@gmail.com",
    }

    s.initCalendarService()

    return &s
}

func (s* Scribe) initCalendarService() {
    ctx := context.Background()
    svc, err := calendar.NewService(ctx, option.WithCredentialsJSON(s.svcAccJSON))
    if err != nil {
        s.log.Fatalf("Failed to create a service: %s", err)
    }
    s.calSvc = svc
}

func (s* Scribe) SyncCalendars() {
    s.log.Infof("Synchronizing events from %s into %s", s.calStern, s.calDecet)

    sgdEvents := filterSgdEvents(s.getCalEvents(s.calStern).Items)
    s.log.Info("Š+GD Šternce events:")
    s.logEvents(sgdEvents)

    for _, event := range(sgdEvents) {
        s.insertOrPatchEvent(s.calDecet, event)
    }

    // Š+GD events from CAL_DECET
    s.log.Info("Š+GD Decet events")
    s.logEvents(filterSgdEvents(s.getCalEvents(s.calDecet).Items))
}

func (s* Scribe) ListEvents() {
    tmin, _ := time.Parse("2.1.2006", "1.1.2024")
    tmax, _ := time.Parse("2.1.2006", "1.1.2025")
    fmt.Printf("%s [%s,%s]\n--------------------\n",
        s.calDecet, tmin, tmax,
    )
    events := s.getCalEventsCustomTimeInterval(s.calDecet, tmin, tmax).Items
    for _, e := range events {
        if e.Summary == "Decet vaja" {
            continue
        }
        t, _ := parseEventStartTime(e)
        fmt.Printf(
            "%s|%s %s\n",
            t.Format("02.01.2006 15:04"),
            e.Summary,
            strings.Replace(e.Description, "\n", " ", -1),
        )
    }
}
