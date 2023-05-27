package db

import (
	"fmt"
	"testing"
	"zakirullin/dumpbot/internal/sched"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/require"
)

func TestAddSchedule(t *testing.T) {
	r := require.New(t)

	redis, err := miniredis.Run()
	if err != nil {
		panic(fmt.Sprintf("Can't create Redis: %s\n", err))
	}
	defer func() {
		redis.Close()
	}()

	db := NewDB(redis)
	err = db.AddToSchedule(-1, "filename", 1, "cron")
	r.Nil(err)

	val, _ := redis.Get("schedule:-1")
	r.Equal(`{"filename":{"RunAt":1,"Cron":"cron","Cmd":"move"}}`, val)
}

func TestGetSchedule(t *testing.T) {
	r := require.New(t)

	redis, err := miniredis.Run()
	if err != nil {
		panic(fmt.Sprintf("Can't create Redis: %s\n", err))
	}
	defer func() {
		redis.Close()
	}()

	db := NewDB(redis)
	err = db.AddToSchedule(-1, "filename", 1, "cron")
	r.Nil(err)

	sc, err := db.Schedule(-1)
	r.Nil(err)

	r.Equal(map[string]sched.Cron{
		"filename": sched.NewCron(1, "cron"),
	}, sc)
}

func TestDelFromSchedule(t *testing.T) {
	r := require.New(t)

	redis, err := miniredis.Run()
	if err != nil {
		panic(fmt.Sprintf("Can't create Redis: %s\n", err))
	}
	defer func() {
		redis.Close()
	}()

	db := NewDB(redis)
	err = db.AddToSchedule(-1, "filename", 1, "cron")
	r.Nil(err)

	err = db.DelFromSchedule(-1, "filename")
	r.Nil(err)

	sc, err := db.Schedule(-1)

	r.Empty(sc)
}
