package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	_ "net/http/pprof"
	"runtime"
	"strings"
	"time"
	"path/filepath"

	"model"

)

const (
	dictionaryPath          = "src/comparer/vocabulary.txt"
	vocabularyWordDelimiter = "\n"
)

var (
	inputFile        string
	inputConcurrency int
	inputWords       []string
	vocabularyWords  []string
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.IntVar(&inputConcurrency, "concurrency", 500, model.FLAG_CONCURRENCY)
	flag.StringVar(&inputFile, "filepath", "", model.FLAG_FILENAME)
	flag.Parse()

	if inputFile == "" {
		panic(model.ERROR_NO_INPUT_FILE)
	}

	pwd, err := filepath.Abs(dictionaryPath)
	if err != nil {
		panic(fmt.Sprintf(model.ERROR_CANT_READ_FILEPATH,err.Error()))
	}
	vocabularyFile, err := ioutil.ReadFile(pwd)
	if err != nil {
		panic(err.Error())
	}
	vocabularyWords = strings.Split(strings.ToLower(string(vocabularyFile)), vocabularyWordDelimiter)

	sourceFuzzyWordsFile, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(fmt.Sprintf("%s: %s", model.ERROR_CANT_READ_FILE, err.Error()))
	}
	inputWords = strings.Fields(string(sourceFuzzyWordsFile))

}

func main() {

	fmt.Printf(model.MESSAGE_START, inputConcurrency, runtime.NumCPU())

	timeNow := time.Now()

	tasks := FillWorkersWithFuzzyWords()
	p := InitWorkerService(tasks, inputConcurrency)
	p.RunWorkerService()

	fmt.Printf(model.MESSAGE_WORK, time.Since(timeNow), p.CalculateTotalFuzziness())

}
