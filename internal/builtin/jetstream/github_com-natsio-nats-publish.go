// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

// Code generated by 'yaegi extract github.com/PaesslerAG/jsonpath'. DO NOT EDIT.

package jetstream

import (
	"github.com/nats-io/nats.go/jetstream"
	"reflect"
)

func init() {
	Symbols["github.com/nats-io/nats.go/jetstream/publish"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Publish": 		reflect.ValueOf(publish.Publish),
		"PublishAsync": 		reflect.ValueOf(publish.PublishAsync),
		"PublishAsyncComplete": 		reflect.ValueOf(publish.PublishAsyncComplete),
	}
}
