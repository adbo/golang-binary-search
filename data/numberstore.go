package data

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

type NumberStore struct {
	Numbers []int
	logger  *logrus.Logger
}

func NewNumberStore(logger *logrus.Logger) *NumberStore {
	return &NumberStore{
		logger: logger,
	}
}

func (store *NumberStore) LoadNumbers(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		store.logger.WithError(err).Error("Failed to open the file")
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, numStr := range strings.Fields(line) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				store.logger.WithError(err).Errorf("Failed to convert string to int: %s", numStr)
				return err
			}
			store.Numbers = append(store.Numbers, num)
		}
	}

	if err := scanner.Err(); err != nil {
		store.logger.WithError(err).Error("Error occurred during scanning the file")
		return err
	}

	return nil
}

func (store *NumberStore) FindIndex(value int) (int, bool) {
	low := 0
	high := len(store.Numbers) - 1

	for low <= high {
		mid := low + (high-low)/2

		if store.Numbers[mid] == value {
			return mid, true
		}

		if store.Numbers[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1, false
}

func (store *NumberStore) FindClosestIndex(target int, tolerance float64) (int, bool) {
	low := 0
	high := len(store.Numbers) - 1
	resultIndex := -1

	for low <= high {
		mid := low + (high-low)/2
		midValue := store.Numbers[mid]

		if isInTolerance(midValue, target, tolerance) {
			resultIndex = mid
			break
		}

		if midValue < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return resultIndex, resultIndex != -1
}

func isInTolerance(value, target int, tolerance float64) bool {
	diff := float64(target - value)
	allowedDiff := tolerance * float64(target)
	return -allowedDiff <= diff && diff <= allowedDiff
}
