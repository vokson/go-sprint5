package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var (
	errStepsCount = errors.New("steps count must be positive")
	errDuration   = errors.New("duration must be positive")
	errHeight     = errors.New("height must be positive")
	errWeight     = errors.New("weight must be positive")
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	runningCalories, err := RunningSpentCalories(steps, weight, height, duration)

	if err != nil {
		return 0, err
	}

	return walkingCaloriesCoefficient * runningCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errStepsCount
	}
	if weight <= 0 {
		return 0, errWeight
	}
	if height <= 0 {
		return 0, errHeight
	}
	if duration <= 0 {
		return 0, errDuration
	}
	speed := MeanSpeed(steps, height, duration)
	return (weight * speed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return stepLengthCoefficient * height * float64(steps) / mInKm
}
