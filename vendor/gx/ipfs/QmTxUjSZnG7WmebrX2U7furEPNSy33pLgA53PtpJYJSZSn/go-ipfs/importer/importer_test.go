package importer

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"testing"

	chunk "gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/importer/chunk"
	dag "gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/merkledag"
	mdtest "gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/merkledag/test"
	uio "gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/unixfs/io"

	node "gx/ipfs/QmNwUEK7QbwSqyKBu3mMtToo8SUc6wQJ7gdZq4gGGJqfnf/go-ipld-format"
	u "gx/ipfs/QmPsAfmDBnZN3kZGSuNwvCNDZiHneERSKmRcFyG3UkvcT3/go-ipfs-util"
)

func getBalancedDag(t testing.TB, size int64, blksize int64) (node.Node, dag.DAGService) {
	ds := mdtest.Mock()
	r := io.LimitReader(u.NewTimeSeededRand(), size)
	nd, err := BuildDagFromReader(ds, chunk.NewSizeSplitter(r, blksize))
	if err != nil {
		t.Fatal(err)
	}
	return nd, ds
}

func getTrickleDag(t testing.TB, size int64, blksize int64) (node.Node, dag.DAGService) {
	ds := mdtest.Mock()
	r := io.LimitReader(u.NewTimeSeededRand(), size)
	nd, err := BuildTrickleDagFromReader(ds, chunk.NewSizeSplitter(r, blksize))
	if err != nil {
		t.Fatal(err)
	}
	return nd, ds
}

func TestBalancedDag(t *testing.T) {
	ds := mdtest.Mock()
	buf := make([]byte, 10000)
	u.NewTimeSeededRand().Read(buf)
	r := bytes.NewReader(buf)

	nd, err := BuildDagFromReader(ds, chunk.DefaultSplitter(r))
	if err != nil {
		t.Fatal(err)
	}

	dr, err := uio.NewDagReader(context.Background(), nd, ds)
	if err != nil {
		t.Fatal(err)
	}

	out, err := ioutil.ReadAll(dr)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(out, buf) {
		t.Fatal("bad read")
	}
}

func BenchmarkBalancedReadSmallBlock(b *testing.B) {
	b.StopTimer()
	nbytes := int64(10000000)
	nd, ds := getBalancedDag(b, nbytes, 4096)

	b.SetBytes(nbytes)
	b.StartTimer()
	runReadBench(b, nd, ds)
}

func BenchmarkTrickleReadSmallBlock(b *testing.B) {
	b.StopTimer()
	nbytes := int64(10000000)
	nd, ds := getTrickleDag(b, nbytes, 4096)

	b.SetBytes(nbytes)
	b.StartTimer()
	runReadBench(b, nd, ds)
}

func BenchmarkBalancedReadFull(b *testing.B) {
	b.StopTimer()
	nbytes := int64(10000000)
	nd, ds := getBalancedDag(b, nbytes, chunk.DefaultBlockSize)

	b.SetBytes(nbytes)
	b.StartTimer()
	runReadBench(b, nd, ds)
}

func BenchmarkTrickleReadFull(b *testing.B) {
	b.StopTimer()
	nbytes := int64(10000000)
	nd, ds := getTrickleDag(b, nbytes, chunk.DefaultBlockSize)

	b.SetBytes(nbytes)
	b.StartTimer()
	runReadBench(b, nd, ds)
}

func runReadBench(b *testing.B, nd node.Node, ds dag.DAGService) {
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithCancel(context.Background())
		read, err := uio.NewDagReader(ctx, nd, ds)
		if err != nil {
			b.Fatal(err)
		}

		_, err = read.WriteTo(ioutil.Discard)
		if err != nil && err != io.EOF {
			b.Fatal(err)
		}
		cancel()
	}
}
