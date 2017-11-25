package store

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"bytes"

	"github.com/boltdb/bolt"
	"github.com/umputun/rlb-stats/app/parse"
)

var bucketCandles = []byte("stats")
var bucketLogs = []byte("raw_logs")

// Bolt implements store.Engine with boltdb
type Bolt struct {
	db *bolt.DB
}

// NewBolt makes persistent boltdb based store
func NewBolt(dbFile string, collectDuration time.Duration) (*Bolt, error) {
	log.Printf("[INFO] bolt (persitent) store, %s", dbFile)
	result := Bolt{}
	db, err := bolt.Open(dbFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
	db.Update(func(tx *bolt.Tx) error {
		_, e := tx.CreateBucketIfNotExists(bucketCandles)
		return e
	})
	db.Update(func(tx *bolt.Tx) error {
		_, e := tx.CreateBucketIfNotExists(bucketLogs)
		return e
	})
	result.db = db
	result.activateCollector(collectDuration)
	return &result, err
}

// Save with ts-ip as a key. ts prefix for bolt range query
func (s *Bolt) Save(entry *parse.LogEntry) (err error) {
	key := fmt.Sprintf("%d-%s", entry.Date.Unix(), entry.SourceIP)
	total := 0
	err = s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketLogs)
		total = b.Stats().KeyN
		jdata, jerr := json.Marshal(entry)
		if jerr != nil {
			return err
		}
		return b.Put([]byte(key), jdata)
	})

	log.Printf("[DEBUG] saved, time=%v, total=%d", entry.Date.Unix(), total+1)
	return err
}

// Load LogEntries by period
func (s *Bolt) loadLogEntry(periodStart, periodEnd time.Time) (result []*parse.LogEntry, err error) {

	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketLogs)
		c := b.Cursor()

		// ip just a place holder to make keys sorted properly by ts prefix
		min := []byte(fmt.Sprintf("%d-127.0.0.1", periodStart.Unix()))
		max := []byte(fmt.Sprintf("%d-127.0.0.1", periodEnd.Unix()))

		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			log.Printf("[DEBUG] found entry %s", string(k))
			entry := &parse.LogEntry{}
			json.Unmarshal(b.Get(k), entry)
			result = append(result, entry)
			_ = v
		}

		return nil
	})

	return
}

// Load Candles by period
func (s *Bolt) Load(periodStart, periodEnd time.Time) (result []*Candle, err error) {
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketCandles)
		c := b.Cursor()

		min := []byte(fmt.Sprintf("%d", periodStart.Unix()))
		max := []byte(fmt.Sprintf("%d", periodEnd.Unix()))

		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			log.Printf("[DEBUG] found candle %s", string(k))
			entry := &Candle{}
			json.Unmarshal(b.Get(k), entry)
			result = append(result, entry)
			_ = v
		}

		return nil
	})

	return
}

// TODO: write logEntry->candle function
// TODO: dedupe same ips in candle (by which rule?)

// activateCollector runs periodic cleanups to aggregate data into candles
// detection based on ts (unix time) prefix of the key.
func (s *Bolt) activateCollector(every time.Duration) {
	log.Printf("[INFO] collecter activated, every %v", every)

	ticker := time.NewTicker(every)
	go func() {
		for range ticker.C {

			// TODO: collect data for previous period and group into candles

		}
	}()
}
