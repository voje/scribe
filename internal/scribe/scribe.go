package scribe

import (
	"context"
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
                s.log.Fatal("Failed to create a service: %s", err)
        }
        s.calSvc = svc
}

func (s* Scribe) SyncCalendars() {
    s.log.Info("Synchronizing events from %s into %s", s.calStern, s.calDecet)

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

