package syncer

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/multisig-labs/pandasia/pkg/db"
	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	_, queries := db.OpenDB("../../pandasia.db")

	err := SyncPChain(context.Background(), queries, "http://100.83.243.106:9650", nil)
	require.NoError(t, err)

	t.Fatal()
}

func Test2(t *testing.T) {
	c := colly.NewCollector(colly.AllowURLRevisit())
	c.OnResponse(func(r *colly.Response) {
		t.Logf("%s", r.Body)
		// rs := &map[string]interface{}{}
		// err := json.Unmarshal(r.Body, rs)
		// require.NoError(t, err)
		// t.Logf("%v", rs)
	})

	q, _ := queue.New(
		1, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)
	body := strings.NewReader(fmt.Sprintf(`{"id":0,"jsonrpc":"2.0","method":"platform.getBlockByHeight","params":{"height":"0","encoding":"hex"}}`))
	u, _ := url.Parse("http://100.83.243.106:9650/ext/bc/P")

	r := &colly.Request{
		URL:     u,
		Headers: &http.Header{"Content-Type": {"application/json"}},
		Method:  "POST",
		Body:    body,
	}
	err := q.AddRequest(r)
	require.NoError(t, err)
	err = q.Run(c)
	require.NoError(t, err)
	t.Fatal()
}

func Test3(t *testing.T) {
	mc := MutexCounter{}
	ch := make(chan *colly.Response, 10)
	go func() {
		for range ch {
			mc.Inc()
			time.Sleep(time.Second)
		}
		fmt.Printf("Counter: %d\n", mc.Load())
	}()

	superFetch("http://100.83.243.106:9650/ext/bc/P", 0, 10, ch)
	t.Fatal()
}

func TestUpdateRewards(t *testing.T) {
	ctx := context.Background()
	dbFile, _ := db.OpenDB("../../data/pandasia-dev.db")
	UpdateRewards(ctx, dbFile)
	t.Fatal()
}

func superFetch(uri string, startIdx int, numToFetch int, ch chan *colly.Response) error {
	threads := runtime.NumCPU()
	c := colly.NewCollector(colly.AllowURLRevisit())
	c.OnResponse(func(r *colly.Response) {
		ch <- r
	})
	q, _ := queue.New(
		threads, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: numToFetch},
	)
	u, _ := url.Parse(uri)

	for i := startIdx; i < startIdx+numToFetch; i++ {
		body := strings.NewReader(fmt.Sprintf(`{"id":0,"jsonrpc":"2.0","method":"platform.getBlockByHeight","params":{"height":"%d","encoding":"hex"}}`, i))
		r := &colly.Request{
			URL:     u,
			Headers: &http.Header{"Content-Type": {"application/json"}},
			Method:  "POST",
			Body:    body,
		}
		q.AddRequest(r)
	}
	err := q.Run(c)
	close(ch)
	return err
}

type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (c *MutexCounter) Inc() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.counter++
}

func (c *MutexCounter) Load() int64 {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.counter
}
