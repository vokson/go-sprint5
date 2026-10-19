package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	personaldata.Personal

	Steps    int
	Duration time.Duration
}

const (
	expectedFields  = 2
	dayActionFormat = `Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`
)

var (
	errSplitString = errors.New("failed to split string")
	errStepsCount  = errors.New("steps count must be positive")
	errDuration    = errors.New("duration must be positive")
)

func (ds *DaySteps) Parse(datastring string) (err error) {
	pair := strings.Split(datastring, ",")
	if len(pair) != expectedFields {
		return errSplitString
	}

	steps, err := strconv.Atoi(pair[0])
	if err != nil {
		return fmt.Errorf("parsing steps count: %w", err)
	}
	if steps <= 0 {
		return errStepsCount
	}

	duration, err := time.ParseDuration(pair[1])
	if err != nil {
		return fmt.Errorf("parsing duration: %w", err)
	}
	if duration <= 0 {
		return errDuration
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("calculating calories: %w", err)
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)
	return fmt.Sprintf(dayActionFormat, ds.Steps, distance, calories), nil
}
