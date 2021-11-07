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

func TestGetTestCaseIOData(t *testing.T) {
	testData := []struct {
		testNumber int
		inputData  []string
		outputData string
	}{
		{1, []string{"12345"}, "54321"},
		{2, []string{"123", "123"}, "54321"},
	}

	ast := assert.New(t)

	for _, data := range testData {
		inputData, outputData, err := getTestCaseIOData(data.testNumber, getDirWithTestData())

		ast.Equal(data.inputData, inputData)
		ast.Equal(data.outputData, outputData)
		ast.NoError(err)
	}
}

func TestGetTestCaseIODataForTestWithNoIOData(t *testing.T) {
	inputData, outputData, err := getTestCaseIOData(3, getDirWithTestData())

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
