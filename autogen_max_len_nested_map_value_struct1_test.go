// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package skyencoder

import (
	"errors"
	"math"

	"github.com/skycoin/skycoin/src/cipher/encoder"
)

// EncodeSizeMaxLenNestedMapValueStruct1 computes the size of an encoded object of type MaxLenNestedMapValueStruct1
func EncodeSizeMaxLenNestedMapValueStruct1(obj *MaxLenNestedMapValueStruct1) int {
	i0 := 0

	// obj.Foo
	i0 += 4
	for _, v := range obj.Foo {
		i1 := 0

		// k
		i1 += 8

		// v.Foo
		i1 += 4 + len(v.Foo)

		i0 += i1
	}

	return i0
}

// EncodeMaxLenNestedMapValueStruct1 encodes an object of type MaxLenNestedMapValueStruct1 to the buffer in encoder.Encoder
func EncodeMaxLenNestedMapValueStruct1(e *encoder.Encoder, obj *MaxLenNestedMapValueStruct1) error {

	// obj.Foo

	// obj.Foo length check
	if len(obj.Foo) > math.MaxUint32 {
		return errors.New("obj.Foo length exceeds math.MaxUint32")
	}

	// obj.Foo length
	e.Uint32(uint32(len(obj.Foo)))

	for k, v := range obj.Foo {

		// k
		e.Int64(k)

		// v.Foo maxlen check
		if len(v.Foo) > 3 {
			return encoder.ErrMaxLenExceeded
		}

		// v.Foo length check
		if len(v.Foo) > math.MaxUint32 {
			return errors.New("v.Foo length exceeds math.MaxUint32")
		}

		// v.Foo
		e.ByteSlice([]byte(v.Foo))

	}

	return nil
}

// DecodeMaxLenNestedMapValueStruct1 decodes an object of type MaxLenNestedMapValueStruct1 from the buffer in encoder.Decoder
func DecodeMaxLenNestedMapValueStruct1(d *encoder.Decoder, obj *MaxLenNestedMapValueStruct1) error {
	{
		// obj.Foo

		ul, err := d.Uint32()
		if err != nil {
			return err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return encoder.ErrBufferUnderflow
		}

		if length != 0 {
			obj.Foo = make(map[int64]MaxLenStringStruct1)

			for counter := 0; counter < length; counter++ {
				var k1 int64

				{
					// k1
					i, err := d.Int64()
					if err != nil {
						return err
					}
					k1 = i
				}

				if _, ok := obj.Foo[k1]; ok {
					return encoder.ErrMapDuplicateKeys
				}

				var v1 MaxLenStringStruct1

				{
					// v1.Foo

					ul, err := d.Uint32()
					if err != nil {
						return err
					}

					length := int(ul)
					if length < 0 || length > len(d.Buffer) {
						return encoder.ErrBufferUnderflow
					}

					if length > 3 {
						return encoder.ErrMaxLenExceeded
					}

					v1.Foo = string(d.Buffer[:length])
					d.Buffer = d.Buffer[length:]
				}

				obj.Foo[k1] = v1
			}
		}
	}

	if len(d.Buffer) != 0 {
		return encoder.ErrRemainingBytes
	}

	return nil
}
