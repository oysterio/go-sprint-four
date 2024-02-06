package ftracker

import (
	"fmt"
	"math"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// action — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// action — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration — длительность тренировки в часах.
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	distance := distance(action)
	return distance / duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// action — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// trainingType — вид тренировки(Бег, Ходьба, Плавание).
// duration — длительность тренировки в часах.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	// ваш код здесь
	switch {
	case trainingType == "Бег":
		distance := distance(action)                               // вызовите здесь необходимую функцию
		speed := meanSpeed(action, duration)                       // вызовите здесь необходимую функцию
		calories := RunningSpentCalories(action, weight, duration) // вызовите здесь необходимую функцию
		return fmt.Sprintf("Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case trainingType == "Ходьба":
		distance := distance(action)                                       // вызовите здесь необходимую функцию
		speed := meanSpeed(action, duration)                               // вызовите здесь необходимую функцию
		calories := WalkingSpentCalories(action, duration, weight, height) // вызовите здесь необходимую функцию
		return fmt.Sprintf("Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case trainingType == "Плавание":
		distance := distance(action)                                               // вызовите здесь необходимую функцию
		speed := swimmingMeanSpeed(lengthPool, countPool, duration)                // вызовите здесь необходимую функцию
		calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight) // вызовите здесь необходимую функцию
		return fmt.Sprintf("Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18   // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 1.79 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// action — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// weight — вес пользователя.
// duration — длительность тренировки в часах.
func RunningSpentCalories(action int, weight, duration float64) float64 {
	// ваш код здесь
	durationMin := duration * minInH
	return ((runningCaloriesMeanSpeedMultiplier * meanSpeed(action, duration) * runningCaloriesMeanSpeedShift) * weight / mInKm * durationMin)
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// action — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration — длительность тренировки в часах.
// weight — вес пользователя.
// height — рост пользователя.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	// ваш код здесь
	meanSpeed := meanSpeed(action, duration) * kmhInMsec
	durationMin := duration * minInH
	return ((walkingCaloriesWeightMultiplier*weight + (math.Pow(meanSpeed, 2)/height*cmInM)*walkingSpeedHeightMultiplier*weight) * durationMin)
}

// Константы для расчета калорий, расходуемых при плавании.
const (
	swimmingCaloriesMeanSpeedShift   = 1.1 // среднее количество сжигаемых колорий при плавании относительно скорости.
	swimmingCaloriesWeightMultiplier = 2   // множитель веса при плавании.
)

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
//
// Параметры:
//
// lengthPool — длина бассейна в метрах.
// countPool — сколько раз пользователь переплыл бассейн.
// duration — длительность тренировки в часах.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool) * float64(countPool) / mInKm / duration
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
//
// Параметры:
//
// lengthPool — длина бассейна в метрах.
// countPool — сколько раз пользователь переплыл бассейн.
// duration — длительность тренировки в часах.
// weight — вес пользователя.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	// ваш код здесь
	swimmingMeanSpeed := swimmingMeanSpeed(lengthPool, countPool, duration)
	return (swimmingMeanSpeed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}
