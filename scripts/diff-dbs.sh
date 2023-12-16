#!/bin/sh

set -e

if [ $# -ne 4 ]; then
    echo "Compare 2 databases by <table>.<col>"
    echo "Usage: $0 <db1> <db2> <table> <col>"
    exit 1
fi

DB1=$1
DB2=$2
TABLE=$3
COL=$4

[ ! -f "$DB1" ] && (echo "$DB1 not found" && exit 2)
[ ! -f "$DB2" ] && (echo "$DB2 not found" && exit 2)

cat << EOF | sqlite3
ATTACH DATABASE '$DB1' as db1;
ATTACH DATABASE '$DB2' as db2;
SELECT 'db1 $TABLE count: ' || count(*) from db1.$TABLE;
SELECT 'db2 $TABLE count: ' || count(*) from db2.$TABLE;

SELECT
  '$DB2 missing ' AS db_name,
  t1.$COL
FROM
  db1.$TABLE t1
WHERE
  t1.$COL NOT IN (SELECT $COL FROM db2.$TABLE) AND t1.type_id = 12 AND t1.height < 10451817
UNION ALL
SELECT
  '$DB1 missing ' AS db_name,
  t2.$COL
FROM
  db2.$TABLE t2
WHERE
  t2.$COL NOT IN (SELECT $COL FROM $TABLE) AND t2.type_id = 12 AND t2.height < 10451817 ;

EOF

