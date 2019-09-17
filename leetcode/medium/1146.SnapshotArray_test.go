package medium

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type versionItem struct {
	version int
	value   int
}

type arrayItem struct {
	snapshots   *[]versionItem
	currVersion versionItem
}

type SnapshotArray struct {
	items     []arrayItem
	snapCount int
}

func Constructor(length int) SnapshotArray {
	var items []arrayItem
	for i := 0; i < length; i++ {
		item := arrayItem{}
		item.snapshots = &[]versionItem{}
		item.currVersion = versionItem{}
		items = append(items, item)
	}

	arr := SnapshotArray{
		items:     items,
		snapCount: 0,
	}
	return arr
}

func (this *SnapshotArray) Set(index int, val int) {
	this.items[index].currVersion.value = val
}

func (this *SnapshotArray) Snap() int {
	this.snapCount++
	for _, item := range this.items {
		item.currVersion.version = this.snapCount - 1
		if len(*item.snapshots) > 0 && item.currVersion.value == (*item.snapshots)[len(*item.snapshots)-1].value {
			(*item.snapshots)[len(*item.snapshots)-1].version = item.currVersion.version
		} else {
			*item.snapshots = append(*item.snapshots, item.currVersion)
		}
	}
	return this.snapCount - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	for _, s := range *(this.items[index].snapshots) {
		if s.version >= snap_id {
			return s.value
		}
	}

	snap := this.items[index].snapshots
	return (*snap)[snap_id].value
}

func TestSnapshotArray(t *testing.T) {
	Convey("test", t, func() {
		arr := Constructor(5)

		arr.Set(1, 2)
		snapshotId := arr.Snap()
		So(snapshotId, ShouldEqual, 0)
		So(arr.Get(1, snapshotId), ShouldEqual, 2)

		arr.Set(1, 3)
		snapshotId = arr.Snap()
		So(snapshotId, ShouldEqual, 1)
		So(arr.Get(1, snapshotId), ShouldEqual, 3)
	})
}
