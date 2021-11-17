package algo_tester

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
	"time"
)

var inputFileNotFoundErr = errors.New("input file not found")

const (
	inputFileTemplate = "%s/test.%d.in"
	outputFileTemplate = "%s/test.%d.out"
)

type Task func(data []string) string

func RunTests(t *testing.T, task Task, pathToDirWithIOData string) error {
	if err := checkIfDirWithIODataExists(pathToDirWithIOData); err != nil {
		return err
	}

	testNumber := 0
	for {
		inputData, outputData, err := getTaskIOTestData(testNumber, pathToDirWithIOData)
		if errors.Is(err, inputFileNotFoundErr) {
			fmt.Println(">> No more test data, finishing..")
			break
		}
		if err != nil {
			return errors.Wrap(err, "get test case io data")
		}



		testPassed, testDuration := runTest(t, task, trimInputData(inputData), outputData)

		fmt.Printf(">> Test #%d passed: %t, test duration: %d ms\n\n", testNumber, testPassed, testDuration.Milliseconds())

		testNumber++
	}

	return nil
}

func runTest(t *testing.T, task Task, inputData []string, expected string) (bool, time.Duration) {
	testStartedAt := time.Now()

	result := task(inputData)

	testFinishedAt := time.Now()

	return assert.Equal(t, expected, result), testFinishedAt.Sub(testStartedAt)
}

func checkIfDirWithIODataExists(pathToDirWithIOData string) error {
	_, err := os.Stat(pathToDirWithIOData)
	return err
}

func getTaskIOTestData(testNumber int, pathToDirWithIOData string) ([]string, string, error) {
	pathToInputFile := fmt.Sprintf(inputFileTemplate, pathToDirWithIOData, testNumber)
	pathToOutputFile := fmt.Sprintf(outputFileTemplate, pathToDirWithIOData, testNumber)

	if _, err := os.Stat(pathToInputFile); err != nil {
		return nil, "", inputFileNotFoundErr
	}
	if _, err := os.Stat(pathToOutputFile); err != nil {
		return nil, "", errors.Wrap(err, "get output file stat")
	}

	inputData, err := os.ReadFile(pathToInputFile)
	if err != nil {
		return nil, "", errors.Wrap(err, "read input file")
	}
	outputData, err := os.ReadFile(pathToOutputFile)
	if err != nil {
		return nil, "", errors.Wrap(err, "read output file")
	}

	return strings.Split(string(inputData), "\n"), string(outputData), nil
}

func trimInputData(data []string) []string {
	trimmed := make([]string, 0, len(data))

	for _, d := range data {
		trimmed = append(trimmed, strings.Trim(d, "\n\r"))
	}

	return trimmed
}