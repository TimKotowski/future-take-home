CREATE EXTENSION IF NOT EXISTS btree_gist;

CREATE TYPE appointments_status AS ENUM ('ACTIVE', 'CANCELLED', 'COMPLETED');
create table if not exists trainers (
  id                SERIAL PRIMARY KEY,
  first_name        text NOT NULL,
  last_name         text,
  years_experience  int DEFAULT 0
);

create table if not exists users (
  id                SERIAL PRIMARY KEY,
  first_name        text NOT NULL,
  last_name         text,
  email             text NOT NULL
);

-- https://www.postgresql.org/docs/12/rangetypes.html#RANGETYPES-INDEXING
create table if not exists appointments (
  id                    uuid PRIMARY KEY,
  trainer_id            bigint NOT NULL,
  user_id               bigint NOT NULL,
  start_slot            timestamptz NOT NULL,
  end_slot              timestamptz NOT NULL,
  status                appointments_status NOT NULL DEFAULT 'ACTIVE',
  created_at            timestamp DEFAULT now(),
  updated_at            timestamp DEFAULT now(),

  CONSTRAINT            "timeslot_range_valid" CHECK (start_slot < end_slot),
  CONSTRAINT            "timeslot_no_overlapping"
        EXCLUDE USING GIST    (trainer_id WITH =, tstzrange(start_slot, end_slot) WITH &&),
  CONSTRAINT "fk_trainers_trainer_id" FOREIGN KEY (trainer_id) REFERENCES trainers(id),
  CONSTRAINT "fk_users_user_id" FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX start_end_slots_trainers_composite_idx ON appointments(start_slot, end_slot, trainer_id);
CREATE INDEX appointments_status_idx on appointments(status);
CREATE INDEX fk_user_idx ON appointments(user_id);


