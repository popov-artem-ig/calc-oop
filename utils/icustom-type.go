package utils

// ICustomType - Общий интерфейс для реализованных типов данных
type ICustomType interface {
	IsEmpty() bool
	ToString() (string, error)
}