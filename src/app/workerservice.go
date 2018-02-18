package main

import (
	"strings"
	"sync"

	"comparer"
)

type WorkerService struct {
	Tasks []*Job

	concurrency  int
	tasksChannel chan *Job
	wg           sync.WaitGroup
}

type Job struct {
	ResultDimension   int
	fuzzyWord         string
	calculateDistance func(string string) int
}

func InitJob(s string, f func(str string) int) *Job {
	return &Job{fuzzyWord: s, calculateDistance: f}
}

func InitWorkerService(tasks []*Job, concurrency int) *WorkerService {
	return &WorkerService{
		Tasks:        tasks,
		concurrency:  concurrency,
		tasksChannel: make(chan *Job),
	}
}

func (p *WorkerService) runWorkerService() {

	p.wg.Add(len(p.Tasks))
	for i := 0; i < p.concurrency; i++ {
		go p.payload()
	}

	for _, task := range p.Tasks {
		p.tasksChannel <- task
	}
	close(p.tasksChannel)
	p.wg.Wait()
}

func (p *WorkerService) payload() {
	for job := range p.tasksChannel {
		job.Run(&p.wg)
	}
}

func (p *WorkerService) calculateTotalFuzziness() (result int) {
	for _, task := range p.Tasks {
		result += task.ResultDimension
	}
	return result
}

func (t *Job) Run(wg *sync.WaitGroup) {
	t.ResultDimension = t.calculateDistance(t.fuzzyWord)
	wg.Done()
}

func fillWorkersWithFuzzyWords() (tasks []*Job) {

	for _, inputWord := range inputSourceWords {
		tasks = append(tasks, InitJob(inputWord, func(str string) int {
			return comparer.DamerauLevenshtein(strings.ToLower(str), vocabularyWords)
		}))
	}

	return tasks
}
