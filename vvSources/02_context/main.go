package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	value int
	err   error
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 150)

	return 666, nil
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respch := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching user data from third party took too long")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

func main() {
	start := time.Now()
	ctx := context.Background()
	userID := 10
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", val)
	fmt.Println("took:", time.Since(start))
}
