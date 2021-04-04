// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Constraint struct {
	_tab flatbuffers.Table
}

func GetRootAsConstraint(buf []byte, offset flatbuffers.UOffsetT) *Constraint {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Constraint{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Constraint) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Constraint) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Constraint) Tvar(obj *Var) *Var {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Var)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Constraint) Kind() Kind {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Constraint) MutateKind(n Kind) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func ConstraintStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ConstraintAddTvar(builder *flatbuffers.Builder, tvar flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(tvar), 0)
}
func ConstraintAddKind(builder *flatbuffers.Builder, kind byte) {
	builder.PrependByteSlot(1, kind, 0)
}
func ConstraintEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
