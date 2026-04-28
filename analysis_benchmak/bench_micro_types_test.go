//go:build bench

package benchmark

import (
	"testing"

	types "github.com/evcoreco/ocpp16types"
)

var (
	sinkDateTimeCustom    types.DateTime
	sinkDateTimePrimitive string

	sinkParentIDTagInfoCustom types.IDTagInfo
	sinkParentIDTagInfoPrim   primitiveIDTagInfo
	sinkParentIDTagValue      string
)

func BenchmarkDateTime_CustomType(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		dt, err := types.NewDateTime(benchmarkTimestamp)
		if err != nil {
			b.Fatal(err)
		}

		sinkDateTimeCustom = dt
	}
}

func BenchmarkDateTime_PrimitiveDirect(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		sinkDateTimePrimitive = benchmarkTimestamp
	}
}

func BenchmarkDateTime_PrimitiveValidated(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := validatePrimitiveTimestampUTC(benchmarkTimestamp); err != nil {
			b.Fatal(err)
		}

		sinkDateTimePrimitive = benchmarkTimestamp
	}
}

func BenchmarkParentIDTag_CustomChain(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		idTagInfo, err := buildCustomParentIDTagInfo("PARENT-TAG-001")
		if err != nil {
			b.Fatal(err)
		}

		sinkParentIDTagInfoCustom = idTagInfo
	}
}

func BenchmarkParentIDTag_PrimitiveDirect(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		parentIDTag := "PARENT-TAG-001"
		idTagInfo := primitiveIDTagInfo{ParentIDTag: &parentIDTag}
		sinkParentIDTagInfoPrim = idTagInfo
	}
}

func BenchmarkParentIDTag_PrimitiveValidated(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		parentIDTag := "PARENT-TAG-001"
		if err := validatePrimitiveCiString20(parentIDTag); err != nil {
			b.Fatal(err)
		}

		idTagInfo := primitiveIDTagInfo{ParentIDTag: &parentIDTag}
		sinkParentIDTagInfoPrim = idTagInfo
	}
}

func BenchmarkParentIDTag_CustomReadPrimitive(b *testing.B) {
	b.ReportAllocs()

	idTagInfo, err := buildCustomParentIDTagInfo("PARENT-TAG-001")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sinkParentIDTagValue = idTagInfo.ParentIDTag().Value().Value()
	}
}

func BenchmarkParentIDTag_PrimitiveRead(b *testing.B) {
	b.ReportAllocs()

	parentIDTag := "PARENT-TAG-001"
	idTagInfo := primitiveIDTagInfo{ParentIDTag: &parentIDTag}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sinkParentIDTagValue = *idTagInfo.ParentIDTag
	}
}

func buildCustomParentIDTagInfo(parentIDTag string) (types.IDTagInfo, error) {
	ciString, err := types.NewCiString20Type(parentIDTag)
	if err != nil {
		return types.IDTagInfo{}, err
	}

	idToken := types.NewIDToken(ciString)

	idTagInfo, err := types.NewIDTagInfo(types.AuthorizationStatusAccepted)
	if err != nil {
		return types.IDTagInfo{}, err
	}

	return idTagInfo.WithParentIDTag(idToken), nil
}
