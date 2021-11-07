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
	inputFileTemplate = "%s/test.%d.in.txt"
	outputFileTemplate = "%s/test.%d.out.txt"
)

type TestCase func(data []string) string

func RunTests(t *testing.T, testCase TestCase, pathToDirWithIOData string) error {
	if err := checkIfDirWithIODataExists(pathToDirWithIOData); err != nil {
		return err
	}

	testNumber := 1
	for {
		inputData, outputData, err := getTestCaseIOData(testNumber, pathToDirWithIOData)
		if errors.Is(err, inputFileNotFoundErr) {
			fmt.Println(">> No more test data, finishing..")
			break
		}
		if err != nil {
			return errors.Wrap(err, "get test case io data")
		}

		testPassed, testDuration := runTestCase(t, testCase, inputData, outputData)

		fmt.Printf(">> Test #%d passed: %t, test duration: %d ms\n\n", testNumber, testPassed, testDuration.Milliseconds())

		testNumber++
	}

	return nil
}

func runTestCase(t *testing.T, testCase TestCase, inputData []string, expected string) (bool, time.Duration) {
	testStartedAt := time.Now()

	result := testCase(inputData)

	testFinishedAt := time.Now()

	return assert.Equal(t, expected, result), testStartedAt.Sub(testFinishedAt)
}

func checkIfDirWithIODataExists(pathToDirWithIOData string) error {
	_, err := os.Stat(pathToDirWithIOData)
	return err
}

func getTestCaseIOData(testNumber int, pathToDirWithIOData string) ([]string, string, error) {
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