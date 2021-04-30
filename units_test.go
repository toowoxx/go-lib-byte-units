package byteunits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_String(t *testing.T) {
	assert.Equal(t, "B", B.String())
	assert.Equal(t, "kB", KB.String())
	assert.Equal(t, "MB", MB.String())
	assert.Equal(t, "GB", GB.String())
	assert.Equal(t, "TB", TB.String())
	assert.Equal(t, "PB", PB.String())
	assert.Equal(t, "EB", EB.String())
	assert.Equal(t, "KiB", KiB.String())
	assert.Equal(t, "MiB", MiB.String())
	assert.Equal(t, "GiB", GiB.String())
	assert.Equal(t, "TiB", TiB.String())
	assert.Equal(t, "PiB", PiB.String())
	assert.Equal(t, "EiB", EiB.String())
	assert.Equal(t, "B", Unit(123456).String())
}

func TestBytes_IECUnit(t *testing.T) {
	assert.Equal(t, KiB, (4 * KiB).IECUnit())
	assert.Equal(t, MiB, (4 * MiB).IECUnit())
	assert.Equal(t, MiB, (2 * MiB).IECUnit())
	assert.Equal(t, GiB, (999 * GiB).IECUnit())
	assert.Equal(t, GiB, (1023 * GiB).IECUnit())
	assert.Equal(t, GiB, (1000 * GiB).IECUnit())
	assert.NotEqual(t, GiB, (1024 * GiB).IECUnit())
	assert.Equal(t, TiB, (4 * TiB).IECUnit())
	assert.Equal(t, EiB, (1 * EiB).IECUnit())
	assert.Equal(t, EiB, (2 * EiB).IECUnit())
}

func TestBytes_MetricUnit(t *testing.T) {
	assert.Equal(t, KB, (4 * KB).MetricUnit())
	assert.Equal(t, MB, (4 * MB).MetricUnit())
	assert.Equal(t, MB, (2 * MB).MetricUnit())
	assert.Equal(t, GB, (999 * GB).MetricUnit())
	assert.NotEqual(t, GB, (1023 * GB).MetricUnit())
	assert.NotEqual(t, GB, (1000 * GB).MetricUnit())
	assert.NotEqual(t, GB, (1024 * GB).MetricUnit())
	assert.Equal(t, TB, (4 * TB).MetricUnit())
	assert.Equal(t, EB, (1 * EB).MetricUnit())
	assert.Equal(t, EB, (2 * EB).MetricUnit())
}

func TestBytes_StringWithSize(t *testing.T) {
	assert.Equal(t, "1.00 MiB", (1 * MiB).StringWithSize(true))
	assert.Equal(t, "2.05 kB", (2 * KiB).StringWithSize(false))
	assert.Equal(t, "2.50 kB", (2500 * B).StringWithSize(false))
	assert.Equal(t, "4.00 KiB", (4096 * B).StringWithSize(true))
	assert.Equal(t, "2.00 EB", (2 * EB).StringWithSize(false))
	assert.Equal(t, "1.52 EB", (1520 * PB).StringWithSize(false))
	assert.Equal(t, "2.00 kB", (1999 * B).StringWithSize(false))
	assert.Equal(t, "1.99 MB", (1994 * KB).StringWithSize(false))
	assert.Equal(t, "128 B", (128 * B).StringWithSize(false))
	assert.Equal(t, "128 B", (128 * B).StringWithSize(true))
	assert.Equal(t, "1 B", (1 * B).StringWithSize(true))
	assert.Equal(t, "1 B", (1 * B).StringWithSize(false))
}
