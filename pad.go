package main

import "C"
import "github.com/ppreeper/pad"

//export pad_left
func pad_left(s *C.char, p *C.char, c C.int) *C.char {
	str := pad.Left(C.GoString(s), C.GoString(p), int(c))
	return C.CString(str)
}

//export pad_right
func pad_right(s *C.char, p *C.char, c C.int) *C.char {
	str := pad.Right(C.GoString(s), C.GoString(p), int(c))
	return C.CString(str)
}

//export pad_trunc_left
func pad_trunc_left(s *C.char, l C.int) *C.char {
	str := pad.TruncLeft(C.GoString(s), int(l))
	return C.CString(str)
}

//export pad_trunc_right
func pad_trunc_right(s *C.char, l C.int) *C.char {
	str := pad.TruncRight(C.GoString(s), int(l))
	return C.CString(str)
}

//export pad_left_len
func pad_left_len(s *C.char, p *C.char, l C.int) *C.char {
	str := pad.LeftLen(C.GoString(s), C.GoString(p), int(l))
	return C.CString(str)
}

//export pad_right_len
func pad_right_len(s *C.char, p *C.char, l C.int) *C.char {
	str := pad.RightLen(C.GoString(s), C.GoString(p), int(l))
	return C.CString(str)
}

//export pad_zfill
func pad_zfill(s *C.char, c C.int) *C.char {
	str := pad.ZFill(C.GoString(s), int(c))
	return C.CString(str)
}

//export pad_zfill_len
func pad_zfill_len(s *C.char, l C.int) *C.char {
	str := pad.ZFillLen(C.GoString(s), int(l))
	return C.CString(str)
}

//export pad_ljust
func pad_ljust(s *C.char, c C.int) *C.char {
	str := pad.LJust(C.GoString(s), int(c))
	return C.CString(str)
}

//export pad_ljust_len
func pad_ljust_len(s *C.char, l C.int) *C.char {
	str := pad.LJustLen(C.GoString(s), int(l))
	return C.CString(str)
}

//export pad_rjust
func pad_rjust(s *C.char, c C.int) *C.char {
	str := pad.RJust(C.GoString(s), int(c))
	return C.CString(str)
}

//export pad_rjust_len
func pad_rjust_len(s *C.char, l C.int) *C.char {
	str := pad.RJustLen(C.GoString(s), int(l))
	return C.CString(str)
}

func main() {}
