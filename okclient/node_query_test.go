package okclient

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestQueryABCIInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryABCIInfo()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryConsenusState(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryConsenusState()
	assertNotEqual(t, err, nil)
	fmt.Println(string(resp))
}

func TestQueryDumpConsenusState(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryDumpConsenusState()
	assertNotEqual(t, err, nil)
	fmt.Println(string(resp))
}

func TestQueryNetInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryNetInfo()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryGenesisFile(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryGenesisFile()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryHealthInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryHealthInfo()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryUnconfirmedTxsNum(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryUnconfirmedTxsNum(30)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}