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

func testUpepFeatures(t *testing.T) {
	t.Parallel()

	query := UpepFeatures(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testUpepFeaturesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = upepFeature.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepFeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUpepFeaturesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = UpepFeatures(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := UpepFeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUpepFeaturesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UpepFeatureSlice{upepFeature}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepFeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testUpepFeaturesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := UpepFeatureExists(tx, upepFeature.ID)
	if err != nil {
		t.Errorf("Unable to check if UpepFeature exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UpepFeatureExistsG to return true, but got false.")
	}
}
func testUpepFeaturesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	upepFeatureFound, err := FindUpepFeature(tx, upepFeature.ID)
	if err != nil {
		t.Error(err)
	}

	if upepFeatureFound == nil {
		t.Error("want a record, got nil")
	}
}
func testUpepFeaturesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = UpepFeatures(tx).Bind(upepFeature); err != nil {
		t.Error(err)
	}
}

func testUpepFeaturesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := UpepFeatures(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUpepFeaturesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeatureOne := &UpepFeature{}
	upepFeatureTwo := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeatureOne, upepFeatureDBTypes, false, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}
	if err = randomize.Struct(seed, upepFeatureTwo, upepFeatureDBTypes, false, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeatureOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = upepFeatureTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := UpepFeatures(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUpepFeaturesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	upepFeatureOne := &UpepFeature{}
	upepFeatureTwo := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeatureOne, upepFeatureDBTypes, false, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}
	if err = randomize.Struct(seed, upepFeatureTwo, upepFeatureDBTypes, false, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeatureOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = upepFeatureTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepFeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func upepFeatureBeforeInsertHook(e boil.Executor, o *UpepFeature) error {
	*o = UpepFeature{}
	return nil
}

func upepFeatureAfterInsertHook(e boil.Executor, o *UpepFeature) error {
	*o = UpepFeature{}
	return nil
}

func upepFeatureAfterSelectHook(e boil.Executor, o *UpepFeature) error {
	*o = UpepFeature{}
	return nil
}

func upepFeatureBeforeUpdateHook(e boil.Executor, o *UpepFeature) error {
	*o = UpepFeature{}
	return nil
}

func upepFeatureAfterUpdateHook(e boil.Executor, o *UpepFeature) error {
	*o = UpepFeature{}
	return nil
}

func upepFeatureBeforeDeleteHook(e boil.Executor, o *UpepFeature) error {
	*o = UpepFeature{}
	return nil
}

func upepFeatureAfterDeleteHook(e boil.Executor, o *UpepFeature) error {
	*o = UpepFeature{}
	return nil
}

func upepFeatureBeforeUpsertHook(e boil.Executor, o *UpepFeature) error {
	*o = UpepFeature{}
	return nil
}

func upepFeatureAfterUpsertHook(e boil.Executor, o *UpepFeature) error {
	*o = UpepFeature{}
	return nil
}

func testUpepFeaturesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &UpepFeature{}
	o := &UpepFeature{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, upepFeatureDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UpepFeature object: %s", err)
	}

	AddUpepFeatureHook(boil.BeforeInsertHook, upepFeatureBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	upepFeatureBeforeInsertHooks = []UpepFeatureHook{}

	AddUpepFeatureHook(boil.AfterInsertHook, upepFeatureAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	upepFeatureAfterInsertHooks = []UpepFeatureHook{}

	AddUpepFeatureHook(boil.AfterSelectHook, upepFeatureAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	upepFeatureAfterSelectHooks = []UpepFeatureHook{}

	AddUpepFeatureHook(boil.BeforeUpdateHook, upepFeatureBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	upepFeatureBeforeUpdateHooks = []UpepFeatureHook{}

	AddUpepFeatureHook(boil.AfterUpdateHook, upepFeatureAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	upepFeatureAfterUpdateHooks = []UpepFeatureHook{}

	AddUpepFeatureHook(boil.BeforeDeleteHook, upepFeatureBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	upepFeatureBeforeDeleteHooks = []UpepFeatureHook{}

	AddUpepFeatureHook(boil.AfterDeleteHook, upepFeatureAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	upepFeatureAfterDeleteHooks = []UpepFeatureHook{}

	AddUpepFeatureHook(boil.BeforeUpsertHook, upepFeatureBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	upepFeatureBeforeUpsertHooks = []UpepFeatureHook{}

	AddUpepFeatureHook(boil.AfterUpsertHook, upepFeatureAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	upepFeatureAfterUpsertHooks = []UpepFeatureHook{}
}
func testUpepFeaturesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepFeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUpepFeaturesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx, upepFeatureColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := UpepFeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUpepFeatureToOneUpepRefSeqEntryUsingRefSeqEntry(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local UpepFeature
	var foreign UpepRefSeqEntry

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, upepFeatureDBTypes, false, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
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

	slice := UpepFeatureSlice{&local}
	if err = local.L.LoadRefSeqEntry(tx, false, (*[]*UpepFeature)(&slice)); err != nil {
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

func testUpepFeatureToOneSetOpUpepRefSeqEntryUsingRefSeqEntry(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a UpepFeature
	var b, c UpepRefSeqEntry

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, upepFeatureDBTypes, false, strmangle.SetComplement(upepFeaturePrimaryKeyColumns, upepFeatureColumnsWithoutDefault)...); err != nil {
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

		if x.R.RefSeqEntryUpepFeatures[0] != &a {
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
func testUpepFeaturesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = upepFeature.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testUpepFeaturesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := UpepFeatureSlice{upepFeature}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testUpepFeaturesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := UpepFeatures(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	upepFeatureDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `FeatureEnd`: `integer`, `FeatureStart`: `integer`, `ID`: `bigint`, `Name`: `character varying`, `PartialEnd`: `boolean`, `PartialStart`: `boolean`, `RefSeqEntryID`: `bigint`, `UpdatedAt`: `timestamp with time zone`}
	_                  = bytes.MinRead
)

func testUpepFeaturesUpdate(t *testing.T) {
	t.Parallel()

	if len(upepFeatureColumns) == len(upepFeaturePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepFeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	if err = upepFeature.Update(tx); err != nil {
		t.Error(err)
	}
}

func testUpepFeaturesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(upepFeatureColumns) == len(upepFeaturePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	upepFeature := &UpepFeature{}
	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeatureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := UpepFeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, upepFeature, upepFeatureDBTypes, true, upepFeaturePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(upepFeatureColumns, upepFeaturePrimaryKeyColumns) {
		fields = upepFeatureColumns
	} else {
		fields = strmangle.SetComplement(
			upepFeatureColumns,
			upepFeaturePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(upepFeature))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := UpepFeatureSlice{upepFeature}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testUpepFeaturesUpsert(t *testing.T) {
	t.Parallel()

	if len(upepFeatureColumns) == len(upepFeaturePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	upepFeature := UpepFeature{}
	if err = randomize.Struct(seed, &upepFeature, upepFeatureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = upepFeature.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert UpepFeature: %s", err)
	}

	count, err := UpepFeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &upepFeature, upepFeatureDBTypes, false, upepFeaturePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UpepFeature struct: %s", err)
	}

	if err = upepFeature.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert UpepFeature: %s", err)
	}

	count, err = UpepFeatures(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
