package hw10programoptimization

import (
	"archive/zip"
	"testing"
)

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		r, _ := zip.OpenReader("testdata/users.dat.zip")

		data, _ := r.File[0].Open()
		b.StartTimer()
		_, _ = GetDomainStat(data, "biz")
		b.StopTimer()

		r.Close()
	}
}

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		r, _ := zip.OpenReader("testdata/users.dat.zip")

		data, _ := r.File[0].Open()

		b.StartTimer()
		_, _ = GetDomainStatOld(data, "biz")
		b.StopTimer()

		r.Close()
	}
}
