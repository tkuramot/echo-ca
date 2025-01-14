package task

import (
	errDomain "github/tkuramot/echo-practice/internal/domain/error"
	"github/tkuramot/echo-practice/pkg/ulid"
	"unicode/utf8"
)

type Status string

const (
	Backlog    Status = "backlog"
	InProgress Status = "in_progress"
	Done       Status = "done"
	Cancelled  Status = "canceled"
)

var (
	ErrInvalidStatus = errDomain.NewError(
		errDomain.InvalidArgument,
		"無効なステータスです",
	)
)

type Task struct {
	id          string
	title       string
	description string
	status      Status
}

const (
	titleLengthMin       = 1
	titleLengthMax       = 255
	descriptionLengthMax = 1000
)

func Reconstruct(id, title, description string, status Status) (*Task, error) {
	return newTask(id, title, description, status)
}

func NewTask(title, description string) (*Task, error) {
	return newTask(ulid.NewULID(), title, description, Backlog)
}

func newTask(id, title, description string, status Status) (*Task, error) {
	if id == "" {
		return nil, errDomain.NewError(
			errDomain.InvalidArgument,
			"IDは必須です",
		)
	}

	if utf8.RuneCountInString(title) < titleLengthMin || utf8.RuneCountInString(title) > titleLengthMax {
		return nil, errDomain.NewError(
			errDomain.InvalidArgument,
			"タイトルは1文字以上、255文字以下で入力してください",
		)
	}

	if utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errDomain.NewError(
			errDomain.InvalidArgument,
			"説明は1000文字以下で入力してください",
		)
	}

	if !IsValidStatus(status) {
		return nil, ErrInvalidStatus
	}

	return &Task{
		id:          id,
		title:       title,
		description: description,
		status:      status,
	}, nil
}

func (t *Task) ID() string {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) Status() Status {
	return t.status
}

func IsValidStatus(status Status) bool {
	return status == Backlog || status == InProgress || status == Done || status == Cancelled
}
