// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColUInt64 represents UInt64 column.
type ColUInt64 []uint64

// Compile-time assertions for ColUInt64.
var (
	_ ColInput  = ColUInt64{}
	_ ColResult = (*ColUInt64)(nil)
	_ Column    = (*ColUInt64)(nil)
)

// Rows returns count of rows in column.
func (c ColUInt64) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColUInt64) Reset() {
	*c = (*c)[:0]
}

// Type returns ColumnType of UInt64.
func (ColUInt64) Type() ColumnType {
	return ColumnTypeUInt64
}

// Row returns i-th row of column.
func (c ColUInt64) Row(i int) uint64 {
	return c[i]
}

// Append uint64 to column.
func (c *ColUInt64) Append(v uint64) {
	*c = append(*c, v)
}

// Append uint64 slice to column.
func (c *ColUInt64) AppendArr(vs []uint64) {
	*c = append(*c, vs...)
}

// LowCardinality returns LowCardinality for UInt64.
func (c *ColUInt64) LowCardinality() *ColLowCardinality[uint64] {
	return &ColLowCardinality[uint64]{
		index: c,
	}
}

// Array is helper that creates Array of uint64.
func (c *ColUInt64) Array() *ColArr[uint64] {
	return &ColArr[uint64]{
		Data: c,
	}
}

// Nullable is helper that creates Nullable(uint64).
func (c *ColUInt64) Nullable() *ColNullable[uint64] {
	return &ColNullable[uint64]{
		Values: c,
	}
}

// NewArrUInt64 returns new Array(UInt64).
func NewArrUInt64() *ColArr[uint64] {
	return &ColArr[uint64]{
		Data: new(ColUInt64),
	}
}
