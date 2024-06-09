package main

import "log/slog"

type MethodExecutor struct{}

func (m *MethodExecutor) AdminLogin() {
	slog.Info("called admin login")
}

func (m *MethodExecutor) AdminLoginV2() {
	slog.Info("called admin login v2")
}
