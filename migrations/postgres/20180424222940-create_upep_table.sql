
-- +migrate Up

create table if not exists upep.upep_ref_seq_db
(
	id bigserial not null
		constraint upep_ref_seq_db_pkey
			primary key,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	name text not null,
	version integer default 0 not null
)
;

create table if not exists upep.upep_organisms
(
	id bigserial not null
		constraint upep_organisms_pkey
			primary key,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	name text not null,
	upep_ref_seq_db_id bigint not null
		constraint upep_organisms_upep_ref_seq_db_id_fk
			references upep.upep_ref_seq_db
			ON DELETE CASCADE ON UPDATE CASCADE
)
;

create table if not exists upep.upep_gene_identifiers
(
	id bigserial not null
		constraint upep_gis_id_pk
			primary key,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	gi bigint not null,
	upep_ref_seq_db_id bigint not null
		constraint upep_gene_identifiers_upep_ref_seq_db_id_fk
			references upep.upep_ref_seq_db
)
;

create table if not exists upep.upep_accessions
(
	id bigserial not null
		constraint upep_accessions_pkey
			primary key,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	accession text not null,
	upep_ref_seq_db_id bigint not null
		constraint upep_accessions_upep_ref_seq_db_id_fk
			references upep.upep_ref_seq_db
				on update cascade on delete cascade
)
;

create table if not exists upep.upep_molecular_types
(
	id bigserial not null
		constraint upep_molecular_types_pkey
			primary key,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	name varchar(10) not null
)
;

create table if not exists upep.upep_ref_seq_entries
(
	id bigserial not null
		constraint upep_ref_seq_entries_pkey
			primary key,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	name varchar(20),
	organism_id bigint
		constraint upep_ref_seq_entries_upep_organisms_id_fk
			references upep.upep_organisms
				on update cascade on delete cascade,
	molecular_type_id bigint
		constraint upep_ref_seq_entries_upep_molecular_types_id_fk
			references upep.upep_molecular_types
				on update cascade on delete cascade,
	accession_id bigint
		constraint upep_ref_seq_entries_upep_accessions_id_fk
			references upep.upep_accessions
				on update cascade on delete cascade,
	gi_id bigint
		constraint upep_ref_seq_entries_upep_gis_id_fk
			references upep.upep_gene_identifiers
				on update cascade on delete cascade,
	ref_seq_db_id bigint
		constraint upep_ref_seq_entries_upep_ref_seq_db_id_fk
			references upep.upep_ref_seq_db
				on update cascade on delete cascade,
	ref_seq_sequence text not null
)
;

create table if not exists upep.upep_features
(
	id bigserial not null
		constraint upep_features_pkey
			primary key,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	feature_start integer not null,
	feature_end integer not null,
	name varchar(10) not null,
	partial_start boolean default false not null,
	partial_end boolean default false not null,
	ref_seq_entry_id bigint not null
		constraint upep_features_upep_ref_seq_entries_id_fk
			references upep.upep_ref_seq_entries
				on update cascade on delete cascade
)
;

create table if not exists upep.upep_blast_db
(
	id bigserial not null
		constraint upep_blast_db_pkey
			primary key,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	title text not null,
	path text not null,
	description text not null,
	upep_ref_seq_db_id bigint not null
		constraint upep_blast_db_upep_ref_seq_db_id_fk
			references upep.upep_ref_seq_db
				on update cascade on delete cascade,
	starting_codon_id bigint not null
		constraint upep_blast_db_upep_codon_id_fk
			references upep.upep_codon
				on update cascade on delete cascade,
	ending_codon_id bigint not null
		constraint upep_blast_db_upep_codon_id_fk_2
			references upep.upep_codon
				on update cascade on delete cascade
)
;


CREATE TABLE upep.upep_codon
(
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    starting_codon boolean DEFAULT false NOT NULL,
    ending_codon boolean DEFAULT false NOT NULL,
    sequence char(3) NOT NULL,
    id bigint DEFAULT nextval('upep.upep_codon_id_seq'::regclass) PRIMARY KEY NOT NULL
);
CREATE UNIQUE INDEX upep_codon_sequence_uindex ON upep.upep_codon (sequence);
INSERT INTO upep.upep_codon (updated_at) VALUES ('2018-05-29 14:18:03.555813');
INSERT INTO upep.upep_codon (updated_at) VALUES ('2018-05-29 14:18:03.557814');
INSERT INTO upep.upep_codon (updated_at) VALUES ('2018-05-29 14:18:03.558814');
INSERT INTO upep.upep_codon (updated_at) VALUES ('2018-05-29 14:18:03.558814');
INSERT INTO upep.upep_codon (updated_at) VALUES ('2018-05-29 14:18:03.558814');
INSERT INTO upep.upep_codon (updated_at) VALUES ('2018-05-29 14:18:03.558814');
INSERT INTO upep.upep_codon (updated_at) VALUES ('2018-05-29 14:18:03.558814');
INSERT INTO upep.upep_codon (updated_at) VALUES ('2018-05-29 14:18:03.558814');
INSERT INTO upep.upep_codon (updated_at) VALUES ('2018-05-29 14:18:03.559814');
INSERT INTO upep.upep_codon (updated_at) VALUES ('2018-05-29 14:18:03.559814');
INSERT INTO upep.upep_codon (updated_at) VALUES ('2018-05-29 14:18:03.559814');
INSERT INTO upep.upep_codon (updated_at) VALUES ('2018-05-29 14:18:03.559814');

create unique index if not exists upep_codon_sequence_uindex
	on upep.upep_codon (sequence)
;

create table if not exists upep.upep_sorf_positions
(
	id bigserial not null
		constraint upep_sorf_pos_pkey
			primary key,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	starting_position integer not null,
	ending_position integer not null,
	ref_seq_entry_id bigint not null
		constraint upep_sorf_pos_upep_ref_seq_entries_id_fk
			references upep.upep_ref_seq_entries
				on update cascade on delete cascade,
	starting_codon_id bigint not null
		constraint upep_sorf_pos_upep_codon_id_fk
			references upep.upep_codon
				on update cascade on delete cascade,
	ending_codon_id bigint not null
		constraint upep_sorf_pos_upep_codon_id_fk_2
			references upep.upep_codon
				on update cascade on delete cascade
)
;

-- +migrate Down
drop schema upep cascade ;
create schema upep;