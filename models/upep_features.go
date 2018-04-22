// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
	"gopkg.in/volatiletech/null.v6"
)

// UpepFeature is an object representing the database table.
type UpepFeature struct {
	ID            int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt     null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt     null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	FeatureStart  int       `boil:"feature_start" json:"feature_start" toml:"feature_start" yaml:"feature_start"`
	FeatureEnd    int       `boil:"feature_end" json:"feature_end" toml:"feature_end" yaml:"feature_end"`
	Name          string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	PartialStart  bool      `boil:"partial_start" json:"partial_start" toml:"partial_start" yaml:"partial_start"`
	PartialEnd    bool      `boil:"partial_end" json:"partial_end" toml:"partial_end" yaml:"partial_end"`
	RefSeqEntryID int64     `boil:"ref_seq_entry_id" json:"ref_seq_entry_id" toml:"ref_seq_entry_id" yaml:"ref_seq_entry_id"`

	R *upepFeatureR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L upepFeatureL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UpepFeatureColumns = struct {
	ID            string
	CreatedAt     string
	UpdatedAt     string
	FeatureStart  string
	FeatureEnd    string
	Name          string
	PartialStart  string
	PartialEnd    string
	RefSeqEntryID string
}{
	ID:            "id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	FeatureStart:  "feature_start",
	FeatureEnd:    "feature_end",
	Name:          "name",
	PartialStart:  "partial_start",
	PartialEnd:    "partial_end",
	RefSeqEntryID: "ref_seq_entry_id",
}

// upepFeatureR is where relationships are stored.
type upepFeatureR struct {
	RefSeqEntry *UpepRefSeqEntry
}

// upepFeatureL is where Load methods for each relationship are stored.
type upepFeatureL struct{}

var (
	upepFeatureColumns               = []string{"id", "created_at", "updated_at", "feature_start", "feature_end", "name", "partial_start", "partial_end", "ref_seq_entry_id"}
	upepFeatureColumnsWithoutDefault = []string{"created_at", "updated_at", "feature_start", "feature_end", "name", "ref_seq_entry_id"}
	upepFeatureColumnsWithDefault    = []string{"id", "partial_start", "partial_end"}
	upepFeaturePrimaryKeyColumns     = []string{"id"}
)

type (
	// UpepFeatureSlice is an alias for a slice of pointers to UpepFeature.
	// This should generally be used opposed to []UpepFeature.
	UpepFeatureSlice []*UpepFeature
	// UpepFeatureHook is the signature for custom UpepFeature hook methods
	UpepFeatureHook func(boil.Executor, *UpepFeature) error

	upepFeatureQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	upepFeatureType                 = reflect.TypeOf(&UpepFeature{})
	upepFeatureMapping              = queries.MakeStructMapping(upepFeatureType)
	upepFeaturePrimaryKeyMapping, _ = queries.BindMapping(upepFeatureType, upepFeatureMapping, upepFeaturePrimaryKeyColumns)
	upepFeatureInsertCacheMut       sync.RWMutex
	upepFeatureInsertCache          = make(map[string]insertCache)
	upepFeatureUpdateCacheMut       sync.RWMutex
	upepFeatureUpdateCache          = make(map[string]updateCache)
	upepFeatureUpsertCacheMut       sync.RWMutex
	upepFeatureUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var upepFeatureBeforeInsertHooks []UpepFeatureHook
var upepFeatureBeforeUpdateHooks []UpepFeatureHook
var upepFeatureBeforeDeleteHooks []UpepFeatureHook
var upepFeatureBeforeUpsertHooks []UpepFeatureHook

var upepFeatureAfterInsertHooks []UpepFeatureHook
var upepFeatureAfterSelectHooks []UpepFeatureHook
var upepFeatureAfterUpdateHooks []UpepFeatureHook
var upepFeatureAfterDeleteHooks []UpepFeatureHook
var upepFeatureAfterUpsertHooks []UpepFeatureHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UpepFeature) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range upepFeatureBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UpepFeature) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range upepFeatureBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UpepFeature) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range upepFeatureBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UpepFeature) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range upepFeatureBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UpepFeature) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range upepFeatureAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UpepFeature) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range upepFeatureAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UpepFeature) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range upepFeatureAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UpepFeature) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range upepFeatureAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UpepFeature) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range upepFeatureAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUpepFeatureHook registers your hook function for all future operations.
func AddUpepFeatureHook(hookPoint boil.HookPoint, upepFeatureHook UpepFeatureHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		upepFeatureBeforeInsertHooks = append(upepFeatureBeforeInsertHooks, upepFeatureHook)
	case boil.BeforeUpdateHook:
		upepFeatureBeforeUpdateHooks = append(upepFeatureBeforeUpdateHooks, upepFeatureHook)
	case boil.BeforeDeleteHook:
		upepFeatureBeforeDeleteHooks = append(upepFeatureBeforeDeleteHooks, upepFeatureHook)
	case boil.BeforeUpsertHook:
		upepFeatureBeforeUpsertHooks = append(upepFeatureBeforeUpsertHooks, upepFeatureHook)
	case boil.AfterInsertHook:
		upepFeatureAfterInsertHooks = append(upepFeatureAfterInsertHooks, upepFeatureHook)
	case boil.AfterSelectHook:
		upepFeatureAfterSelectHooks = append(upepFeatureAfterSelectHooks, upepFeatureHook)
	case boil.AfterUpdateHook:
		upepFeatureAfterUpdateHooks = append(upepFeatureAfterUpdateHooks, upepFeatureHook)
	case boil.AfterDeleteHook:
		upepFeatureAfterDeleteHooks = append(upepFeatureAfterDeleteHooks, upepFeatureHook)
	case boil.AfterUpsertHook:
		upepFeatureAfterUpsertHooks = append(upepFeatureAfterUpsertHooks, upepFeatureHook)
	}
}

// OneP returns a single upepFeature record from the query, and panics on error.
func (q upepFeatureQuery) OneP() *UpepFeature {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single upepFeature record from the query.
func (q upepFeatureQuery) One() (*UpepFeature, error) {
	o := &UpepFeature{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for upep_features")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all UpepFeature records from the query, and panics on error.
func (q upepFeatureQuery) AllP() UpepFeatureSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all UpepFeature records from the query.
func (q upepFeatureQuery) All() (UpepFeatureSlice, error) {
	var o []*UpepFeature

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UpepFeature slice")
	}

	if len(upepFeatureAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all UpepFeature records in the query, and panics on error.
func (q upepFeatureQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all UpepFeature records in the query.
func (q upepFeatureQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count upep_features rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q upepFeatureQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q upepFeatureQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if upep_features exists")
	}

	return count > 0, nil
}

// RefSeqEntryG pointed to by the foreign key.
func (o *UpepFeature) RefSeqEntryG(mods ...qm.QueryMod) upepRefSeqEntryQuery {
	return o.RefSeqEntry(boil.GetDB(), mods...)
}

// RefSeqEntry pointed to by the foreign key.
func (o *UpepFeature) RefSeqEntry(exec boil.Executor, mods ...qm.QueryMod) upepRefSeqEntryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.RefSeqEntryID),
	}

	queryMods = append(queryMods, mods...)

	query := UpepRefSeqEntries(exec, queryMods...)
	queries.SetFrom(query.Query, "\"upep\".\"upep_ref_seq_entries\"")

	return query
} // LoadRefSeqEntry allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (upepFeatureL) LoadRefSeqEntry(e boil.Executor, singular bool, maybeUpepFeature interface{}) error {
	var slice []*UpepFeature
	var object *UpepFeature

	count := 1
	if singular {
		object = maybeUpepFeature.(*UpepFeature)
	} else {
		slice = *maybeUpepFeature.(*[]*UpepFeature)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &upepFeatureR{}
		}
		args[0] = object.RefSeqEntryID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &upepFeatureR{}
			}
			args[i] = obj.RefSeqEntryID
		}
	}

	query := fmt.Sprintf(
		"select * from \"upep\".\"upep_ref_seq_entries\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UpepRefSeqEntry")
	}
	defer results.Close()

	var resultSlice []*UpepRefSeqEntry
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UpepRefSeqEntry")
	}

	if len(upepFeatureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.RefSeqEntry = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RefSeqEntryID == foreign.ID {
				local.R.RefSeqEntry = foreign
				break
			}
		}
	}

	return nil
}

// SetRefSeqEntryG of the upep_feature to the related item.
// Sets o.R.RefSeqEntry to related.
// Adds o to related.R.RefSeqEntryUpepFeatures.
// Uses the global database handle.
func (o *UpepFeature) SetRefSeqEntryG(insert bool, related *UpepRefSeqEntry) error {
	return o.SetRefSeqEntry(boil.GetDB(), insert, related)
}

// SetRefSeqEntryP of the upep_feature to the related item.
// Sets o.R.RefSeqEntry to related.
// Adds o to related.R.RefSeqEntryUpepFeatures.
// Panics on error.
func (o *UpepFeature) SetRefSeqEntryP(exec boil.Executor, insert bool, related *UpepRefSeqEntry) {
	if err := o.SetRefSeqEntry(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetRefSeqEntryGP of the upep_feature to the related item.
// Sets o.R.RefSeqEntry to related.
// Adds o to related.R.RefSeqEntryUpepFeatures.
// Uses the global database handle and panics on error.
func (o *UpepFeature) SetRefSeqEntryGP(insert bool, related *UpepRefSeqEntry) {
	if err := o.SetRefSeqEntry(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetRefSeqEntry of the upep_feature to the related item.
// Sets o.R.RefSeqEntry to related.
// Adds o to related.R.RefSeqEntryUpepFeatures.
func (o *UpepFeature) SetRefSeqEntry(exec boil.Executor, insert bool, related *UpepRefSeqEntry) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"upep\".\"upep_features\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"ref_seq_entry_id"}),
		strmangle.WhereClause("\"", "\"", 2, upepFeaturePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RefSeqEntryID = related.ID

	if o.R == nil {
		o.R = &upepFeatureR{
			RefSeqEntry: related,
		}
	} else {
		o.R.RefSeqEntry = related
	}

	if related.R == nil {
		related.R = &upepRefSeqEntryR{
			RefSeqEntryUpepFeatures: UpepFeatureSlice{o},
		}
	} else {
		related.R.RefSeqEntryUpepFeatures = append(related.R.RefSeqEntryUpepFeatures, o)
	}

	return nil
}

// UpepFeaturesG retrieves all records.
func UpepFeaturesG(mods ...qm.QueryMod) upepFeatureQuery {
	return UpepFeatures(boil.GetDB(), mods...)
}

// UpepFeatures retrieves all the records using an executor.
func UpepFeatures(exec boil.Executor, mods ...qm.QueryMod) upepFeatureQuery {
	mods = append(mods, qm.From("\"upep\".\"upep_features\""))
	return upepFeatureQuery{NewQuery(exec, mods...)}
}

// FindUpepFeatureG retrieves a single record by ID.
func FindUpepFeatureG(id int64, selectCols ...string) (*UpepFeature, error) {
	return FindUpepFeature(boil.GetDB(), id, selectCols...)
}

// FindUpepFeatureGP retrieves a single record by ID, and panics on error.
func FindUpepFeatureGP(id int64, selectCols ...string) *UpepFeature {
	retobj, err := FindUpepFeature(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindUpepFeature retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUpepFeature(exec boil.Executor, id int64, selectCols ...string) (*UpepFeature, error) {
	upepFeatureObj := &UpepFeature{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"upep\".\"upep_features\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(upepFeatureObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from upep_features")
	}

	return upepFeatureObj, nil
}

// FindUpepFeatureP retrieves a single record by ID with an executor, and panics on error.
func FindUpepFeatureP(exec boil.Executor, id int64, selectCols ...string) *UpepFeature {
	retobj, err := FindUpepFeature(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *UpepFeature) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *UpepFeature) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *UpepFeature) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *UpepFeature) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no upep_features provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	if o.UpdatedAt.Time.IsZero() {
		o.UpdatedAt.Time = currTime
		o.UpdatedAt.Valid = true
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(upepFeatureColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	upepFeatureInsertCacheMut.RLock()
	cache, cached := upepFeatureInsertCache[key]
	upepFeatureInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			upepFeatureColumns,
			upepFeatureColumnsWithDefault,
			upepFeatureColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(upepFeatureType, upepFeatureMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(upepFeatureType, upepFeatureMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"upep\".\"upep_features\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"upep\".\"upep_features\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into upep_features")
	}

	if !cached {
		upepFeatureInsertCacheMut.Lock()
		upepFeatureInsertCache[key] = cache
		upepFeatureInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single UpepFeature record. See Update for
// whitelist behavior description.
func (o *UpepFeature) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single UpepFeature record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *UpepFeature) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the UpepFeature, and panics on error.
// See Update for whitelist behavior description.
func (o *UpepFeature) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the UpepFeature.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *UpepFeature) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	upepFeatureUpdateCacheMut.RLock()
	cache, cached := upepFeatureUpdateCache[key]
	upepFeatureUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			upepFeatureColumns,
			upepFeaturePrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update upep_features, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"upep\".\"upep_features\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, upepFeaturePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(upepFeatureType, upepFeatureMapping, append(wl, upepFeaturePrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update upep_features row")
	}

	if !cached {
		upepFeatureUpdateCacheMut.Lock()
		upepFeatureUpdateCache[key] = cache
		upepFeatureUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q upepFeatureQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q upepFeatureQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for upep_features")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o UpepFeatureSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o UpepFeatureSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o UpepFeatureSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UpepFeatureSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), upepFeaturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"upep\".\"upep_features\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, upepFeaturePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in upepFeature slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *UpepFeature) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *UpepFeature) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *UpepFeature) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *UpepFeature) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no upep_features provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(upepFeatureColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()

	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	upepFeatureUpsertCacheMut.RLock()
	cache, cached := upepFeatureUpsertCache[key]
	upepFeatureUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			upepFeatureColumns,
			upepFeatureColumnsWithDefault,
			upepFeatureColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			upepFeatureColumns,
			upepFeaturePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert upep_features, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(upepFeaturePrimaryKeyColumns))
			copy(conflict, upepFeaturePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"upep\".\"upep_features\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(upepFeatureType, upepFeatureMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(upepFeatureType, upepFeatureMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert upep_features")
	}

	if !cached {
		upepFeatureUpsertCacheMut.Lock()
		upepFeatureUpsertCache[key] = cache
		upepFeatureUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single UpepFeature record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *UpepFeature) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single UpepFeature record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *UpepFeature) DeleteG() error {
	if o == nil {
		return errors.New("models: no UpepFeature provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single UpepFeature record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *UpepFeature) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single UpepFeature record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UpepFeature) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no UpepFeature provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), upepFeaturePrimaryKeyMapping)
	sql := "DELETE FROM \"upep\".\"upep_features\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from upep_features")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q upepFeatureQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q upepFeatureQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no upepFeatureQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from upep_features")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o UpepFeatureSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o UpepFeatureSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no UpepFeature slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o UpepFeatureSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UpepFeatureSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no UpepFeature slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(upepFeatureBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), upepFeaturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"upep\".\"upep_features\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, upepFeaturePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from upepFeature slice")
	}

	if len(upepFeatureAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *UpepFeature) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *UpepFeature) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *UpepFeature) ReloadG() error {
	if o == nil {
		return errors.New("models: no UpepFeature provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UpepFeature) Reload(exec boil.Executor) error {
	ret, err := FindUpepFeature(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *UpepFeatureSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *UpepFeatureSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UpepFeatureSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty UpepFeatureSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UpepFeatureSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	upepFeatures := UpepFeatureSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), upepFeaturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"upep\".\"upep_features\".* FROM \"upep\".\"upep_features\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, upepFeaturePrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&upepFeatures)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UpepFeatureSlice")
	}

	*o = upepFeatures

	return nil
}

// UpepFeatureExists checks if the UpepFeature row exists.
func UpepFeatureExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"upep\".\"upep_features\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if upep_features exists")
	}

	return exists, nil
}

// UpepFeatureExistsG checks if the UpepFeature row exists.
func UpepFeatureExistsG(id int64) (bool, error) {
	return UpepFeatureExists(boil.GetDB(), id)
}

// UpepFeatureExistsGP checks if the UpepFeature row exists. Panics on error.
func UpepFeatureExistsGP(id int64) bool {
	e, err := UpepFeatureExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// UpepFeatureExistsP checks if the UpepFeature row exists. Panics on error.
func UpepFeatureExistsP(exec boil.Executor, id int64) bool {
	e, err := UpepFeatureExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}