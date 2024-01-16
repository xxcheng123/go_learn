package main

import "fmt"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/16 10:52
 */

type Subject interface {
	Raining()
}

type Student struct {
}

func (s *Student) Raining() {
	fmt.Println("The student go home with an umbrella.")
}

type Worker struct {
}

func (w *Worker) Raining() {
	fmt.Println("Worker drive the car go home.")
}

type Observer interface {
	AddListener(s Subject)
	RemoveListener(s Subject)
	Notify()
}
type RainObserver struct {
	Subjects []Subject
}

func (r *RainObserver) AddListener(s Subject) {
	r.Subjects = append(r.Subjects, s)
}
func (r *RainObserver) RemoveListener(s Subject) {
	for i, v := range r.Subjects {
		if v == s {
			r.Subjects = append(r.Subjects[:i], r.Subjects[i+1:]...)
			break
		}
	}
}
func (r *RainObserver) Notify() {
	for _, v := range r.Subjects {
		v.Raining()
	}
}
func main() {
	s := &Student{}
	w := &Worker{}
	r := &RainObserver{}
	r.AddListener(s)
	r.AddListener(w)
	r.RemoveListener(s)
	r.Notify()
}
