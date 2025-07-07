package entity

import (
	"encoding/json"
	"strconv"

	dttm "github.com/dmastr/task-manager-cli/internal/common/datetimeformattter"
)

type TaskID int

func IdFromString(id string) (t TaskID, err error) {
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		return t, err
	}
	return TaskID(idParsed), err
}


type Task struct {
	Id          TaskID        `json:"id"`
	Text        string        `json:"text"`
	IsCompleted bool          `json:"is_completed"`
	CreatedAt   dttm.DateTime `json:"created_at"`
	UpdatedAt   dttm.DateTime `json:"updated_at"`
}

func NewTask(text string) Task {
	return Task{
		Text:        text,
		IsCompleted: false,
		CreatedAt:   dttm.DateTimeNow(),
		UpdatedAt:   dttm.DateTimeNow(),
	}
}

func (t *Task) ToString() string {
	res, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(res)
}
