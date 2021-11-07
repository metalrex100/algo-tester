package algo_tester

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestRunTests(t *testing.T) {
	testCase := func(data []string) string {
		fmt.Println(data)

		return "54321"
	}

	err := RunTests(t, testCase, getDirWithTestData())

	assert.NoError(t, err)
}

func TestRunTestsWithMissingDirWithTestData(t *testing.T) {
	testCase := func(data []string) string {
		fmt.Println(data)

		return "54321"
	}

	err := RunTests(t, testCase, "")

	assert.Error(t, err)
}

func TestGetTaskIOTestData(t *testing.T) {
	data := []struct {
		testNumber int
		inputData  []string
		outputData string
	}{
		{1, []string{"12345"}, "54321"},
		{2, []string{"123", "123"}, "54321"},
	}

	ast := assert.New(t)

	for _, d := range data {
		inputData, outputData, err := getTaskIOTestData(d.testNumber, getDirWithTestData())

		ast.Equal(d.inputData, inputData)
		ast.Equal(d.outputData, outputData)
		ast.NoError(err)
	}
}

func TestGetTaskIOTestDataForTaskWithNoIOData(t *testing.T) {
	inputData, outputData, err := getTaskIOTestData(3, getDirWithTestData())

	ast := assert.New(t)

	ast.Nil(inputData)
	ast.Empty(outputData)
	ast.Error(err)
}

func getDirWithTestData() string {
	return fmt.Sprintf("%s/test-data", getPathToCurrentDir())
}

func getPathToCurrentDir() string {
	_, b, _, _ := runtime.Caller(0)

	return filepath.Dir(b)
}
