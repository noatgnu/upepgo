create table upep.upep_gene_identifiers
(
	id bigserial not null
		constraint upep_gis_id_pk
			primary key,
	created_at timestamp,
	updated_at timestamp,
	gi bigint not null
)
;

create table upep.upep_accessions
(
	id bigserial not null
		constraint upep_accessions_pkey
			primary key,
	created_at timestamp,
	updated_at timestamp,
	accession varchar(20) not null
)
;

create table upep.upep_ref_seq_db
(
	id bigserial not null
		constraint upep_ref_seq_db_pkey
			primary key,
	created_at timestamp,
	updated_at timestamp,
	name text not null,
	version integer default 0 not null
)
;

create table upep.upep_organisms
(
	id bigserial not null
		constraint upep_organisms_pkey
			primary key,
	created_at timestamp,
	updated_at timestamp,
	name text not null
)
;

create table upep.upep_molecular_types
(
	id bigserial not null
		constraint upep_molecular_types_pkey
			primary key,
	created_at timestamp,
	updated_at timestamp,
	name varchar(10) not null
)
;

create table upep.upep_ref_seq_entries
(
	id bigserial not null
		constraint upep_ref_seq_entries_pkey
			primary key,
	created_at timestamp,
	updated_at timestamp,
	name varchar(20),
	organism_id bigint
		constraint upep_ref_seq_entries_upep_organisms_id_fk
			references upep_organisms,
	molecular_type_id bigint
		constraint upep_ref_seq_entries_upep_molecular_types_id_fk
			references upep_molecular_types,
	accession_id bigint
		constraint upep_ref_seq_entries_upep_accessions_id_fk
			references upep_accessions,
	gi_id bigint
		constraint upep_ref_seq_entries_upep_gis_id_fk
			references upep_gene_identifiers,
	ref_seq_db_id bigint
		constraint upep_ref_seq_entries_upep_ref_seq_db_id_fk
			references upep_ref_seq_db,
	ref_seq_sequence text not null
)
;

create table upep.upep_features
(
	id bigserial not null
		constraint upep_features_pkey
			primary key,
	created_at timestamp,
	updated_at timestamp,
	feature_start integer not null,
	feature_end integer not null,
	name varchar(10) not null,
	partial_start boolean default false not null,
	partial_end boolean default false not null,
	ref_seq_entry_id bigint not null
		constraint upep_features_upep_ref_seq_entries_id_fk
			references upep_ref_seq_entries
)
;

