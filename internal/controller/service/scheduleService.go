package service

import (
	"binary_tree/internal/errors"
	"binary_tree/internal/model"
	"binary_tree/internal/model/dto"

	"gorm.io/gorm"

	"net/http"
)

type ScheduleService interface {
	GetMySchedules(userID uint) ([]model.Schedule, int, error)
	GetSchedules(userID uint) ([]model.Schedule, int, error)
	GetMyCoupleSchedules(userID uint) ([]model.Schedule, int, error)
	CreateSchedule(userID uint, createScheduleDTO dto.CreateScheduleDTO) (uint, int, error)
}

type scheduleService struct {
	DB *gorm.DB
}

func NewScheduleService(db *gorm.DB) ScheduleService {
	return &scheduleService{
		DB: db,
	}
}

// 사용자가 작성한 캘린더을 조회
func (s *scheduleService) GetMySchedules(userID uint) ([]model.Schedule, int, error) {
	var schedules []model.Schedule

	if err := s.DB.Where("author_id = ?", userID).Find(&schedules).Error; err != nil {
		return nil, http.StatusInternalServerError, errors.ErrCannotFindSchedules
	}

	return schedules, http.StatusOK, nil
}

// 사용자와 사용자의 커플이 서로 작성한 캘린더을 조회
func (s *scheduleService) GetSchedules(userID uint) ([]model.Schedule, int, error) {
	coupleID, err := model.GetCoupleByUserID(s.DB, userID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var schedules []model.Schedule
	if err := s.DB.Where("couple_id = ?", coupleID.ID).Find(&schedules).Error; err != nil {
		return nil, http.StatusInternalServerError, errors.ErrCannotFindSchedules
	}

	return schedules, http.StatusOK, nil
}

// 사용자의 커플이 작성한 캘린더 조회
func (s *scheduleService) GetMyCoupleSchedules(userID uint) ([]model.Schedule, int, error) {
	user, err := model.FindUserByID(s.DB, userID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if user.PartnerID == nil {
		return nil, http.StatusBadRequest, errors.ErrCannotFindCouple
	}

	var schedules []model.Schedule

	if err := s.DB.Where("author_id = ?", user.PartnerID).Find(&schedules).Error; err != nil {
		return nil, http.StatusInternalServerError, errors.ErrCannotFindSchedules
	}

	return schedules, http.StatusOK, nil
}

/* 캘린더/캘린더 추가 */
func (s *scheduleService) CreateSchedule(userID uint, createScheduleDTO dto.CreateScheduleDTO) (uint, int, error) {
	couple, err := model.GetCoupleByUserID(s.DB, userID)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	var createdSchedule model.Schedule
	createdSchedule.CoupleID = couple.ID
	createdSchedule.AuthorID = userID

	createdSchedule.Title = createScheduleDTO.Title
	createdSchedule.Description = createScheduleDTO.Description
	createdSchedule.StartDate = createScheduleDTO.StartDate
	createdSchedule.EndDate = createScheduleDTO.EndDate
	createdSchedule.EventType = createScheduleDTO.EventType
	createdSchedule.RepeatType = createScheduleDTO.RepeatType
	createdSchedule.RepeatUntil = createScheduleDTO.RepeatUntil

	if err := createdSchedule.Save(s.DB); err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return couple.ID, http.StatusCreated, nil
}
