package service

import (
	"log"
	"mikhael-project-go/internal/usecases"

	"time"

	"github.com/robfig/cron/v3"
)

type (
	CronJob interface {
		StartSchedulerSendEmail()
	}

	cronJob struct {
		SchedulerService usecases.SchedulerService
	}
)

func NewCronJob(schedulerService usecases.SchedulerService) CronJob {

	return &cronJob{
		SchedulerService: schedulerService,
	}
}

// Todo -> untuk panggil
func (s *cronJob) StartSchedulerSendEmail() {
	c := cron.New(cron.WithLocation(time.FixedZone("Asia/Jakarta", 7*3600)))
	// menit / jam / hari
	_, err := c.AddFunc("27 15 * * *", func() {
		log.Println("Running job: SendEmailProductJob")
		// Todo -> Panggil function nya
		s.SchedulerService.SendEmailProduct()
	})
	if err != nil {
		log.Fatal("Failed to add cron job:", err)
	}

	c.Start()
}
