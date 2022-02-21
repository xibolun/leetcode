package nc

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSolve(t *testing.T) {
	Convey("Test solve", t, func() {
		So(solve("12", "21"), ShouldEqual, "33")
		So(solve("120", "21"), ShouldEqual, "141")
		So(solve("10000", ""), ShouldEqual, "10000")
	})
}
