package scribe

import (
        "fmt"
        "regexp"
        "time"

        log "github.com/sirupsen/logrus"
        "google.golang.org/api/calendar/v3"
        "google.golang.org/api/googleapi"
)

const STAR = "⭐"

func (s *Scribe) getCalEventsCustomTimeInterval(cal_id string, tmin, tmax time.Time) *calendar.Events {
        events, err := s.calSvc.Events.
                List(cal_id).
                ShowDeleted(true).
                SingleEvents(true).
                TimeMin(tmin.Format(time.RFC3339)).
                TimeMax(tmax.Format(time.RFC3339)).
                OrderBy("startTime").
        Do()
        if err != nil {
                log.Fatalf("Could not list events for calendar: %s. %s", cal_id, err)
        }
        return events
}

func (s *Scribe) getCalEvents(cal_id string) *calendar.Events {
        tmin := time.Now()
        tmax := tmin.Add(s.timeInterval)
        return s.getCalEventsCustomTimeInterval(cal_id, tmin, tmax)
}

func (s *Scribe) logEvents(events [](*calendar.Event)) {
        for _, event := range events {
                s.log.Info(stringifyEvent(event))
        }
}

func parseEventStartTime(e *calendar.Event) (time.Time, error) {
        ts := e.Start.DateTime
        if ts == "" {
                ts = e.Start.Date
                return time.Parse("2006-01-02", ts)
        }
        return time.Parse(time.RFC3339, ts)
}

func stringifyEvent(event *calendar.Event) string {
        t := ""
        if event.Start.Date != "" {
                t = event.Start.Date
        } else {
                t = event.Start.DateTime
        }
        return fmt.Sprintf("event: [%s] %s (%s)\n", event.Id, event.Summary, t)
}

func filterSgdEvents(in [](*calendar.Event)) (out [](*calendar.Event)) {
        r := regexp.MustCompile(`Š\+GD`)
        for _, event := range in {
                if r.MatchString(event.Summary) {
                        out = append(out, event)
                }
        }
        return out
}

func (s *Scribe) insertOrPatchEvent(calendarID string, event *calendar.Event) {
        // Add some flavor
        event.Summary = STAR + " " + event.Summary

        // Try inserting
        // Google canendar manually deleted events stay in the system as "cancelled"
        // Copying an event that was deleted will result in an error
        // We want to patch the deleted event, updating its "cancelled" status
        _, err := s.calSvc.Events.Insert(calendarID, event).Do()
        if err != nil {
                switch err.(type) {
                case *googleapi.Error:
                        log.Infof("Patching existing event: %s", event.Id)
                        _, err = s.calSvc.Events.Patch(s.calDecet, event.Id, event).Do()
                        if err != nil {
                                log.Error(err)
                        }
                default:
                        log.Errorf("Inserting event failed: %s", err)
                }
        }
}
