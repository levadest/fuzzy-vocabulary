package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	_ "net/http/pprof"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"model"
)

const (
	dictionaryPath          = "src/model/vocabulary.txt"
	vocabularyWordDelimiter = "\n"
)

var (
	inputSourceFile  string
	inputConcurrency int
	inputSourceWords []string
	vocabularyWords  []string
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.IntVar(&inputConcurrency, "concurrency", 500, model.FLAG_CONCURRENCY)
	flag.StringVar(&inputSourceFile, "filepath", "", model.FLAG_FILENAME)
	flag.Parse()

	if inputSourceFile == "" {
		panic(model.ERROR_NO_INPUT_FILE)
	}

	dictionaryFilePath, err := filepath.Abs(dictionaryPath)
	if err != nil {
		panic(fmt.Sprintf(model.ERROR_CANT_READ_FILEPATH, err.Error()))
	}
	vocabularyFile, err := ioutil.ReadFile(dictionaryFilePath)
	if err != nil {
		panic(err.Error())
	}
	vocabularyWords = strings.Split(strings.ToLower(string(vocabularyFile)), vocabularyWordDelimiter)

	sourceFuzzyWordsFile, err := ioutil.ReadFile(inputSourceFile)
	if err != nil {
		panic(fmt.Sprintf("%s: %s", model.ERROR_CANT_READ_FILE, err.Error()))
	}
	inputSourceWords = strings.Fields(string(sourceFuzzyWordsFile))

}

func main() {

	fmt.Printf(model.MESSAGE_START, inputConcurrency, runtime.NumCPU())

	timeNow := time.Now()

	tasks := fillWorkersWithFuzzyWords()
	p := InitWorkerService(tasks, inputConcurrency)
	p.runWorkerService()

	fmt.Printf(model.MESSAGE_WORK, time.Since(timeNow), p.calculateTotalFuzziness())

}
