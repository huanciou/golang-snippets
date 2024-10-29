package main

import (
	"io"
	"net/http"
	"sync"
	"testing"
)

func Test_client(t *testing.T) {
	var wg sync.WaitGroup
	const numsOfReq = 50

	for i := 0; i < numsOfReq; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			resp, err := http.Post("http://localhost:3000/ping", "", nil)
			if err != nil {
				t.Error(err)
				return
			}

			body, _ := io.ReadAll(resp.Body)
			defer resp.Body.Close()

			t.Logf("%s", body)
		}()
	}

	wg.Wait()
}
