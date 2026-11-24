//go:build !windows && !nogui

package gui

// askInstallFolder: only Windows installs a download itself.
func askInstallFolder() func(string) { return nil }
