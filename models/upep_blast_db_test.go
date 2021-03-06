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

func testUpepBlastDBS(t *testing.T) {
	t.Parallel()

	query := UpepBlastDBS(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testUpepBlastDBSDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = upepBlastDB.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepBlastDBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUpepBlastDBSQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = UpepBlastDBS(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := UpepBlastDBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUpepBlastDBSSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UpepBlastDBSlice{upepBlastDB}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepBlastDBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testUpepBlastDBSExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := UpepBlastDBExists(tx, upepBlastDB.ID)
	if err != nil {
		t.Errorf("Unable to check if UpepBlastDB exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UpepBlastDBExistsG to return true, but got false.")
	}
}
func testUpepBlastDBSFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	upepBlastDBFound, err := FindUpepBlastDB(tx, upepBlastDB.ID)
	if err != nil {
		t.Error(err)
	}

	if upepBlastDBFound == nil {
		t.Error("want a record, got nil")
	}
}
func testUpepBlastDBSBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = UpepBlastDBS(tx).Bind(upepBlastDB); err != nil {
		t.Error(err)
	}
}

func testUpepBlastDBSOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := UpepBlastDBS(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUpepBlastDBSAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDBOne := &UpepBlastDB{}
	upepBlastDBTwo := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDBOne, upepBlastDBDBTypes, false, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}
	if err = randomize.Struct(seed, upepBlastDBTwo, upepBlastDBDBTypes, false, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDBOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = upepBlastDBTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := UpepBlastDBS(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUpepBlastDBSCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upepBlastDBOne := &UpepBlastDB{}
	upepBlastDBTwo := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDBOne, upepBlastDBDBTypes, false, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}
	if err = randomize.Struct(seed, upepBlastDBTwo, upepBlastDBDBTypes, false, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDBOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = upepBlastDBTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepBlastDBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func upepBlastDBBeforeInsertHook(e boil.Executor, o *UpepBlastDB) error {
	*o = UpepBlastDB{}
	return nil
}

func upepBlastDBAfterInsertHook(e boil.Executor, o *UpepBlastDB) error {
	*o = UpepBlastDB{}
	return nil
}

func upepBlastDBAfterSelectHook(e boil.Executor, o *UpepBlastDB) error {
	*o = UpepBlastDB{}
	return nil
}

func upepBlastDBBeforeUpdateHook(e boil.Executor, o *UpepBlastDB) error {
	*o = UpepBlastDB{}
	return nil
}

func upepBlastDBAfterUpdateHook(e boil.Executor, o *UpepBlastDB) error {
	*o = UpepBlastDB{}
	return nil
}

func upepBlastDBBeforeDeleteHook(e boil.Executor, o *UpepBlastDB) error {
	*o = UpepBlastDB{}
	return nil
}

func upepBlastDBAfterDeleteHook(e boil.Executor, o *UpepBlastDB) error {
	*o = UpepBlastDB{}
	return nil
}

func upepBlastDBBeforeUpsertHook(e boil.Executor, o *UpepBlastDB) error {
	*o = UpepBlastDB{}
	return nil
}

func upepBlastDBAfterUpsertHook(e boil.Executor, o *UpepBlastDB) error {
	*o = UpepBlastDB{}
	return nil
}

func testUpepBlastDBSHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &UpepBlastDB{}
	o := &UpepBlastDB{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, upepBlastDBDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB object: %s", err)
	}

	AddUpepBlastDBHook(boil.BeforeInsertHook, upepBlastDBBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	upepBlastDBBeforeInsertHooks = []UpepBlastDBHook{}

	AddUpepBlastDBHook(boil.AfterInsertHook, upepBlastDBAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	upepBlastDBAfterInsertHooks = []UpepBlastDBHook{}

	AddUpepBlastDBHook(boil.AfterSelectHook, upepBlastDBAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	upepBlastDBAfterSelectHooks = []UpepBlastDBHook{}

	AddUpepBlastDBHook(boil.BeforeUpdateHook, upepBlastDBBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	upepBlastDBBeforeUpdateHooks = []UpepBlastDBHook{}

	AddUpepBlastDBHook(boil.AfterUpdateHook, upepBlastDBAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	upepBlastDBAfterUpdateHooks = []UpepBlastDBHook{}

	AddUpepBlastDBHook(boil.BeforeDeleteHook, upepBlastDBBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	upepBlastDBBeforeDeleteHooks = []UpepBlastDBHook{}

	AddUpepBlastDBHook(boil.AfterDeleteHook, upepBlastDBAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	upepBlastDBAfterDeleteHooks = []UpepBlastDBHook{}

	AddUpepBlastDBHook(boil.BeforeUpsertHook, upepBlastDBBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	upepBlastDBBeforeUpsertHooks = []UpepBlastDBHook{}

	AddUpepBlastDBHook(boil.AfterUpsertHook, upepBlastDBAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	upepBlastDBAfterUpsertHooks = []UpepBlastDBHook{}
}
func testUpepBlastDBSInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepBlastDBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUpepBlastDBSInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx, upepBlastDBColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := UpepBlastDBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUpepBlastDBToOneUpepRefSeqDBUsingUpepRefSeqDB(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UpepBlastDB
	var foreign UpepRefSeqDB

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, upepBlastDBDBTypes, false, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, upepRefSeqDBDBTypes, false, upepRefSeqDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepRefSeqDB struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.UpepRefSeqDBID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.UpepRefSeqDB(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UpepBlastDBSlice{&local}
	if err = local.L.LoadUpepRefSeqDB(tx, false, (*[]*UpepBlastDB)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.UpepRefSeqDB == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.UpepRefSeqDB = nil
	if err = local.L.LoadUpepRefSeqDB(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.UpepRefSeqDB == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUpepBlastDBToOneUpepCodonUsingStartingCodon(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UpepBlastDB
	var foreign UpepCodon

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, upepBlastDBDBTypes, false, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
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

	slice := UpepBlastDBSlice{&local}
	if err = local.L.LoadStartingCodon(tx, false, (*[]*UpepBlastDB)(&slice)); err != nil {
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

func testUpepBlastDBToOneUpepCodonUsingEndingCodon(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UpepBlastDB
	var foreign UpepCodon

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, upepBlastDBDBTypes, false, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
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

	slice := UpepBlastDBSlice{&local}
	if err = local.L.LoadEndingCodon(tx, false, (*[]*UpepBlastDB)(&slice)); err != nil {
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

func testUpepBlastDBToOneSetOpUpepRefSeqDBUsingUpepRefSeqDB(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UpepBlastDB
	var b, c UpepRefSeqDB

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, upepBlastDBDBTypes, false, strmangle.SetComplement(upepBlastDBPrimaryKeyColumns, upepBlastDBColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, upepRefSeqDBDBTypes, false, strmangle.SetComplement(upepRefSeqDBPrimaryKeyColumns, upepRefSeqDBColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, upepRefSeqDBDBTypes, false, strmangle.SetComplement(upepRefSeqDBPrimaryKeyColumns, upepRefSeqDBColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*UpepRefSeqDB{&b, &c} {
		err = a.SetUpepRefSeqDB(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.UpepRefSeqDB != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UpepBlastDBS[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UpepRefSeqDBID != x.ID {
			t.Error("foreign key was wrong value", a.UpepRefSeqDBID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UpepRefSeqDBID))
		reflect.Indirect(reflect.ValueOf(&a.UpepRefSeqDBID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UpepRefSeqDBID != x.ID {
			t.Error("foreign key was wrong value", a.UpepRefSeqDBID, x.ID)
		}
	}
}
func testUpepBlastDBToOneSetOpUpepCodonUsingStartingCodon(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UpepBlastDB
	var b, c UpepCodon

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, upepBlastDBDBTypes, false, strmangle.SetComplement(upepBlastDBPrimaryKeyColumns, upepBlastDBColumnsWithoutDefault)...); err != nil {
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

		if x.R.StartingCodonUpepBlastDBS[0] != &a {
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
func testUpepBlastDBToOneSetOpUpepCodonUsingEndingCodon(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UpepBlastDB
	var b, c UpepCodon

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, upepBlastDBDBTypes, false, strmangle.SetComplement(upepBlastDBPrimaryKeyColumns, upepBlastDBColumnsWithoutDefault)...); err != nil {
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

		if x.R.EndingCodonUpepBlastDBS[0] != &a {
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
func testUpepBlastDBSReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = upepBlastDB.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testUpepBlastDBSReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UpepBlastDBSlice{upepBlastDB}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testUpepBlastDBSSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := UpepBlastDBS(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	upepBlastDBDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `Description`: `text`, `EndingCodonID`: `bigint`, `ID`: `bigint`, `Path`: `text`, `StartingCodonID`: `bigint`, `Title`: `text`, `UpdatedAt`: `timestamp with time zone`, `UpepRefSeqDBID`: `bigint`}
	_                  = bytes.MinRead
)

func testUpepBlastDBSUpdate(t *testing.T) {
	t.Parallel()

	if len(upepBlastDBColumns) == len(upepBlastDBPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepBlastDBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	if err = upepBlastDB.Update(tx); err != nil {
		t.Error(err)
	}
}

func testUpepBlastDBSSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(upepBlastDBColumns) == len(upepBlastDBPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	upepBlastDB := &UpepBlastDB{}
	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepBlastDBS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, upepBlastDB, upepBlastDBDBTypes, true, upepBlastDBPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(upepBlastDBColumns, upepBlastDBPrimaryKeyColumns) {
		fields = upepBlastDBColumns
	} else {
		fields = strmangle.SetComplement(
			upepBlastDBColumns,
			upepBlastDBPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(upepBlastDB))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := UpepBlastDBSlice{upepBlastDB}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testUpepBlastDBSUpsert(t *testing.T) {
	t.Parallel()

	if len(upepBlastDBColumns) == len(upepBlastDBPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	upepBlastDB := UpepBlastDB{}
	if err = randomize.Struct(seed, &upepBlastDB, upepBlastDBDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepBlastDB.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert UpepBlastDB: %s", err)
	}

	count, err := UpepBlastDBS(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &upepBlastDB, upepBlastDBDBTypes, false, upepBlastDBPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UpepBlastDB struct: %s", err)
	}

	if err = upepBlastDB.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert UpepBlastDB: %s", err)
	}

	count, err = UpepBlastDBS(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
