package logger

import "io"

type Sugar struct {
}

func NewSugar(writer io.Writer) (*Sugar, error) {
	return nil, nil
}

func (*Sugar) Debug(msg string, v ...any) {}
func (*Sugar) Info(msg string, v ...any) {}
func (*Sugar) Warn(msg string, v ...any) {}
func (*Sugar) Panic(msg string, v ...any) {}
func (*Sugar) Error(msg string, v ...any) {}
func (*Sugar) Fatal(msg string, v ...any) {}
