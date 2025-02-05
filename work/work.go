package work

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
)

// Job 是一个示例任务，接受一个整数参数 id

type Worker struct {
	Pool *ants.Pool
}

func New(count int) *Worker {
	return &Worker{
		Pool: nil,
	}
}

func (w *Worker) C创建(count int) (err error) {
	w.Pool, err = ants.NewPool(count)
	return err
}

func (w *Worker) T提交任务(job func()) (err error) {
	if w.Pool == nil || w.Pool.IsClosed() {
		return fmt.Errorf("this pool has been closed")
	}
	return w.Pool.Submit(job)
}

func (w *Worker) T停止() {
	if w.Pool != nil && !w.Pool.IsClosed() {
		fmt.Println("停止")
		w.Pool.Release()
	}
}

func (w *Worker) S是否关闭() bool {
	if w.Pool == nil {
		return true
	}
	return w.Pool.IsClosed()
}

func (w *Worker) H获取线程池数量() int {
	if w.Pool == nil {
		return 0
	}
	fmt.Println(w.Pool.Running())
	fmt.Println(w.Pool.Waiting())
	fmt.Println(w.Pool.Free())
	return w.Pool.Cap()
}

func (w *Worker) H获取剩余任务数() int {
	if w.Pool == nil {
		return 0
	}
	fmt.Println(w.Pool.Waiting())
	return w.Pool.Waiting()
}
