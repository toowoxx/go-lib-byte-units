// Package byteunits provides byte units
// as defined by the IEC and metric system.
// It does NOT follow JEDEC memory standards.
package byteunits

import (
	"fmt"
	"math/big"
)

type Bytes int64
type Unit = Bytes

const (
	// B is one byte (8 bits)
	B Unit = 1
)

const (
	// KB (written as kB) is 1000 B
	KB = 1000 * B
	// MB is 1000 kB
	MB = 1000 * KB
	// GB is 1000 MB
	GB = 1000 * MB
	// TB is 1000 GB
	TB = 1000 * GB
	// PB is 1000 TB
	PB = 1000 * TB
	// EB is 1000 PB
	EB = 1000 * PB
)

const (
	_ = B << (iota * 10)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
)

var stringRepresentations = map[Bytes]string{
	B: "B",
	KB: "kB",
	MB: "MB",
	GB: "GB",
	TB: "TB",
	PB: "PB",
	EB: "EB",
	KiB: "KiB",
	MiB: "MiB",
	GiB: "GiB",
	TiB: "TiB",
	PiB: "PiB",
	EiB: "EiB",
}

// String returns the string representation of
// the unit according to IEC or metric.
//
// Examples:
//     byteunits.MiB.String()
//     byteunits.KB.String()
// Results:
//     MiB
//     kB
func (b Unit) String() string {
	str, exists := stringRepresentations[b]
	if !exists {
		str = stringRepresentations[B]
	}
	return str
}

// IECUnit returns the IEC unit that would best fit the size
//
// Example:
//     (2 * MiB).IECUnit()
// Result:
//     MiB
func (b Bytes) IECUnit() Unit {
	switch {
	case b < KiB:
		return B
	case b < MiB:
		return KiB
	case b < GiB:
		return MiB
	case b < TiB:
		return GiB
	case b < PiB:
		return TiB
	case b < EiB:
		return PiB
	case b >= EiB:
		return EiB
	}
	return EiB
}

// MetricUnit returns the metric unit that would best fit the size
//
// Example:
//     (3 * MB).MetricUnit()
// Result:
//     MB
func (b Bytes) MetricUnit() Unit {
	switch {
	case b < KB:
		return B
	case b < MB:
		return KB
	case b < GB:
		return MB
	case b < TB:
		return GB
	case b < PB:
		return TB
	case b < EB:
		return PB
	case b >= EB:
		return EB
	}
	return EB
}

// StringWithSize returns a size + unit string
//
// If iec is true, the IEC standard will be used.
// IEC units have an "i" in the middle.
//
// Examples:
//     (10 * MiB).StringWithSize(true)
//     (2 * KiB).StringWithSize(false)
//     (3 * GB).StringWithSize(false)
// Results:
//     10 MiB
//     2.05 kB
//     3 GB
func (b Bytes) StringWithSize(iec bool) string {
	finalSize := big.NewFloat(0)
	bytes := big.NewFloat(0).SetInt64(int64(b))
	unit := B

	if iec {
		switch {
		case b < KiB:
			return fmt.Sprintf("%d %v", b, b)
		default:
			unit = b.IECUnit()
		}
	} else {
		switch {
		case b < KB:
			return fmt.Sprintf("%d %v", b, b)
		default:
			unit = b.MetricUnit()
		}
	}

	finalSize = bytes.Quo(bytes, big.NewFloat(0).SetInt64(int64(unit)))

	return fmt.Sprintf("%.2f %v", finalSize, unit)
}
