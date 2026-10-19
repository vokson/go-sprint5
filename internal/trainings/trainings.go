package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const (
	trainingFields     = 3 // Количество полей в строке данных.
	trainingInfoFormat = `Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`
)

var (
	errSplitString = errors.New("failed to split string")
	errStepsCount  = errors.New("steps count must be positive")
	errDuration    = errors.New("duration must be positive")
	// Здесь по-русски из-за тестов
	errUnknownTraining = errors.New("неизвестный тип тренировки")
)

type Training struct {
	personaldata.Personal

	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	arr := strings.Split(datastring, ",")
	if len(arr) != trainingFields {
		return errSplitString
	}

	steps, err := strconv.Atoi(arr[0])
	if err != nil {
		return fmt.Errorf("parsing steps count: %w", err)
	}
	if steps <= 0 {
		return errStepsCount
	}

	duration, err := time.ParseDuration(arr[2])
	if err != nil {
		return fmt.Errorf("parsing duration: %w", err)
	}
	if duration <= 0 {
		return errDuration
	}

	t.Steps = steps
	t.TrainingType = arr[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var caloriesFunc func(int, float64, float64, time.Duration) (float64, error)

	switch t.TrainingType {
	case "Бег":
		caloriesFunc = spentenergy.RunningSpentCalories
	case "Ходьба":
		caloriesFunc = spentenergy.WalkingSpentCalories
	default:
		return "", errUnknownTraining
	}

	calories, err := caloriesFunc(t.Steps, t.Weight, t.Height, t.Duration)
	if err != nil {
		return "", fmt.Errorf("calculating spent calories: %w", err)
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	return fmt.Sprintf(trainingInfoFormat, t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}
