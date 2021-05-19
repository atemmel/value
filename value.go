package value

type Value interface {
	AsInt() *int
	AsUint() *uint
	AsBool() *bool
	AsString() *string
	AsF32() *float32
	AsF64() *float64
	AsByte() *byte

	AsIntArray() []int
	AsUintArray() []uint
	AsBoolArray() []bool
	AsStringArray() []string
	AsF32Array() []float32
	AsF64Array() []float64
	AsByteArray() []byte

	AsValueArray() []Value
}

func Int(value int) Value {
	return &intValue{value}
}

func Uint(value uint) Value {
	return &uintValue{value}
}

func Bool(value bool) Value {
	return &boolValue{value}
}

func String(value string) Value {
	return &stringValue{value}
}

func F32(value float32) Value {
	return &f32Value{value}
}

func F64(value float64) Value {
	return &f64Value{value}
}

func Byte(value byte) Value {
	return &byteValue{value}
}

func IntArray(value []int) Value {
	return &intArrayValue{value}
}

func UintArray(value []uint) Value {
	return &uintArrayValue{value}
}

func BoolArray(value []bool) Value {
	return &boolArrayValue{value}
}

func StringArray(value []string) Value {
	return &stringArrayValue{value}
}

func F32Array(value []float32) Value {
	return &f32ArrayValue{value}
}

func F64Array(value []float64) Value {
	return &f64ArrayValue{value}
}

func ByteArray(value []byte) Value {
	return &byteArrayValue{value}
}

func Array(value []Value) Value {
	return &arrayValue{value}
}

func ValueTable() map[string]Value {
	return make(map[string]Value)
}

type intValue struct {
	value int
}

func (v *intValue) AsInt() *int {
	return &v.value
}

func (v *intValue) AsUint() *uint {
	return nil
}

func (v *intValue) AsBool() *bool {
	return nil
}

func (v *intValue) AsString() *string {
	return nil
}

func (v *intValue) AsF32() *float32 {
	return nil
}

func (v *intValue) AsF64() *float64 {
	return nil
}

func (v *intValue) AsByte() *byte {
	return nil
}

func (v *intValue) AsIntArray() []int {
	return nil
}

func (v *intValue) AsUintArray() []uint {
	return nil
}

func (v *intValue) AsBoolArray() []bool {
	return nil
}

func (v *intValue) AsStringArray() []string {
	return nil
}

func (v *intValue) AsF32Array() []float32 {
	return nil
}

func (v *intValue) AsF64Array() []float64 {
	return nil
}

func (v *intValue) AsByteArray() []byte {
	return nil
}

func (v *intValue) AsValueArray() []Value {
	return nil
}

type uintValue struct {
	value uint
}

func (v *uintValue) AsInt() *int {
	return nil
}

func (v *uintValue) AsUint() *uint {
	return &v.value
}

func (v *uintValue) AsBool() *bool {
	return nil
}

func (v *uintValue) AsString() *string {
	return nil
}

func (v *uintValue) AsF32() *float32 {
	return nil
}

func (v *uintValue) AsF64() *float64 {
	return nil
}

func (v *uintValue) AsByte() *byte {
	return nil
}

func (v *uintValue) AsIntArray() []int {
	return nil
}

func (v *uintValue) AsUintArray() []uint {
	return nil
}

func (v *uintValue) AsBoolArray() []bool {
	return nil
}

func (v *uintValue) AsStringArray() []string {
	return nil
}

func (v *uintValue) AsF32Array() []float32 {
	return nil
}

func (v *uintValue) AsF64Array() []float64 {
	return nil
}

func (v *uintValue) AsByteArray() []byte {
	return nil
}

func (v *uintValue) AsValueArray() []Value {
	return nil
}

type boolValue struct {
	value bool
}

func (v *boolValue) AsInt() *int {
	return nil
}

func (v *boolValue) AsUint() *uint {
	return nil
}

func (v *boolValue) AsBool() *bool {
	return &v.value
}

func (v *boolValue) AsString() *string {
	return nil
}

func (v *boolValue) AsF32() *float32 {
	return nil
}

func (v *boolValue) AsF64() *float64 {
	return nil
}

func (v *boolValue) AsByte() *byte {
	return nil
}

func (v *boolValue) AsIntArray() []int {
	return nil
}

func (v *boolValue) AsUintArray() []uint {
	return nil
}

func (v *boolValue) AsBoolArray() []bool {
	return nil
}

func (v *boolValue) AsStringArray() []string {
	return nil
}

func (v *boolValue) AsF32Array() []float32 {
	return nil
}

func (v *boolValue) AsF64Array() []float64 {
	return nil
}

func (v *boolValue) AsByteArray() []byte {
	return nil
}

func (v *boolValue) AsValueArray() []Value {
	return nil
}

type stringValue struct {
	value string
}

func (v *stringValue) AsInt() *int {
	return nil
}

func (v *stringValue) AsUint() *uint {
	return nil
}

func (v *stringValue) AsBool() *bool {
	return nil
}

func (v *stringValue) AsString() *string {
	return &v.value
}

func (v *stringValue) AsF32() *float32 {
	return nil
}

func (v *stringValue) AsF64() *float64 {
	return nil
}

func (v *stringValue) AsByte() *byte {
	return nil
}

func (v *stringValue) AsIntArray() []int {
	return nil
}

func (v *stringValue) AsUintArray() []uint {
	return nil
}

func (v *stringValue) AsBoolArray() []bool {
	return nil
}

func (v *stringValue) AsStringArray() []string {
	return nil
}

func (v *stringValue) AsF32Array() []float32 {
	return nil
}

func (v *stringValue) AsF64Array() []float64 {
	return nil
}

func (v *stringValue) AsByteArray() []byte {
	return nil
}

func (v *stringValue) AsValueArray() []Value {
	return nil
}

type f32Value struct {
	value float32
}

func (v *f32Value) AsInt() *int {
	return nil
}

func (v *f32Value) AsUint() *uint {
	return nil
}

func (v *f32Value) AsBool() *bool {
	return nil
}

func (v *f32Value) AsString() *string {
	return nil
}

func (v *f32Value) AsF32() *float32 {
	return &v.value
}

func (v *f32Value) AsF64() *float64 {
	return nil
}

func (v *f32Value) AsByte() *byte {
	return nil
}

func (v *f32Value) AsIntArray() []int {
	return nil
}

func (v *f32Value) AsUintArray() []uint {
	return nil
}

func (v *f32Value) AsBoolArray() []bool {
	return nil
}

func (v *f32Value) AsStringArray() []string {
	return nil
}

func (v *f32Value) AsF32Array() []float32 {
	return nil
}

func (v *f32Value) AsF64Array() []float64 {
	return nil
}

func (v *f32Value) AsByteArray() []byte {
	return nil
}

func (v *f32Value) AsValueArray() []Value {
	return nil
}

type f64Value struct {
	value float64
}

func (v *f64Value) AsInt() *int {
	return nil
}

func (v *f64Value) AsUint() *uint {
	return nil
}

func (v *f64Value) AsBool() *bool {
	return nil
}

func (v *f64Value) AsString() *string {
	return nil
}

func (v *f64Value) AsF32() *float32 {
	return nil
}

func (v *f64Value) AsF64() *float64 {
	return &v.value
}

func (v *f64Value) AsByte() *byte {
	return nil
}

func (v *f64Value) AsIntArray() []int {
	return nil
}

func (v *f64Value) AsUintArray() []uint {
	return nil
}

func (v *f64Value) AsBoolArray() []bool {
	return nil
}

func (v *f64Value) AsStringArray() []string {
	return nil
}

func (v *f64Value) AsF32Array() []float32 {
	return nil
}

func (v *f64Value) AsF64Array() []float64 {
	return nil
}

func (v *f64Value) AsByteArray() []byte {
	return nil
}

func (v *f64Value) AsValueArray() []Value {
	return nil
}

type byteValue struct {
	value byte
}

func (v *byteValue) AsInt() *int {
	return nil
}

func (v *byteValue) AsUint() *uint {
	return nil
}

func (v *byteValue) AsBool() *bool {
	return nil
}

func (v *byteValue) AsString() *string {
	return nil
}

func (v *byteValue) AsF32() *float32 {
	return nil
}

func (v *byteValue) AsF64() *float64 {
	return nil
}

func (v *byteValue) AsByte() *byte {
	return &v.value
}

func (v *byteValue) AsIntArray() []int {
	return nil
}

func (v *byteValue) AsUintArray() []uint {
	return nil
}

func (v *byteValue) AsBoolArray() []bool {
	return nil
}

func (v *byteValue) AsStringArray() []string {
	return nil
}

func (v *byteValue) AsF32Array() []float32 {
	return nil
}

func (v *byteValue) AsF64Array() []float64 {
	return nil
}

func (v *byteValue) AsByteArray() []byte {
	return nil
}

func (v *byteValue) AsValueArray() []Value {
	return nil
}

type intArrayValue struct {
	value []int
}

func (v *intArrayValue) AsInt() *int {
	return nil
}

func (v *intArrayValue) AsUint() *uint {
	return nil
}

func (v *intArrayValue) AsBool() *bool {
	return nil
}

func (v *intArrayValue) AsString() *string {
	return nil
}

func (v *intArrayValue) AsF32() *float32 {
	return nil
}

func (v *intArrayValue) AsF64() *float64 {
	return nil
}

func (v *intArrayValue) AsByte() *byte {
	return nil
}

func (v *intArrayValue) AsIntArray() []int {
	return v.value
}

func (v *intArrayValue) AsUintArray() []uint {
	return nil
}

func (v *intArrayValue) AsBoolArray() []bool {
	return nil
}

func (v *intArrayValue) AsStringArray() []string {
	return nil
}

func (v *intArrayValue) AsF32Array() []float32 {
	return nil
}

func (v *intArrayValue) AsF64Array() []float64 {
	return nil
}

func (v *intArrayValue) AsByteArray() []byte {
	return nil
}

func (v *intArrayValue) AsValueArray() []Value {
	return nil
}

type uintArrayValue struct {
	value []uint
}

func (v *uintArrayValue) AsInt() *int {
	return nil
}

func (v *uintArrayValue) AsUint() *uint {
	return nil
}

func (v *uintArrayValue) AsBool() *bool {
	return nil
}

func (v *uintArrayValue) AsString() *string {
	return nil
}

func (v *uintArrayValue) AsF32() *float32 {
	return nil
}

func (v *uintArrayValue) AsF64() *float64 {
	return nil
}

func (v *uintArrayValue) AsByte() *byte {
	return nil
}

func (v *uintArrayValue) AsIntArray() []int {
	return nil
}

func (v *uintArrayValue) AsUintArray() []uint {
	return v.value
}

func (v *uintArrayValue) AsBoolArray() []bool {
	return nil
}

func (v *uintArrayValue) AsStringArray() []string {
	return nil
}

func (v *uintArrayValue) AsF32Array() []float32 {
	return nil
}

func (v *uintArrayValue) AsF64Array() []float64 {
	return nil
}

func (v *uintArrayValue) AsByteArray() []byte {
	return nil
}

func (v *uintArrayValue) AsValueArray() []Value {
	return nil
}

type boolArrayValue struct {
	value []bool
}

func (v *boolArrayValue) AsInt() *int {
	return nil
}

func (v *boolArrayValue) AsUint() *uint {
	return nil
}

func (v *boolArrayValue) AsBool() *bool {
	return nil
}

func (v *boolArrayValue) AsString() *string {
	return nil
}

func (v *boolArrayValue) AsF32() *float32 {
	return nil
}

func (v *boolArrayValue) AsF64() *float64 {
	return nil
}

func (v *boolArrayValue) AsByte() *byte {
	return nil
}

func (v *boolArrayValue) AsIntArray() []int {
	return nil
}

func (v *boolArrayValue) AsUintArray() []uint {
	return nil
}

func (v *boolArrayValue) AsBoolArray() []bool {
	return v.value
}

func (v *boolArrayValue) AsStringArray() []string {
	return nil
}

func (v *boolArrayValue) AsF32Array() []float32 {
	return nil
}

func (v *boolArrayValue) AsF64Array() []float64 {
	return nil
}

func (v *boolArrayValue) AsByteArray() []byte {
	return nil
}

func (v *boolArrayValue) AsValueArray() []Value {
	return nil
}

type stringArrayValue struct {
	value []string
}

func (v *stringArrayValue) AsInt() *int {
	return nil
}

func (v *stringArrayValue) AsUint() *uint {
	return nil
}

func (v *stringArrayValue) AsBool() *bool {
	return nil
}

func (v *stringArrayValue) AsString() *string {
	return nil
}

func (v *stringArrayValue) AsF32() *float32 {
	return nil
}

func (v *stringArrayValue) AsF64() *float64 {
	return nil
}

func (v *stringArrayValue) AsByte() *byte {
	return nil
}

func (v *stringArrayValue) AsIntArray() []int {
	return nil
}

func (v *stringArrayValue) AsUintArray() []uint {
	return nil
}

func (v *stringArrayValue) AsBoolArray() []bool {
	return nil
}

func (v *stringArrayValue) AsStringArray() []string {
	return v.value
}

func (v *stringArrayValue) AsF32Array() []float32 {
	return nil
}

func (v *stringArrayValue) AsF64Array() []float64 {
	return nil
}

func (v *stringArrayValue) AsByteArray() []byte {
	return nil
}

func (v *stringArrayValue) AsValueArray() []Value {
	return nil
}

type f32ArrayValue struct {
	value []float32
}

func (v *f32ArrayValue) AsInt() *int {
	return nil
}

func (v *f32ArrayValue) AsUint() *uint {
	return nil
}

func (v *f32ArrayValue) AsBool() *bool {
	return nil
}

func (v *f32ArrayValue) AsString() *string {
	return nil
}

func (v *f32ArrayValue) AsF32() *float32 {
	return nil
}

func (v *f32ArrayValue) AsF64() *float64 {
	return nil
}

func (v *f32ArrayValue) AsByte() *byte {
	return nil
}

func (v *f32ArrayValue) AsIntArray() []int {
	return nil
}

func (v *f32ArrayValue) AsUintArray() []uint {
	return nil
}

func (v *f32ArrayValue) AsBoolArray() []bool {
	return nil
}

func (v *f32ArrayValue) AsStringArray() []string {
	return nil
}

func (v *f32ArrayValue) AsF32Array() []float32 {
	return v.value
}

func (v *f32ArrayValue) AsF64Array() []float64 {
	return nil
}

func (v *f32ArrayValue) AsByteArray() []byte {
	return nil
}

func (v *f32ArrayValue) AsValueArray() []Value {
	return nil
}

type f64ArrayValue struct {
	value []float64
}

func (v *f64ArrayValue) AsInt() *int {
	return nil
}

func (v *f64ArrayValue) AsUint() *uint {
	return nil
}

func (v *f64ArrayValue) AsBool() *bool {
	return nil
}

func (v *f64ArrayValue) AsString() *string {
	return nil
}

func (v *f64ArrayValue) AsF32() *float32 {
	return nil
}

func (v *f64ArrayValue) AsF64() *float64 {
	return nil
}

func (v *f64ArrayValue) AsByte() *byte {
	return nil
}

func (v *f64ArrayValue) AsIntArray() []int {
	return nil
}

func (v *f64ArrayValue) AsUintArray() []uint {
	return nil
}

func (v *f64ArrayValue) AsBoolArray() []bool {
	return nil
}

func (v *f64ArrayValue) AsStringArray() []string {
	return nil
}

func (v *f64ArrayValue) AsF32Array() []float32 {
	return nil
}

func (v *f64ArrayValue) AsF64Array() []float64 {
	return v.value
}

func (v *f64ArrayValue) AsByteArray() []byte {
	return nil
}

func (v *f64ArrayValue) AsValueArray() []Value {
	return nil
}

type byteArrayValue struct {
	value []byte
}

func (v *byteArrayValue) AsInt() *int {
	return nil
}

func (v *byteArrayValue) AsUint() *uint {
	return nil
}

func (v *byteArrayValue) AsBool() *bool {
	return nil
}

func (v *byteArrayValue) AsString() *string {
	return nil
}

func (v *byteArrayValue) AsF32() *float32 {
	return nil
}

func (v *byteArrayValue) AsF64() *float64 {
	return nil
}

func (v *byteArrayValue) AsByte() *byte {
	return nil
}

func (v *byteArrayValue) AsIntArray() []int {
	return nil
}

func (v *byteArrayValue) AsUintArray() []uint {
	return nil
}

func (v *byteArrayValue) AsBoolArray() []bool {
	return nil
}

func (v *byteArrayValue) AsStringArray() []string {
	return nil
}

func (v *byteArrayValue) AsF32Array() []float32 {
	return nil
}

func (v *byteArrayValue) AsF64Array() []float64 {
	return nil
}

func (v *byteArrayValue) AsByteArray() []byte {
	return v.value
}

func (v *byteArrayValue) AsValueArray() []Value {
	return nil
}

type arrayValue struct {
	value []Value
}

func (v *arrayValue) AsInt() *int {
	return nil
}

func (v *arrayValue) AsUint() *uint {
	return nil
}

func (v *arrayValue) AsBool() *bool {
	return nil
}

func (v *arrayValue) AsString() *string {
	return nil
}

func (v *arrayValue) AsF32() *float32 {
	return nil
}

func (v *arrayValue) AsF64() *float64 {
	return nil
}

func (v *arrayValue) AsByte() *byte {
	return nil
}

func (v *arrayValue) AsIntArray() []int {
	return nil
}

func (v *arrayValue) AsUintArray() []uint {
	return nil
}

func (v *arrayValue) AsBoolArray() []bool {
	return nil
}

func (v *arrayValue) AsStringArray() []string {
	return nil
}

func (v *arrayValue) AsF32Array() []float32 {
	return nil
}

func (v *arrayValue) AsF64Array() []float64 {
	return nil
}

func (v *arrayValue) AsByteArray() []byte {
	return nil
}

func (v *arrayValue) AsValueArray() []Value {
	return v.value
}
