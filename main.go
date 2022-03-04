package main

import (
	"encoding/json"
	"fmt"
	"time"
)

var df = "2006-01-02"

type Date time.Time

// current local date
func NewDate() Date {
	return Date(time.Now())
}

// string to Date
func (d *Date) UnmarshalText(text []byte) error {
	t, err := time.Parse(df, string(text))
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// Date to string
func (d Date) MarshalText() (text []byte, err error) {
	s := time.Time(d).Format(df)
	return []byte(s), nil
}

func (d Date) String() string {
	return time.Time(d).Format(df)
}

func main() {
	t := time.Now()
	fmt.Println(t) // 2022-03-04 12:39:08.360021 +0800 CST m=+0.000235750

	d := Date(t)
	fmt.Println(d) // 2022-03-04

	b, _ := json.Marshal(d)
	s := string(b)
	fmt.Println(s) // "2022-03-04"

	json.Unmarshal([]byte(s), &d)
	fmt.Println(d) // 2022-03-04
}
