package bloomst

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBloomst(t *testing.T) {

	Convey("test", t, func() {

		Convey("test new", func() {

			b, e := New("bloomst.bolt")
			So(e, ShouldBeNil)

			has, e := b.TestAndAdd([]byte("key"), []byte("love"))
			So(e, ShouldBeNil)
			So(has, ShouldEqual, false)

			has, e = b.TestAndAdd([]byte("key"), []byte("love"))
			So(e, ShouldBeNil)
			So(has, ShouldEqual, true)

			has, e = b.TestAndAdd([]byte("key"), []byte("love"))
			So(e, ShouldBeNil)
			So(has, ShouldEqual, true)
		})

	})
}
