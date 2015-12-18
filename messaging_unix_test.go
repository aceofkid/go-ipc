// Copyright 2015 Aleksandr Demakin. All rights reserved.

// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package ipc

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

const (
	testMqName = "go-ipc.testmq"
)

func TestCreateMq(t *testing.T) {
	_, err := NewMessageQueue(testMqName, O_OPEN_OR_CREATE|O_READWRITE, 0666, 4)
	if assert.NoError(t, err) {
		assert.NoError(t, DestroyMessageQueue(testMqName))
	}
}

func TestCreateMqExcl(t *testing.T) {
	_, err := NewMessageQueue(testMqName, O_OPEN_OR_CREATE|O_READWRITE, 0666, 4)
	if !assert.NoError(t, err) {
		return
	}
	_, err = NewMessageQueue(testMqName, O_CREATE_ONLY|O_READWRITE, 0666, 4)
	assert.Error(t, err)
}

func TestCreateMqOpenOnly(t *testing.T) {
	_, err := NewMessageQueue(testMqName, O_OPEN_OR_CREATE|O_READWRITE, 0666, 4)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.NoError(t, DestroyMessageQueue(testMqName)) {
		return
	}
	_, err = NewMessageQueue(testMqName, O_OPEN_ONLY|O_READWRITE, 0666, 4)
	assert.Error(t, err)
}

func TestMqSendIntSameProcess(t *testing.T) {
	var message int = 1122
	mq, err := NewMessageQueue(testMqName, O_OPEN_OR_CREATE|O_READWRITE, 0666, int(unsafe.Sizeof(message)))
	if !assert.NoError(t, err) {
		return
	}
	defer mq.Destroy()
	go func() {
		assert.NoError(t, mq.Send(message, 1))
	}()
	var received int
	var prio int
	if !assert.NoError(t, mq.Receive(&received, &prio)) {
		return
	}
	assert.Equal(t, received, message)
	assert.Equal(t, 1, prio)
}

func TestMqSendSliceSameProcess(t *testing.T) {
	mq, err := NewMessageQueue(testMqName, O_OPEN_OR_CREATE|O_READWRITE, 0666, 1024)
	if !assert.NoError(t, err) {
		return
	}
	defer mq.Destroy()
	message := make([]byte, 16)
	for i, _ := range message {
		message[i] = byte(i)
	}
	go func() {
		assert.NoError(t, mq.Send(message, 1))
	}()
	received := make([]byte, 16)
	if !assert.NoError(t, mq.Receive(received, nil)) {
		return
	}
	assert.Equal(t, received, message)
}
