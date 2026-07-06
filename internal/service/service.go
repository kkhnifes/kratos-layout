// Package service holds transport adapters that convert DTO ↔ DO.
package service

import "github.com/kkhnifes/kratos-layout/internal/biz"

// New builds this layer's services from their usecase dependencies.
func New(todoUC *biz.TodoUsecase) *TodoService {
	return NewTodoService(todoUC)
}
