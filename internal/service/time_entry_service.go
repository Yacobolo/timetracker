package service

import (
	"context"
	"timetracker/internal/db"
	"timetracker/internal/repository"
)

type TimeEntryService interface {
	CreateTimeEntry(ctx context.Context, input db.CreateTimeEntryParams) (db.TimeEntry, error)
	DeleteTimeEntry(ctx context.Context, id int32) error
	GetTimeEntry(ctx context.Context, id int32) (db.TimeEntry, error)
	ListTimeEntries(ctx context.Context) ([]db.TimeEntry, error)
}

type timeEntryService struct {
	repo repository.TimeEntryRepository
}

func NewTimeEntryService(repo repository.TimeEntryRepository) TimeEntryService {
	return &timeEntryService{repo: repo}
}

func (s *timeEntryService) CreateTimeEntry(ctx context.Context, input db.CreateTimeEntryParams) (db.TimeEntry, error) {

	timeEntry, err := s.repo.CreateTimeEntry(ctx, input)

	if err != nil {
		return db.TimeEntry{}, err
	}

	return timeEntry, nil
}

func (s *timeEntryService) DeleteTimeEntry(ctx context.Context, id int32) error {
	return s.repo.DeleteTimeEntry(ctx, id)
}

func (s *timeEntryService) GetTimeEntry(ctx context.Context, id int32) (db.TimeEntry, error) {
	return s.repo.GetTimeEntry(ctx, id)
}

func (s *timeEntryService) ListTimeEntries(ctx context.Context) ([]db.TimeEntry, error) {
	return s.repo.ListTimeEntries(ctx)
}
