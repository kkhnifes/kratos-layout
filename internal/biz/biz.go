// Package biz holds domain models, usecases, and repo interfaces.
package biz

// New builds this layer's usecases from their repo dependencies.
func New(todoRepo TodoRepo) *TodoUsecase {
	return NewTodoUsecase(todoRepo)
}
