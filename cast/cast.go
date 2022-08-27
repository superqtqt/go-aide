package cast

// builder for Options
type Options struct {
	base    int
	bitSize int
}
type Option func(options *Options)

func WithBase(base int) Option {
	return func(options *Options) {
		options.base = base
	}
}

func ToInt(v interface{}, opt ...func(o *Options)) int {
	rs, _ := ToIntWithError(v, opt...)
	return rs
}

func ToInt32(v interface{}, opt ...func(o *Options)) int32 {
	rs, _ := ToIntWithError(v, opt...)
	return int32(rs)
}

func ToInt64(v interface{}, opt ...func(o *Options)) int64 {
	rs, _ := ToIntWithError(v, opt...)
	return int64(rs)
}

func ToUint(v interface{}, opt ...func(o *Options)) uint {
	rs, _ := ToIntWithError(v, opt...)
	return uint(rs)
}

func ToUint32(v interface{}, opt ...func(o *Options)) uint32 {
	rs, _ := ToIntWithError(v, opt...)
	return uint32(rs)
}

func ToUint64(v interface{}, opt ...func(o *Options)) uint64 {
	rs, _ := ToIntWithError(v, opt...)
	return uint64(rs)
}

func ToUint8(v interface{}, opt ...func(o *Options)) uint8 {
	rs, _ := ToIntWithError(v, opt...)
	return uint8(rs)
}

func ToUint16(v interface{}, opt ...func(o *Options)) uint16 {
	rs, _ := ToIntWithError(v, opt...)
	return uint16(rs)
}
func ToFloat64(v interface{}, opt ...func(o *Options)) float64 {
	rs, _ := ToFloatWithError(v, opt...)
	return rs
}

func ToFloat32(v interface{}, opt ...func(o *Options)) float32 {
	rs, _ := ToFloatWithError(v, opt...)
	return float32(rs)
}
