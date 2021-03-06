package benchmark

import (
	"net"
	"testing"
)

func BenchmarkStreamingTcpListener(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.wg.Add(1)
			streamTcpItem(t.conn)
		}
	})

	t.wg.Wait()
}

func BenchmarkStreamingTcpListenerParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t.wg.Add(1)
		streamTcpItem(t.conn)
	}

	t.wg.Wait()
}

func streamTcpItem(c net.Conn) {
	_, err := c.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
}
