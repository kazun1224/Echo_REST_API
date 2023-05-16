package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// インターフェース(型)
type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

// 構造対 classみたいな
type taskValidator struct{}

// コンストラクタ
func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

// メゾット
func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task, validation.Field(&task.Title, validation.Required.Error("title is required"), validation.RuneLength(1, 10).Error("limited max 10 char")))
}
