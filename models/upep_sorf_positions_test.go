// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testUpepSorfPositions(t *testing.T) {
	t.Parallel()

	query := UpepSorfPositions(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testUpepSorfPositionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = upepSorfPosition.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepSorfPositions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUpepSorfPositionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = UpepSorfPositions(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := UpepSorfPositions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUpepSorfPositionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UpepSorfPositionSlice{upepSorfPosition}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepSorfPositions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testUpepSorfPositionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := UpepSorfPositionExists(tx, upepSorfPosition.ID)
	if err != nil {
		t.Errorf("Unable to check if UpepSorfPosition exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UpepSorfPositionExistsG to return true, but got false.")
	}
}
func testUpepSorfPositionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	upepSorfPositionFound, err := FindUpepSorfPosition(tx, upepSorfPosition.ID)
	if err != nil {
		t.Error(err)
	}

	if upepSorfPositionFound == nil {
		t.Error("want a record, got nil")
	}
}
func testUpepSorfPositionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = UpepSorfPositions(tx).Bind(upepSorfPosition); err != nil {
		t.Error(err)
	}
}

func testUpepSorfPositionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := UpepSorfPositions(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUpepSorfPositionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPositionOne := &UpepSorfPosition{}
	upepSorfPositionTwo := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPositionOne, upepSorfPositionDBTypes, false, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}
	if err = randomize.Struct(seed, upepSorfPositionTwo, upepSorfPositionDBTypes, false, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPositionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = upepSorfPositionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := UpepSorfPositions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUpepSorfPositionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upepSorfPositionOne := &UpepSorfPosition{}
	upepSorfPositionTwo := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPositionOne, upepSorfPositionDBTypes, false, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}
	if err = randomize.Struct(seed, upepSorfPositionTwo, upepSorfPositionDBTypes, false, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPositionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = upepSorfPositionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepSorfPositions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func upepSorfPositionBeforeInsertHook(e boil.Executor, o *UpepSorfPosition) error {
	*o = UpepSorfPosition{}
	return nil
}

func upepSorfPositionAfterInsertHook(e boil.Executor, o *UpepSorfPosition) error {
	*o = UpepSorfPosition{}
	return nil
}

func upepSorfPositionAfterSelectHook(e boil.Executor, o *UpepSorfPosition) error {
	*o = UpepSorfPosition{}
	return nil
}

func upepSorfPositionBeforeUpdateHook(e boil.Executor, o *UpepSorfPosition) error {
	*o = UpepSorfPosition{}
	return nil
}

func upepSorfPositionAfterUpdateHook(e boil.Executor, o *UpepSorfPosition) error {
	*o = UpepSorfPosition{}
	return nil
}

func upepSorfPositionBeforeDeleteHook(e boil.Executor, o *UpepSorfPosition) error {
	*o = UpepSorfPosition{}
	return nil
}

func upepSorfPositionAfterDeleteHook(e boil.Executor, o *UpepSorfPosition) error {
	*o = UpepSorfPosition{}
	return nil
}

func upepSorfPositionBeforeUpsertHook(e boil.Executor, o *UpepSorfPosition) error {
	*o = UpepSorfPosition{}
	return nil
}

func upepSorfPositionAfterUpsertHook(e boil.Executor, o *UpepSorfPosition) error {
	*o = UpepSorfPosition{}
	return nil
}

func testUpepSorfPositionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &UpepSorfPosition{}
	o := &UpepSorfPosition{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, upepSorfPositionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition object: %s", err)
	}

	AddUpepSorfPositionHook(boil.BeforeInsertHook, upepSorfPositionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	upepSorfPositionBeforeInsertHooks = []UpepSorfPositionHook{}

	AddUpepSorfPositionHook(boil.AfterInsertHook, upepSorfPositionAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	upepSorfPositionAfterInsertHooks = []UpepSorfPositionHook{}

	AddUpepSorfPositionHook(boil.AfterSelectHook, upepSorfPositionAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	upepSorfPositionAfterSelectHooks = []UpepSorfPositionHook{}

	AddUpepSorfPositionHook(boil.BeforeUpdateHook, upepSorfPositionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	upepSorfPositionBeforeUpdateHooks = []UpepSorfPositionHook{}

	AddUpepSorfPositionHook(boil.AfterUpdateHook, upepSorfPositionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	upepSorfPositionAfterUpdateHooks = []UpepSorfPositionHook{}

	AddUpepSorfPositionHook(boil.BeforeDeleteHook, upepSorfPositionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	upepSorfPositionBeforeDeleteHooks = []UpepSorfPositionHook{}

	AddUpepSorfPositionHook(boil.AfterDeleteHook, upepSorfPositionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	upepSorfPositionAfterDeleteHooks = []UpepSorfPositionHook{}

	AddUpepSorfPositionHook(boil.BeforeUpsertHook, upepSorfPositionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	upepSorfPositionBeforeUpsertHooks = []UpepSorfPositionHook{}

	AddUpepSorfPositionHook(boil.AfterUpsertHook, upepSorfPositionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	upepSorfPositionAfterUpsertHooks = []UpepSorfPositionHook{}
}
func testUpepSorfPositionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepSorfPositions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUpepSorfPositionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx, upepSorfPositionColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := UpepSorfPositions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUpepSorfPositionToOneUpepRefSeqEntryUsingRefSeqEntry(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UpepSorfPosition
	var foreign UpepRefSeqEntry

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, upepSorfPositionDBTypes, false, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, upepRefSeqEntryDBTypes, false, upepRefSeqEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepRefSeqEntry struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.RefSeqEntryID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.RefSeqEntry(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UpepSorfPositionSlice{&local}
	if err = local.L.LoadRefSeqEntry(tx, false, (*[]*UpepSorfPosition)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.RefSeqEntry == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.RefSeqEntry = nil
	if err = local.L.LoadRefSeqEntry(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.RefSeqEntry == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUpepSorfPositionToOneUpepCodonUsingStartingCodon(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UpepSorfPosition
	var foreign UpepCodon

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, upepSorfPositionDBTypes, false, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, upepCodonDBTypes, false, upepCodonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepCodon struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.StartingCodonID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.StartingCodon(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UpepSorfPositionSlice{&local}
	if err = local.L.LoadStartingCodon(tx, false, (*[]*UpepSorfPosition)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.StartingCodon == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StartingCodon = nil
	if err = local.L.LoadStartingCodon(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.StartingCodon == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUpepSorfPositionToOneUpepCodonUsingEndingCodon(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UpepSorfPosition
	var foreign UpepCodon

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, upepSorfPositionDBTypes, false, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, upepCodonDBTypes, false, upepCodonColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepCodon struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.EndingCodonID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.EndingCodon(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UpepSorfPositionSlice{&local}
	if err = local.L.LoadEndingCodon(tx, false, (*[]*UpepSorfPosition)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.EndingCodon == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.EndingCodon = nil
	if err = local.L.LoadEndingCodon(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.EndingCodon == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUpepSorfPositionToOneSetOpUpepRefSeqEntryUsingRefSeqEntry(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UpepSorfPosition
	var b, c UpepRefSeqEntry

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, upepSorfPositionDBTypes, false, strmangle.SetComplement(upepSorfPositionPrimaryKeyColumns, upepSorfPositionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, upepRefSeqEntryDBTypes, false, strmangle.SetComplement(upepRefSeqEntryPrimaryKeyColumns, upepRefSeqEntryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, upepRefSeqEntryDBTypes, false, strmangle.SetComplement(upepRefSeqEntryPrimaryKeyColumns, upepRefSeqEntryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*UpepRefSeqEntry{&b, &c} {
		err = a.SetRefSeqEntry(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.RefSeqEntry != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RefSeqEntryUpepSorfPositions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RefSeqEntryID != x.ID {
			t.Error("foreign key was wrong value", a.RefSeqEntryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RefSeqEntryID))
		reflect.Indirect(reflect.ValueOf(&a.RefSeqEntryID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RefSeqEntryID != x.ID {
			t.Error("foreign key was wrong value", a.RefSeqEntryID, x.ID)
		}
	}
}
func testUpepSorfPositionToOneSetOpUpepCodonUsingStartingCodon(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UpepSorfPosition
	var b, c UpepCodon

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, upepSorfPositionDBTypes, false, strmangle.SetComplement(upepSorfPositionPrimaryKeyColumns, upepSorfPositionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, upepCodonDBTypes, false, strmangle.SetComplement(upepCodonPrimaryKeyColumns, upepCodonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, upepCodonDBTypes, false, strmangle.SetComplement(upepCodonPrimaryKeyColumns, upepCodonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*UpepCodon{&b, &c} {
		err = a.SetStartingCodon(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StartingCodon != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.StartingCodonUpepSorfPositions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StartingCodonID != x.ID {
			t.Error("foreign key was wrong value", a.StartingCodonID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StartingCodonID))
		reflect.Indirect(reflect.ValueOf(&a.StartingCodonID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StartingCodonID != x.ID {
			t.Error("foreign key was wrong value", a.StartingCodonID, x.ID)
		}
	}
}
func testUpepSorfPositionToOneSetOpUpepCodonUsingEndingCodon(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UpepSorfPosition
	var b, c UpepCodon

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, upepSorfPositionDBTypes, false, strmangle.SetComplement(upepSorfPositionPrimaryKeyColumns, upepSorfPositionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, upepCodonDBTypes, false, strmangle.SetComplement(upepCodonPrimaryKeyColumns, upepCodonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, upepCodonDBTypes, false, strmangle.SetComplement(upepCodonPrimaryKeyColumns, upepCodonColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*UpepCodon{&b, &c} {
		err = a.SetEndingCodon(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.EndingCodon != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.EndingCodonUpepSorfPositions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.EndingCodonID != x.ID {
			t.Error("foreign key was wrong value", a.EndingCodonID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.EndingCodonID))
		reflect.Indirect(reflect.ValueOf(&a.EndingCodonID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EndingCodonID != x.ID {
			t.Error("foreign key was wrong value", a.EndingCodonID, x.ID)
		}
	}
}
func testUpepSorfPositionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = upepSorfPosition.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testUpepSorfPositionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UpepSorfPositionSlice{upepSorfPosition}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testUpepSorfPositionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := UpepSorfPositions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	upepSorfPositionDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `EndingCodonID`: `bigint`, `EndingPosition`: `integer`, `ID`: `bigint`, `RefSeqEntryID`: `bigint`, `StartingCodonID`: `bigint`, `StartingPosition`: `integer`, `UpdatedAt`: `timestamp with time zone`}
	_                       = bytes.MinRead
)

func testUpepSorfPositionsUpdate(t *testing.T) {
	t.Parallel()

	if len(upepSorfPositionColumns) == len(upepSorfPositionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepSorfPositions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	if err = upepSorfPosition.Update(tx); err != nil {
		t.Error(err)
	}
}

func testUpepSorfPositionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(upepSorfPositionColumns) == len(upepSorfPositionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	upepSorfPosition := &UpepSorfPosition{}
	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepSorfPositions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, upepSorfPosition, upepSorfPositionDBTypes, true, upepSorfPositionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(upepSorfPositionColumns, upepSorfPositionPrimaryKeyColumns) {
		fields = upepSorfPositionColumns
	} else {
		fields = strmangle.SetComplement(
			upepSorfPositionColumns,
			upepSorfPositionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(upepSorfPosition))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := UpepSorfPositionSlice{upepSorfPosition}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testUpepSorfPositionsUpsert(t *testing.T) {
	t.Parallel()

	if len(upepSorfPositionColumns) == len(upepSorfPositionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	upepSorfPosition := UpepSorfPosition{}
	if err = randomize.Struct(seed, &upepSorfPosition, upepSorfPositionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepSorfPosition.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert UpepSorfPosition: %s", err)
	}

	count, err := UpepSorfPositions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &upepSorfPosition, upepSorfPositionDBTypes, false, upepSorfPositionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UpepSorfPosition struct: %s", err)
	}

	if err = upepSorfPosition.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert UpepSorfPosition: %s", err)
	}

	count, err = UpepSorfPositions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
