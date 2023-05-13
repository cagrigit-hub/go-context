package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	ctx := context.Background()
	userId := 1
	val, err := fetchUserData(ctx, userId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result: ", val)
	fmt.Println("Time: ", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	respch := make(chan Response)
	go func() {
		val, err := fetchThirdpartyStuffWhichCanBeSlow()
		respch <- Response{val, err}
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case resp := <-respch:
		return resp.value, resp.err
	}
}

func fetchThirdpartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(5 * time.Second)
	return 666, nil
}
