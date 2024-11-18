package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

var urls = []string{"hh.ru", "habr.com", "ya.ru"}

func SendRequestToService(ctx context.Context, str string, id int) error {
	client := http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", "https://"+str, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code %d", resp.StatusCode)
	}
	fmt.Printf("%s %d, id: %d\n", str, resp.StatusCode, id)
	return nil
}

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	for i, url := range urls {
		//url, i := url, i
		eg.Go(func() error {
			return SendRequestToService(ctx, url, i)
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal("Error in eGroup")
	}
}
