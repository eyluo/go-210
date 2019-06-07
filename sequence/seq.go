//
// seq.go
//

package seq

import (
	"errors"
	"strings"
	"sync"
)

//
// Seq - the sequence type, which is an opaque slice of interface type.
//
type Seq struct {
	arr []interface{}
}

//
// Tuple - the tuple type, which is not natively supported by Go yet used by
// some sequence functions.
//
type Tuple struct {
	A interface{}
	B interface{}
}

//
// Nth - returns the Nth element in a sequence, or returns an error if out of
// bounds.
//
// Parameter s - the sequence to index into.
// Parameter n - the index to fetch from the sequence.
//
// Return interface{} - the nth value in the sequence.
// Return error - non-nil if n is out of bounds.
//
func Nth(s *Seq, n int) (interface{}, error) {
	if n < 0 || n >= len(s.arr) {
		return nil, errors.New("seq.Nth: out of bounds")
	}
	return s.arr[n], nil
}

//
// Length - returns the length of sequence s.
//
// Return int - the length of sequence s.
//
func Length(s *Seq) int {
	return len(s.arr)
}

//
// ToSlice - returns an index-preserving array representation of a sequence.
//
// Return []interface{} - an array representation of a sequence.
//
func ToSlice(s *Seq) []interface{} {
	result := make([]interface{}, len(s.arr))
	copy(result, s.arr)

	return result
}

//
// ToString - evaluates a string representation of s. Each element of s is
// converted to a string by an application of f.
//
// Parameter f - a function to convert each element of the sequence into a
// string.
//
// Return string - a string representation of the sequence.
//
func ToString(f func(interface{}) string, s *Seq) string {
	wg := new(sync.WaitGroup)
	strReps := make([]string, len(s.arr))
	for i, v := range s.arr {
		wg.Add(1)

		go func(idx int, val interface{}) {
			strReps[idx] = f(val)
			wg.Done()
		}(i, v)
	}

	wg.Wait()

	return strings.Join(strReps, ", ")
}

//
// Equal - returns whether or not s and t contain exactly the same elements, in
// the same order. Equality of element pairs is determined by f.
//
func Equal(f func(a, b interface{}) bool, s *Seq, t *Seq) bool {
	return true
}

//
// Empty -
//
func Empty() *Seq {
	return nil
}

//
// Singleton -
//
func Singleton(e interface{}) *Seq {
	return nil
}

//
// Tabulate -
//
func Tabulate(f func(i int) interface{}, n int) *Seq {
	return nil
}

//
// FromSlice -
//
func FromSlice(a []interface{}) *Seq {
	return nil
}

//
// Rev -
//
func Rev(s *Seq) *Seq {
	return nil
}

//
// Append -
//
func Append(s *Seq, t *Seq) *Seq {
	return nil
}

//
// Flatten -
//
func Flatten(s *Seq) *Seq {
	return nil
}

//
// Filter -
//
func Filter(f func(e interface{}) bool, s *Seq) *Seq {
	return nil
}

//
// Map -
//
func Map(f func(e interface{}) interface{}, s *Seq) *Seq {
	return nil
}

//
// Zip -
//
func Zip(s *Seq, t *Seq) *Seq {
	return nil
}

//
// ZipWith -
//
func ZipWith(f func(a, b interface{}) interface{}, s *Seq, t *Seq) *Seq {
	return nil
}

//
// Enum -
//
func Enum(s *Seq) *Seq {
	return nil
}

//
// FilterIdx -
//
func FilterIdx(f func(i int, a interface{}) bool, s *Seq) *Seq {
	return nil
}

//
// MapIdx -
//
func MapIdx(f func(i int, a interface{}) interface{}, s *Seq) *Seq {
	return nil
}

//
// Update -
//
func Update(s *Seq, i int, a interface{}) (*Seq, error) {
	return nil, nil
}

//
// Inject -
//
func Inject(s *Seq, t *Seq) (*Seq, error) {
	return nil, nil
}

//
// Subseq -
//
func Subseq(s *Seq, i int, j int) (*Seq, error) {
	return nil, nil
}
