#!/bin/bash
set -e
export PGPASSWORD=POSTGRES_PASSWORD;
pg_basebackup -P -R -X stream -c fast -h cart-postgres -U "$CART_DB_USER" -D ./