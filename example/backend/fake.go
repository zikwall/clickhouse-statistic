package main

import "math/rand"

var events = []string{
	"clicked", "touched", "cancelled", "visited", "viewed",
}

var apps = []string{
	"android-official-client", "ios-official-client", "web-client", "custom-app",
}

var hosts = []string{
	"msk-host-9720", "data-x-3021", "france-bridge-007", "::1",
}

func faker(s []string) string {
	randomIndex := rand.Intn(len(s))
	return s[randomIndex]
}

func user() uint32 {
	return uint32(rand.Intn(30-10) + 10)
}
