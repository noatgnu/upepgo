// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessions)
	t.Run("UpepFeatures", testUpepFeatures)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiers)
	t.Run("UpepMolecularTypes", testUpepMolecularTypes)
	t.Run("UpepOrganisms", testUpepOrganisms)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBS)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntries)
}

func TestDelete(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsDelete)
	t.Run("UpepFeatures", testUpepFeaturesDelete)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersDelete)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesDelete)
	t.Run("UpepOrganisms", testUpepOrganismsDelete)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSDelete)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsQueryDeleteAll)
	t.Run("UpepFeatures", testUpepFeaturesQueryDeleteAll)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersQueryDeleteAll)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesQueryDeleteAll)
	t.Run("UpepOrganisms", testUpepOrganismsQueryDeleteAll)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSQueryDeleteAll)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsSliceDeleteAll)
	t.Run("UpepFeatures", testUpepFeaturesSliceDeleteAll)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersSliceDeleteAll)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesSliceDeleteAll)
	t.Run("UpepOrganisms", testUpepOrganismsSliceDeleteAll)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSSliceDeleteAll)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsExists)
	t.Run("UpepFeatures", testUpepFeaturesExists)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersExists)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesExists)
	t.Run("UpepOrganisms", testUpepOrganismsExists)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSExists)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesExists)
}

func TestFind(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsFind)
	t.Run("UpepFeatures", testUpepFeaturesFind)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersFind)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesFind)
	t.Run("UpepOrganisms", testUpepOrganismsFind)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSFind)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesFind)
}

func TestBind(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsBind)
	t.Run("UpepFeatures", testUpepFeaturesBind)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersBind)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesBind)
	t.Run("UpepOrganisms", testUpepOrganismsBind)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSBind)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesBind)
}

func TestOne(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsOne)
	t.Run("UpepFeatures", testUpepFeaturesOne)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersOne)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesOne)
	t.Run("UpepOrganisms", testUpepOrganismsOne)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSOne)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesOne)
}

func TestAll(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsAll)
	t.Run("UpepFeatures", testUpepFeaturesAll)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersAll)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesAll)
	t.Run("UpepOrganisms", testUpepOrganismsAll)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSAll)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesAll)
}

func TestCount(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsCount)
	t.Run("UpepFeatures", testUpepFeaturesCount)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersCount)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesCount)
	t.Run("UpepOrganisms", testUpepOrganismsCount)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSCount)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesCount)
}

func TestHooks(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsHooks)
	t.Run("UpepFeatures", testUpepFeaturesHooks)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersHooks)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesHooks)
	t.Run("UpepOrganisms", testUpepOrganismsHooks)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSHooks)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsInsert)
	t.Run("UpepAccessions", testUpepAccessionsInsertWhitelist)
	t.Run("UpepFeatures", testUpepFeaturesInsert)
	t.Run("UpepFeatures", testUpepFeaturesInsertWhitelist)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersInsert)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersInsertWhitelist)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesInsert)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesInsertWhitelist)
	t.Run("UpepOrganisms", testUpepOrganismsInsert)
	t.Run("UpepOrganisms", testUpepOrganismsInsertWhitelist)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSInsert)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSInsertWhitelist)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesInsert)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("UpepFeatureToUpepRefSeqEntryUsingRefSeqEntry", testUpepFeatureToOneUpepRefSeqEntryUsingRefSeqEntry)
	t.Run("UpepRefSeqEntryToUpepOrganismUsingOrganism", testUpepRefSeqEntryToOneUpepOrganismUsingOrganism)
	t.Run("UpepRefSeqEntryToUpepMolecularTypeUsingMolecularType", testUpepRefSeqEntryToOneUpepMolecularTypeUsingMolecularType)
	t.Run("UpepRefSeqEntryToUpepAccessionUsingAccession", testUpepRefSeqEntryToOneUpepAccessionUsingAccession)
	t.Run("UpepRefSeqEntryToUpepGeneIdentifierUsingGi", testUpepRefSeqEntryToOneUpepGeneIdentifierUsingGi)
	t.Run("UpepRefSeqEntryToUpepRefSeqDBUsingRefSeqDB", testUpepRefSeqEntryToOneUpepRefSeqDBUsingRefSeqDB)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("UpepAccessionToAccessionUpepRefSeqEntries", testUpepAccessionToManyAccessionUpepRefSeqEntries)
	t.Run("UpepGeneIdentifierToGiUpepRefSeqEntries", testUpepGeneIdentifierToManyGiUpepRefSeqEntries)
	t.Run("UpepMolecularTypeToMolecularTypeUpepRefSeqEntries", testUpepMolecularTypeToManyMolecularTypeUpepRefSeqEntries)
	t.Run("UpepOrganismToOrganismUpepRefSeqEntries", testUpepOrganismToManyOrganismUpepRefSeqEntries)
	t.Run("UpepRefSeqDBToRefSeqDBUpepRefSeqEntries", testUpepRefSeqDBToManyRefSeqDBUpepRefSeqEntries)
	t.Run("UpepRefSeqEntryToRefSeqEntryUpepFeatures", testUpepRefSeqEntryToManyRefSeqEntryUpepFeatures)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("UpepFeatureToUpepRefSeqEntryUsingRefSeqEntry", testUpepFeatureToOneSetOpUpepRefSeqEntryUsingRefSeqEntry)
	t.Run("UpepRefSeqEntryToUpepOrganismUsingOrganism", testUpepRefSeqEntryToOneSetOpUpepOrganismUsingOrganism)
	t.Run("UpepRefSeqEntryToUpepMolecularTypeUsingMolecularType", testUpepRefSeqEntryToOneSetOpUpepMolecularTypeUsingMolecularType)
	t.Run("UpepRefSeqEntryToUpepAccessionUsingAccession", testUpepRefSeqEntryToOneSetOpUpepAccessionUsingAccession)
	t.Run("UpepRefSeqEntryToUpepGeneIdentifierUsingGi", testUpepRefSeqEntryToOneSetOpUpepGeneIdentifierUsingGi)
	t.Run("UpepRefSeqEntryToUpepRefSeqDBUsingRefSeqDB", testUpepRefSeqEntryToOneSetOpUpepRefSeqDBUsingRefSeqDB)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("UpepRefSeqEntryToUpepOrganismUsingOrganism", testUpepRefSeqEntryToOneRemoveOpUpepOrganismUsingOrganism)
	t.Run("UpepRefSeqEntryToUpepMolecularTypeUsingMolecularType", testUpepRefSeqEntryToOneRemoveOpUpepMolecularTypeUsingMolecularType)
	t.Run("UpepRefSeqEntryToUpepAccessionUsingAccession", testUpepRefSeqEntryToOneRemoveOpUpepAccessionUsingAccession)
	t.Run("UpepRefSeqEntryToUpepGeneIdentifierUsingGi", testUpepRefSeqEntryToOneRemoveOpUpepGeneIdentifierUsingGi)
	t.Run("UpepRefSeqEntryToUpepRefSeqDBUsingRefSeqDB", testUpepRefSeqEntryToOneRemoveOpUpepRefSeqDBUsingRefSeqDB)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("UpepAccessionToAccessionUpepRefSeqEntries", testUpepAccessionToManyAddOpAccessionUpepRefSeqEntries)
	t.Run("UpepGeneIdentifierToGiUpepRefSeqEntries", testUpepGeneIdentifierToManyAddOpGiUpepRefSeqEntries)
	t.Run("UpepMolecularTypeToMolecularTypeUpepRefSeqEntries", testUpepMolecularTypeToManyAddOpMolecularTypeUpepRefSeqEntries)
	t.Run("UpepOrganismToOrganismUpepRefSeqEntries", testUpepOrganismToManyAddOpOrganismUpepRefSeqEntries)
	t.Run("UpepRefSeqDBToRefSeqDBUpepRefSeqEntries", testUpepRefSeqDBToManyAddOpRefSeqDBUpepRefSeqEntries)
	t.Run("UpepRefSeqEntryToRefSeqEntryUpepFeatures", testUpepRefSeqEntryToManyAddOpRefSeqEntryUpepFeatures)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("UpepAccessionToAccessionUpepRefSeqEntries", testUpepAccessionToManySetOpAccessionUpepRefSeqEntries)
	t.Run("UpepGeneIdentifierToGiUpepRefSeqEntries", testUpepGeneIdentifierToManySetOpGiUpepRefSeqEntries)
	t.Run("UpepMolecularTypeToMolecularTypeUpepRefSeqEntries", testUpepMolecularTypeToManySetOpMolecularTypeUpepRefSeqEntries)
	t.Run("UpepOrganismToOrganismUpepRefSeqEntries", testUpepOrganismToManySetOpOrganismUpepRefSeqEntries)
	t.Run("UpepRefSeqDBToRefSeqDBUpepRefSeqEntries", testUpepRefSeqDBToManySetOpRefSeqDBUpepRefSeqEntries)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("UpepAccessionToAccessionUpepRefSeqEntries", testUpepAccessionToManyRemoveOpAccessionUpepRefSeqEntries)
	t.Run("UpepGeneIdentifierToGiUpepRefSeqEntries", testUpepGeneIdentifierToManyRemoveOpGiUpepRefSeqEntries)
	t.Run("UpepMolecularTypeToMolecularTypeUpepRefSeqEntries", testUpepMolecularTypeToManyRemoveOpMolecularTypeUpepRefSeqEntries)
	t.Run("UpepOrganismToOrganismUpepRefSeqEntries", testUpepOrganismToManyRemoveOpOrganismUpepRefSeqEntries)
	t.Run("UpepRefSeqDBToRefSeqDBUpepRefSeqEntries", testUpepRefSeqDBToManyRemoveOpRefSeqDBUpepRefSeqEntries)
}

func TestReload(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsReload)
	t.Run("UpepFeatures", testUpepFeaturesReload)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersReload)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesReload)
	t.Run("UpepOrganisms", testUpepOrganismsReload)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSReload)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsReloadAll)
	t.Run("UpepFeatures", testUpepFeaturesReloadAll)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersReloadAll)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesReloadAll)
	t.Run("UpepOrganisms", testUpepOrganismsReloadAll)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSReloadAll)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsSelect)
	t.Run("UpepFeatures", testUpepFeaturesSelect)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersSelect)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesSelect)
	t.Run("UpepOrganisms", testUpepOrganismsSelect)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSSelect)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsUpdate)
	t.Run("UpepFeatures", testUpepFeaturesUpdate)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersUpdate)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesUpdate)
	t.Run("UpepOrganisms", testUpepOrganismsUpdate)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSUpdate)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsSliceUpdateAll)
	t.Run("UpepFeatures", testUpepFeaturesSliceUpdateAll)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersSliceUpdateAll)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesSliceUpdateAll)
	t.Run("UpepOrganisms", testUpepOrganismsSliceUpdateAll)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSSliceUpdateAll)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("UpepAccessions", testUpepAccessionsUpsert)
	t.Run("UpepFeatures", testUpepFeaturesUpsert)
	t.Run("UpepGeneIdentifiers", testUpepGeneIdentifiersUpsert)
	t.Run("UpepMolecularTypes", testUpepMolecularTypesUpsert)
	t.Run("UpepOrganisms", testUpepOrganismsUpsert)
	t.Run("UpepRefSeqDBS", testUpepRefSeqDBSUpsert)
	t.Run("UpepRefSeqEntries", testUpepRefSeqEntriesUpsert)
}
