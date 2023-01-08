package service

import (
	"fmt"
	"context"

	"github.com/abhayshiro/golang-graphql/graph/model"
	"github.com/abhayshiro/golang-graphql/database"
)

func JobCreate(ctx context.Context, input model.NewJob) (*model.Job, error) {
	db := database.GetDB()
	job := model.Job{
		Text:        input.Text,
		Description: input.Description,
		Tags:        input.Tags,
		Skills:      input.Skills,
	}

	err := db.Model(job).Create(&job)
	
	if err != nil {
		return nil, fmt.Errorf("Error creating a new job")
	}

	return &job, nil
}

func GetJobs() ([]*model.Job, error) {
	db := database.GetDB()
	var jobs []*model.Job

	db.Table("jobs").Find(&jobs)
	
	return jobs, nil
}
