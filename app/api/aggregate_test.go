package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/umputun/rlb-stats/app/store"
)

var testCandles = []store.Candle{
	{Nodes: map[string]store.Info{
		"n6.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 6, MeanAnswerTime: time.Second * 6, MaxAnswerTime: time.Second * 6, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
		"all":            {Volume: 1, MinAnswerTime: time.Second * 6, MeanAnswerTime: time.Second * 6, MaxAnswerTime: time.Second * 6, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
	},
		StartMinute: time.Time{}},
	{Nodes: map[string]store.Info{
		"n6.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
		"all":            {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
	},
		StartMinute: time.Time{}.Add(time.Minute)},
	{Nodes: map[string]store.Info{
		"n7.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast562.mp3": 1}},
		"all":            {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast562.mp3": 1}},
	},
		StartMinute: time.Time{}.Add(time.Minute * 2)},
	{Nodes: map[string]store.Info{
		"n7.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast562.mp3": 1}},
		"all":            {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast562.mp3": 1}},
	},
		StartMinute: time.Time{}.Add(time.Minute * 3)},
	{Nodes: map[string]store.Info{
		"n6.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
		"all":            {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
	},
		StartMinute: time.Time{}.Add(time.Minute * 4)},
	{Nodes: map[string]store.Info{
		"n6.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
		"all":            {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
	},
		StartMinute: time.Time{}.Add(time.Minute * 5)},
	{Nodes: map[string]store.Info{
		"n6.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
		"all":            {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
	},
		StartMinute: time.Time{}.Add(time.Minute * 10)},
}

var resultCandles = map[int][]store.Candle{
	1: testCandles,
	2: {
		{Nodes: map[string]store.Info{
			"n6.radio-t.com": {Volume: 2, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * (3 + 6) / 2, MaxAnswerTime: time.Second * 6, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 2}},
			"all":            {Volume: 2, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * (3 + 6) / 2, MaxAnswerTime: time.Second * 6, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 2}},
		},
			StartMinute: time.Time{}},

		{Nodes: map[string]store.Info{
			"n7.radio-t.com": {Volume: 2, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast562.mp3": 2}},
			"all":            {Volume: 2, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast562.mp3": 2}},
		},
			StartMinute: time.Time{}.Add(time.Minute * 2)},

		{Nodes: map[string]store.Info{
			"n6.radio-t.com": {Volume: 2, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 2}},
			"all":            {Volume: 2, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 2}},
		},
			StartMinute: time.Time{}.Add(time.Minute * 4)},

		{Nodes: map[string]store.Info{
			"n6.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
			"all":            {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
		},
			StartMinute: time.Time{}.Add(time.Minute * 10)},
	},
	3: {
		{Nodes: map[string]store.Info{
			"n6.radio-t.com": {Volume: 2, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * (3 + 6) / 2, MaxAnswerTime: time.Second * 6, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 2}},
			"n7.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast562.mp3": 1}},
			"all":            {Volume: 3, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * (3 + 6 + 3) / 3, MaxAnswerTime: time.Second * 6, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 2, "/rtfiles/rt_podcast562.mp3": 1}},
		},
			StartMinute: time.Time{}},
		{Nodes: map[string]store.Info{
			"n6.radio-t.com": {Volume: 2, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 2}},
			"n7.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast562.mp3": 1}},
			"all":            {Volume: 3, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 2, "/rtfiles/rt_podcast562.mp3": 1}},
		},
			StartMinute: time.Time{}.Add(time.Minute * 3)},
		{Nodes: map[string]store.Info{
			"n6.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
			"all":            {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
		},
			StartMinute: time.Time{}.Add(time.Minute * 9)},
	},
	5: {
		{Nodes: map[string]store.Info{
			"n6.radio-t.com": {Volume: 3, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * (3*2 + 6) / 3, MaxAnswerTime: time.Second * 6, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 3}},
			"n7.radio-t.com": {Volume: 2, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast562.mp3": 2}},
			"all":            {Volume: 5, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * (3*2 + 6 + 3*2) / 5, MaxAnswerTime: time.Second * 6, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 3, "/rtfiles/rt_podcast562.mp3": 2}},
		},
			StartMinute: time.Time{}},
		{Nodes: map[string]store.Info{
			"n6.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
			"all":            {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
		},
			StartMinute: time.Time{}.Add(time.Minute * 5)},
		{Nodes: map[string]store.Info{
			"n6.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
			"all":            {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
		},
			StartMinute: time.Time{}.Add(time.Minute * 10)}},
	10: {
		{Nodes: map[string]store.Info{
			"n6.radio-t.com": {Volume: 4, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * (3*3 + 6) / 4, MaxAnswerTime: time.Second * 6, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 4}},
			"n7.radio-t.com": {Volume: 2, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast562.mp3": 2}},
			"all":            {Volume: 6, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * (3*3 + 6 + 3*2) / 6, MaxAnswerTime: time.Second * 6, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 4, "/rtfiles/rt_podcast562.mp3": 2}},
		},
			StartMinute: time.Time{}},
		{Nodes: map[string]store.Info{
			"n6.radio-t.com": {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
			"all":            {Volume: 1, MinAnswerTime: time.Second * 3, MeanAnswerTime: time.Second * 3, MaxAnswerTime: time.Second * 3, Files: map[string]int{"/rtfiles/rt_podcast561.mp3": 1}},
		},
			StartMinute: time.Time{}.Add(time.Minute * 10)}},
}

func TestAggregation(t *testing.T) {

	for i, result := range resultCandles {
		testSlice := aggregateCandles(testCandles, time.Duration(i)*time.Minute)
		assert.EqualValues(t, result, testSlice, "candle aggregate for %v minutes match with expected output", i)
	}
	// test less than 1 minute period which should have same output as 1 minute aggregation
	testSlice := aggregateCandles(testCandles, time.Nanosecond)
	assert.EqualValues(t, testCandles, testSlice, "candle aggregate for 1 nanosecond match with expected output")
}
