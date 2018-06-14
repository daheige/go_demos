package time_test

import (
	"testing"
	"fmt"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	time := time.Time{}

	if now.After(time) {
		fmt.Println("yes")
	}
}

func TestGetTime(t *testing.T) {
	d1, _ := time.ParseDuration("-11h")

	t1 := time.Now()
	t2 := time.Now().Add(d1)
	d2 := t2.Sub(t1)
	fmt.Println(int(d2.Hours()))
}

func TestIsToday(t *testing.T) {
	d, _ := time.ParseDuration("-10h")
	fmt.Println(getTime(time.Now().Add(d)))
}

func getTime(t time.Time) string {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if t.After(today) {
		return "今天 " + t.Format("15:04")
	}
	return t.Format("2006-01-02 15:04")
}