package main

import (
        "fmt"
	"context"
	"regexp"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
)

const TIME_INTERVAL time.Duration = time.Duration(24 * time.Hour * 365)
const CAL_DECET string = "garaznidecet@gmail.com"
const CAL_STERN string = "reteskesternce@gmail.com"
const STAR = "⭐"

func getCalEvents(svc *calendar.Service, cal_id string) (*calendar.Events) {
        tmin := time.Now()
        tmax := tmin.Add(TIME_INTERVAL)
        events, err := svc.Events.
                List(cal_id).
                ShowDeleted(true).
                SingleEvents(true).
                TimeMin(tmin.Format(time.RFC3339)).
                TimeMax(tmax.Format(time.RFC3339)).
                OrderBy("startTime").
                Do()
        if err != nil {
                log.Fatal("Could not list events for calendar: %s. %s", cal_id, err)
        }
        return events
}

func print_events(events [](*calendar.Event)) {
        for _, event := range events {
                log.Info(stringify_event(event))
        }
}

func stringify_event(event *calendar.Event) string {
        t := ""
        if event.Start.Date != "" {
                t = event.Start.Date 
        } else {
                t = event.Start.DateTime
        }
        return fmt.Sprintf("event: [%s] %s (%s)\n", event.Id, event.Summary, t)
}

func filter_sgd_events(in [](*calendar.Event)) (out [](*calendar.Event)) {
        r := regexp.MustCompile(`Š\+GD`)
        for _, event := range in {
                if r.MatchString(event.Summary) {
                        out = append(out, event)
                }
        }
        return out
}

func insert_or_patch_event(svc *calendar.Service, calendar_id string, event *calendar.Event) {
        // Add some flavor
        event.Summary = STAR + " " + event.Summary
        
        // Try inserting
        // Google canendar manually deleted events stay in the system as "cancelled"
        // Copying an event that was deleted will result in an error
        // We want to patch the deleted event, updating its "cancelled" status
        _, err := svc.Events.Insert(CAL_DECET, event).Do()
        if err != nil {
                switch err.(type) {
                case *googleapi.Error:
                        log.Infof("Patching existing event: %s", event.Id)
                        _, err = svc.Events.Patch(CAL_DECET, event.Id, event).Do()
                        if err != nil {
                                log.Error(err)
                        }
                default:
                        log.Errorf("Inserting event failed: %s", err)
                }
        }
}

func main() {
        // Init google service
        ctx := context.Background()
        svc, err := calendar.NewService(ctx)
        if err != nil {
                log.Fatal("Failed to create a service: %s", err)
        }

        // Š+GD events from CAL_STERN
        sgd_events := filter_sgd_events(getCalEvents(svc, CAL_STERN).Items)
        log.Info("Š+GD Šternce")
        print_events(sgd_events)

        for _, event := range(sgd_events) {
                insert_or_patch_event(svc, CAL_DECET, event)
        }

        // Š+GD events from CAL_DECET
        log.Info("Š+GD Decet events")
        print_events(filter_sgd_events(getCalEvents(svc, CAL_DECET).Items))
}


