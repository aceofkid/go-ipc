// Copyright 2016 Aleksandr Demakin. All rights reserved.

// +build darwin freebsd

package main

import (
	"fmt"
	"os"

	"github.com/aceofkid/go-ipc/mq"
)

func createMqWithType(name string, perm os.FileMode, typ, opt string) (mq.Messenger, error) {
	switch typ {
	case "default":
		return mq.New(name, os.O_RDWR, perm)
	case "fast":
		mqSize, msgSize := mq.DefaultFastMqMaxSize, mq.DefaultFastMqMessageSize
		if first, second, err := parseTwoInts(opt); err == nil {
			mqSize, msgSize = first, second
		}
		return mq.CreateFastMq(name, 0, perm, mqSize, msgSize)
	case "sysv":
		return mq.CreateSystemVMessageQueue(name, 0, perm)
	default:
		return nil, fmt.Errorf("unknown mq type %q", typ)
	}
}

func openMqWithType(name string, flags int, typ string) (mq.Messenger, error) {
	switch typ {
	case "default":
		return mq.Open(name, flags)
	case "fast":
		return mq.OpenFastMq(name, flags)
	case "sysv":
		return mq.OpenSystemVMessageQueue(name, flags)
	default:
		return nil, fmt.Errorf("unknown mq type %q", typ)
	}
}

func destroyMqWithType(name, typ string) error {
	switch typ {
	case "default":
		return mq.Destroy(name)
	case "fast":
		return mq.DestroyFastMq(name)
	case "sysv":
		return mq.DestroySystemVMessageQueue(name)
	default:
		return fmt.Errorf("unknown mq type %q", typ)
	}
}

func notifywait(name string, timeout int, typ string) error {
	return fmt.Errorf("notifywait is not supported on current platform")
}
