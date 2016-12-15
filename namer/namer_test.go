package namer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var examples map[string]string = map[string]string{
	"060e8dde9cca5505d": "goosey-sprinkler",
	"0fad4fb184ed8dbf7": "mothproof-stitch",
	"009a1487e782619a5": "compatible-excrement",
}

func TestNameExpected(t *testing.T) {
	for example, expected := range examples {
		n, err := Name(example)
		assert.Nil(t, err, "namer returned error")
		assert.Equal(t, n, expected, "name should match expected")
	}
}

func TestNameInstancePrefix(t *testing.T) {
	n, err := Name(fmt.Sprintf("i-%s", "0fad4fb184ed8dbf7"))
	assert.Nil(t, err, "namer returned error")
	assert.Equal(t, n, examples["0fad4fb184ed8dbf7"], "name should match expected")
}

func TestNameInvalidId(t *testing.T) {
	n, err := Name("0fad4fb184ed8dbfz")
	assert.NotNil(t, err, "Should generate error")
	assert.Equal(t, n, "")
}

func TestNameBenchmark(t *testing.T) {
	r := testing.Benchmark(BenchmarkName)
	memUsageMB := float64(r.AllocedBytesPerOp()) / 1024 / 1024
	duration := time.Duration(r.NsPerOp())

	if r.AllocedBytesPerOp() > 100*1024*1024 {
		assert.Fail(t, "Should use less than 100MB memory", fmt.Sprintf("Used %dMB", memUsageMB))
	}

	if r.AllocsPerOp() > 1200000 {
		assert.Fail(t, "Should create less than 1200000 allocations", fmt.Sprintf("Used %d allocs", r.AllocsPerOp()))
	}

	if duration.Seconds() > 0.75 {
		assert.Fail(t, "Should complete in less than 750ms", fmt.Sprintf("Took %f", duration.Seconds()))
	}
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Name("060e8dde9cca5505d")
	}
}
