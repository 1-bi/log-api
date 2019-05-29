package logapi

import "time"

type embbedLoggerBean struct {

	// define name array , value is field type
	fieldProps map[string]Field
}

func (myself *embbedLoggerBean) convertToFields() []Field {

	var fields = make([]Field, 0)
	for _, vf := range myself.fieldProps {
		fields = append(fields, vf)
	}

	return fields
}

func (myself *embbedLoggerBean) LogBinary(key string, value []byte) {
	myself.fieldProps[key] = Binary(key, value)
}

func (myself *embbedLoggerBean) LogByteString(key string, values []byte) {
	myself.fieldProps[key] = ByteString(key, values)
}

func (myself *embbedLoggerBean) LogByteStringArray(key string, values [][]byte) {
	//myself.fieldProps[key] = zap.ByteStrings(key, values)
}

func (myself *embbedLoggerBean) LogString(key string, value string) {
	myself.fieldProps[key] = String(key, value)
}

func (myself *embbedLoggerBean) LogStringArray(key string, values []string) {
	//myself.fieldProps[key] = Strings(key, values)
}

func (myself *embbedLoggerBean) LogBool(key string, value bool) {
	myself.fieldProps[key] = Bool(key, value)
}

func (myself *embbedLoggerBean) LogBoolArray(key string, values []bool) {
	//myself.fieldProps[key] = zap.Bools(key, values)
}

// --- int

func (myself *embbedLoggerBean) LogInt8(key string, value int8) {
	myself.fieldProps[key] = Int8(key, value)
}

func (myself *embbedLoggerBean) LogInt8Array(key string, nums []int8) {
	//myself.fieldProps[key] = zap.Int8s(key, nums)
}

func (myself *embbedLoggerBean) LogInt(key string, value int) {
	myself.fieldProps[key] = Int(key, value)
}

func (myself *embbedLoggerBean) LogIntArray(key string, nums []int) {
	//myself.fieldProps[key] = zap.Ints(key, nums)
}

func (myself *embbedLoggerBean) LogInt16(key string, value int16) {
	myself.fieldProps[key] = Int16(key, value)
}

func (myself *embbedLoggerBean) LogInt16Array(key string, nums []int16) {
	//myself.fieldProps[key] = zap.Int16s(key, nums)
}

func (myself *embbedLoggerBean) LogInt32(key string, value int32) {
	myself.fieldProps[key] = Int32(key, value)
}

func (myself *embbedLoggerBean) LogInt32Array(key string, nums []int32) {
	//myself.fieldProps[key] = zap.Int32s(key, nums)
}

func (myself *embbedLoggerBean) LogInt64(key string, value int64) {
	myself.fieldProps[key] = Int64(key, value)
}

func (myself *embbedLoggerBean) LogInt64Array(key string, nums []int64) {
	//myself.fieldProps[key] = zap.Int64s(key, nums)
}

// --- uint
func (myself *embbedLoggerBean) LogUint8(key string, value uint8) {
	myself.fieldProps[key] = Uint8(key, value)
}

func (myself *embbedLoggerBean) LogUint8Array(key string, nums []uint8) {
	//myself.fieldProps[key] = zap.Uint8s(key, nums)
}

func (myself *embbedLoggerBean) LogUint(key string, value uint) {
	myself.fieldProps[key] = Uint(key, value)
}

func (myself *embbedLoggerBean) LogUintArray(key string, nums []uint) {
	//myself.fieldProps[key] = zap.Uints(key, nums)
}

func (myself *embbedLoggerBean) LogUint16(key string, value uint16) {
	myself.fieldProps[key] = Uint16(key, value)
}

func (myself *embbedLoggerBean) LogUint16Array(key string, nums []uint16) {
	//myself.fieldProps[key] = zap.Uint16s(key, nums)
}

func (myself *embbedLoggerBean) LogUint32(key string, value uint32) {
	myself.fieldProps[key] = Uint32(key, value)
}

func (myself *embbedLoggerBean) LogUint32Array(key string, nums []uint32) {
	//myself.fieldProps[key] = zap.Uint32s(key, nums)
}

// --- float
func (myself *embbedLoggerBean) LogFloat32(key string, value float32) {
	myself.fieldProps[key] = Float32(key, value)
}

func (myself *embbedLoggerBean) LogFloat32Array(key string, values []float32) {
	//myself.fieldProps[key] = zap.Float32s(key, values)
}

func (myself *embbedLoggerBean) LogFloat64(key string, value float64) {
	myself.fieldProps[key] = Float64(key, value)
}

func (myself *embbedLoggerBean) LogFloat64Array(key string, values []float64) {
	//myself.fieldProps[key] = zap.Float64s(key, values)
}

// --- time or duration

func (myself *embbedLoggerBean) LogDuration(key string, value time.Duration) {
	myself.fieldProps[key] = Duration(key, value)
}

func (myself *embbedLoggerBean) LogDurationArray(key string, values []time.Duration) {
	//myself.fieldProps[key] = zap.Durations(key, values)
}

func (myself *embbedLoggerBean) LogTime(key string, value time.Time) {
	myself.fieldProps[key] = Time(key, value)
}

func (myself *embbedLoggerBean) LogTimeArray(key string, values []time.Time) {
	//myself.fieldProps[key] = zap.Times(key, values)
}
