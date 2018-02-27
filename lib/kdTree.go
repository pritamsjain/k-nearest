package kdtree

import (
	"container/heap"
	"k-nearest/types"
	"fmt"
)

type kdTreeNode struct {
	axis           int
	splittingPoint types.Point
	leftChild      *kdTreeNode
	rightChild     *kdTreeNode
}

type KDTree struct {
	root *kdTreeNode
	dim  int
}

func (t *KDTree) Dim() int {
	return t.dim
}

func (t *KDTree) KNearestNeibors(target types.Point, k int) []types.Point {
	hp := &kNearestNHeapHelper{}
	t.searchNearest(t.root, hp, target, k)
	ret := make([]types.Point, 0, hp.Len())
	for hp.Len() > 0 {
		item := heap.Pop(hp).(*kNearestNodeHeap)
		ret = append(ret, item.point)
	}
	for i := len(ret)/2 - 1; i >= 0; i-- {
		opp := len(ret) - 1 - i
		ret[i], ret[opp] = ret[opp], ret[i]
	}
	return ret
}

func (t *KDTree) searchNearest(p *kdTreeNode,
	hp *kNearestNHeapHelper, target types.Point, k int) {
	stk := make([]*kdTreeNode, 0)
	for p != nil {
		stk = append(stk, p)
		if target.GetValue(p.axis) < p.splittingPoint.GetValue(p.axis) {
			p = p.leftChild
		} else {
			p = p.rightChild
		}
	}
	for i := len(stk) - 1; i >= 0; i-- {
		cur := stk[i]
		dist := target.Distance(cur.splittingPoint)
		if hp.Len() == 0 || (*hp)[0].distance >= dist {
			heap.Push(hp, &kNearestNodeHeap{
				point:    cur.splittingPoint,
				distance: dist,
			})
			if hp.Len() > k {
				heap.Pop(hp)
			}
		}
		if hp.Len() < k || target.PlaneDistance(
			cur.splittingPoint.GetValue(cur.axis), cur.axis) <=
			(*hp)[0].distance {
			if target.GetValue(cur.axis) < cur.splittingPoint.GetValue(cur.axis) {
				t.searchNearest(cur.rightChild, hp, target, k)
			} else {
				t.searchNearest(cur.leftChild, hp, target, k)
			}
		}
	}
}

func (t *KDTree) KFurthermostNeibors(target types.Point, k int) []types.Point {
	hp := &kFurtherestNHeapHelper{}
	t.searchFurthermost(t.root, hp, target, k)
	ret := make([]types.Point, 0, hp.Len())
	for hp.Len() > 0 {
		item := heap.Pop(hp).(*kFurtherestNodeHeap)
		ret = append(ret, item.point)
	}
	for i := len(ret)/2 - 1; i >= 0; i-- {
		opp := len(ret) - 1 - i
		ret[i], ret[opp] = ret[opp], ret[i]
	}
	return ret
}

func (t *KDTree) searchFurthermost(p *kdTreeNode,
	hp *kFurtherestNHeapHelper, target types.Point, k int) {
	stk := make([]*kdTreeNode, 0)
	for p != nil {
		stk = append(stk, p)
		if target.GetValue(p.axis) > p.splittingPoint.GetValue(p.axis) {
			p = p.leftChild
		} else {
			p = p.rightChild
		}
	}
	for i := len(stk) - 1; i >= 0; i-- {
		cur := stk[i]
		dist := target.Distance(cur.splittingPoint)
		if hp.Len() == 0 || (*hp)[0].distance <= dist {
			heap.Push(hp, &kFurtherestNodeHeap{
				point:    cur.splittingPoint,
				distance: dist,
			})
			if hp.Len() > k {
				heap.Pop(hp)
			}
		}
		if hp.Len() < k || target.PlaneDistance(
			cur.splittingPoint.GetValue(cur.axis), cur.axis) >=
			(*hp)[0].distance {
			if target.GetValue(cur.axis) > cur.splittingPoint.GetValue(cur.axis) {
				t.searchFurthermost(cur.rightChild, hp, target, k)
			} else {
				t.searchFurthermost(cur.leftChild, hp, target, k)
			}
		}
	}
}
func New(points []types.Point) *KDTree {
	if len(points) == 0 {
		return nil
	}
	ret := &KDTree{
		dim:  points[0].Dim(),
		root: createTree(points, 0),
	}
	return ret
}
//just like binary searchNearest tree but each level of tree is represented by one of axis like longitude or latitude
func createTree(points []types.Point, depth int) *kdTreeNode {
	if len(points) == 0 {
		return nil
	}
	dim := points[0].Dim()
	ret := &kdTreeNode{
		axis: depth % dim,
	}
	if len(points) == 1 {
		ret.splittingPoint = points[0]
		return ret
	}
	idx := selectSplittingPoint(points, ret.axis)
	if idx == -1 {
		return nil
	}
	ret.splittingPoint = points[idx]
	ret.leftChild = createTree(points[0:idx], depth+1)
	ret.rightChild = createTree(points[idx+1:], depth+1)
	return ret
}

type selectionHelper struct {
	axis   int
	points []types.Point
}

func (h *selectionHelper) Len() int {
	return len(h.points)
}

func (h *selectionHelper) Less(i, j int) bool {
	return h.points[i].GetValue(h.axis) < h.points[j].GetValue(h.axis)
}

func (h *selectionHelper) Swap(i, j int) {
	h.points[i], h.points[j] = h.points[j], h.points[i]
}

func selectSplittingPoint(points []types.Point, axis int) int {
	helper := &selectionHelper{
		axis:   axis,
		points: points,
	}
	mid := len(points)/2 + 1
	err := QuickSelect(helper, mid)
	if err != nil {
		fmt.Print(err)
		return -1
	}
	return mid - 1
}

type kNearestNodeHeap struct {
	point    types.Point
	distance float64
}

type kNearestNHeapHelper []*kNearestNodeHeap

func (h kNearestNHeapHelper) Len() int {
	return len(h)
}

func (h kNearestNHeapHelper) Less(i, j int) bool {
	return h[i].distance > h[j].distance
}

func (h kNearestNHeapHelper) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *kNearestNHeapHelper) Push(x interface{}) {
	item := x.(*kNearestNodeHeap)
	*h = append(*h, item)
}

func (h *kNearestNHeapHelper) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type kFurtherestNodeHeap struct {
	point    types.Point
	distance float64
}

type kFurtherestNHeapHelper []*kFurtherestNodeHeap

func (h kFurtherestNHeapHelper) Len() int {
	return len(h)
}

func (h kFurtherestNHeapHelper) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h kFurtherestNHeapHelper) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *kFurtherestNHeapHelper) Push(x interface{}) {
	item := x.(*kFurtherestNodeHeap)
	*h = append(*h, item)
}

func (h *kFurtherestNHeapHelper) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
