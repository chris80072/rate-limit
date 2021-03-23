package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetResponseWithoutRecord(t *testing.T) {

	response := &Response{Ip: "0.0.0.1"}
	var count int64 = 0
	var limit int64 = 60
	expected := "1"

	SetResponse(count, limit, response)

	assert.Equal(t, expected, response.Request, "they should be equal")
}

func TestSetResponseWithRecordReachLimit(t *testing.T) {

	response := &Response{Ip: "0.0.0.1"}
	var count int64 = 59
	var limit int64 = 60
	expected := "60"

	SetResponse(count, limit, response)

	assert.Equal(t, expected, response.Request, "they should be equal")
}

func TestSetResponseWithRecordOverLimit(t *testing.T) {

	response := &Response{Ip: "0.0.0.1"}
	var count int64 = 60
	var limit int64 = 60
	expected := "Error"

	SetResponse(count, limit, response)

	assert.Equal(t, expected, response.Request, "they should be equal")
}
