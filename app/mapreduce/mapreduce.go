package mapreduce

import (
	"errors"
	"log"

	"github.com/zeromicro/go-zero/core/mr"
)

func Run(uids []int) ([]int, error) {
	r, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, uid := range uids {
			source <- uid
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		uid := item.(int)
		ok, err := check(uid)
		if err != nil {
			cancel(err)
		}
		if ok {
			writer.Write(uid)
		}
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		var uids []int
		for p := range pipe {
			uids = append(uids, p.(int))
		}
		writer.Write(uids)
	})
	if err != nil {
		log.Printf("check error: %v", err)
		return nil, err
	}

	return r.([]int), nil
}

func check(uid int) (bool, error) {
	// do something check user legal
	if uid == 0 {
		return false, errors.New("uid wrong")
	}
	return true, nil
}
