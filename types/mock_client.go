// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/okex/okexchain-go-sdk/types (interfaces: BaseClient)

// Package types is a generated GoMock package.
package types

import (
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	gomock "github.com/golang/mock/gomock"
	go_amino "github.com/tendermint/go-amino"
	bytes "github.com/tendermint/tendermint/libs/bytes"
	types2 "github.com/tendermint/tendermint/rpc/core/types"
	reflect "reflect"
)

// MockBaseClient is a mock of BaseClient interface
type MockBaseClient struct {
	ctrl     *gomock.Controller
	recorder *MockBaseClientMockRecorder
}

// MockBaseClientMockRecorder is the mock recorder for MockBaseClient
type MockBaseClientMockRecorder struct {
	mock *MockBaseClient
}

// NewMockBaseClient creates a new mock instance
func NewMockBaseClient(ctrl *gomock.Controller) *MockBaseClient {
	mock := &MockBaseClient{ctrl: ctrl}
	mock.recorder = &MockBaseClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBaseClient) EXPECT() *MockBaseClientMockRecorder {
	return m.recorder
}

// Block mocks base method
func (m *MockBaseClient) Block(arg0 *int64) (*types2.ResultBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Block", arg0)
	ret0, _ := ret[0].(*types2.ResultBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Block indicates an expected call of Block
func (mr *MockBaseClientMockRecorder) Block(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Block", reflect.TypeOf((*MockBaseClient)(nil).Block), arg0)
}

// BlockResults mocks base method
func (m *MockBaseClient) BlockResults(arg0 *int64) (*types2.ResultBlockResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockResults", arg0)
	ret0, _ := ret[0].(*types2.ResultBlockResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockResults indicates an expected call of BlockResults
func (mr *MockBaseClientMockRecorder) BlockResults(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockResults", reflect.TypeOf((*MockBaseClient)(nil).BlockResults), arg0)
}

// Broadcast mocks base method
func (m *MockBaseClient) Broadcast(arg0 []byte, arg1 string) (types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", arg0, arg1)
	ret0, _ := ret[0].(types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockBaseClientMockRecorder) Broadcast(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockBaseClient)(nil).Broadcast), arg0, arg1)
}

// BuildAndBroadcast mocks base method
func (m *MockBaseClient) BuildAndBroadcast(arg0, arg1, arg2 string, arg3 []types.Msg, arg4, arg5 uint64) (types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildAndBroadcast", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildAndBroadcast indicates an expected call of BuildAndBroadcast
func (mr *MockBaseClientMockRecorder) BuildAndBroadcast(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildAndBroadcast", reflect.TypeOf((*MockBaseClient)(nil).BuildAndBroadcast), arg0, arg1, arg2, arg3, arg4, arg5)
}

// BuildStdTx mocks base method
func (m *MockBaseClient) BuildStdTx(arg0, arg1, arg2 string, arg3 []types.Msg, arg4, arg5 uint64) (types0.StdTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildStdTx", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(types0.StdTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildStdTx indicates an expected call of BuildStdTx
func (mr *MockBaseClientMockRecorder) BuildStdTx(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildStdTx", reflect.TypeOf((*MockBaseClient)(nil).BuildStdTx), arg0, arg1, arg2, arg3, arg4, arg5)
}

// BuildTxForSim mocks base method
func (m *MockBaseClient) BuildTxForSim(arg0 []types.Msg, arg1 string, arg2, arg3 uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildTxForSim", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildTxForSim indicates an expected call of BuildTxForSim
func (mr *MockBaseClientMockRecorder) BuildTxForSim(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildTxForSim", reflect.TypeOf((*MockBaseClient)(nil).BuildTxForSim), arg0, arg1, arg2, arg3)
}

// BuildUnsignedStdTxOffline mocks base method
func (m *MockBaseClient) BuildUnsignedStdTxOffline(arg0 []types.Msg, arg1 string) types0.StdTx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildUnsignedStdTxOffline", arg0, arg1)
	ret0, _ := ret[0].(types0.StdTx)
	return ret0
}

// BuildUnsignedStdTxOffline indicates an expected call of BuildUnsignedStdTxOffline
func (mr *MockBaseClientMockRecorder) BuildUnsignedStdTxOffline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildUnsignedStdTxOffline", reflect.TypeOf((*MockBaseClient)(nil).BuildUnsignedStdTxOffline), arg0, arg1)
}

// CalculateGas mocks base method
func (m *MockBaseClient) CalculateGas(arg0 []byte) (types0.StdFee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateGas", arg0)
	ret0, _ := ret[0].(types0.StdFee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateGas indicates an expected call of CalculateGas
func (mr *MockBaseClientMockRecorder) CalculateGas(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateGas", reflect.TypeOf((*MockBaseClient)(nil).CalculateGas), arg0)
}

// Commit mocks base method
func (m *MockBaseClient) Commit(arg0 *int64) (*types2.ResultCommit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(*types2.ResultCommit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit
func (mr *MockBaseClientMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockBaseClient)(nil).Commit), arg0)
}

// GetCodec mocks base method
func (m *MockBaseClient) GetCodec() *go_amino.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCodec")
	ret0, _ := ret[0].(*go_amino.Codec)
	return ret0
}

// GetCodec indicates an expected call of GetCodec
func (mr *MockBaseClientMockRecorder) GetCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodec", reflect.TypeOf((*MockBaseClient)(nil).GetCodec))
}

// GetConfig mocks base method
func (m *MockBaseClient) GetConfig() ClientConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(ClientConfig)
	return ret0
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockBaseClientMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockBaseClient)(nil).GetConfig))
}

// Query mocks base method
func (m *MockBaseClient) Query(arg0 string, arg1 bytes.HexBytes) ([]byte, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Query indicates an expected call of Query
func (mr *MockBaseClientMockRecorder) Query(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockBaseClient)(nil).Query), arg0, arg1)
}

// QueryStore mocks base method
func (m *MockBaseClient) QueryStore(arg0 bytes.HexBytes, arg1, arg2 string) ([]byte, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryStore", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryStore indicates an expected call of QueryStore
func (mr *MockBaseClientMockRecorder) QueryStore(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryStore", reflect.TypeOf((*MockBaseClient)(nil).QueryStore), arg0, arg1, arg2)
}

// Tx mocks base method
func (m *MockBaseClient) Tx(arg0 []byte, arg1 bool) (*types2.ResultTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx", arg0, arg1)
	ret0, _ := ret[0].(*types2.ResultTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tx indicates an expected call of Tx
func (mr *MockBaseClientMockRecorder) Tx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockBaseClient)(nil).Tx), arg0, arg1)
}

// TxSearch mocks base method
func (m *MockBaseClient) TxSearch(arg0 string, arg1 bool, arg2, arg3 int, arg4 string) (*types2.ResultTxSearch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxSearch", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types2.ResultTxSearch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxSearch indicates an expected call of TxSearch
func (mr *MockBaseClientMockRecorder) TxSearch(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxSearch", reflect.TypeOf((*MockBaseClient)(nil).TxSearch), arg0, arg1, arg2, arg3, arg4)
}

// Validators mocks base method
func (m *MockBaseClient) Validators(arg0 *int64, arg1, arg2 int) (*types2.ResultValidators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validators", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types2.ResultValidators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validators indicates an expected call of Validators
func (mr *MockBaseClientMockRecorder) Validators(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validators", reflect.TypeOf((*MockBaseClient)(nil).Validators), arg0, arg1, arg2)
}
