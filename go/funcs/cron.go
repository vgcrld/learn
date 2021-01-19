package funcs

import (
	"fmt"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func RunCron() {

	// Add a wait to keep us running
	wg.Add(1)

	// Create the cron table and add entries
	c := cron.New()
	defer stopCron(*c)
	addCronEntries(c)

	// Start and wait
	pmsg("Starting...")
	c.Start()
	wg.Wait()
}

func addCronEntries(c *cron.Cron) {
	c.AddFunc("*/1 * * * * ", func() { pmsg("every 1 second") })
	c.AddFunc("*/11 * * * * ", func() { pmsg("every 11 seconds") })
	c.AddFunc("*/66 * * * * ", func() { pmsg("every 66 seconds") })
	c.AddFunc("0 30 * * * ", func() { pmsg("every xx:30") })
}

func pmsg(msg string) {
	fmt.Println(msg, time.Now())
}

func printEntries(c *cron.Cron) {
	for i, v := range c.Entries() {
		fmt.Println("job: " + fmt.Sprint(i))
		spew.Dump(v)

	}
}

func stopCron(c cron.Cron) {
	fmt.Println("Stopping the cron function.")
	c.Stop()
}

func logMessage1(msg string) {
	f := log.Fields{
		"name": "richard",
	}
	log.WithFields(f).Info(msg)
}
