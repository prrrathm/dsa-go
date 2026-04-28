
# bash
#!/bin/bash
# Script to generate DDL for a PostgreSQL database using a single connection string
# Replace the connection string with your actual values
# postgres://postgres:[YOUR-PASSWORD]@db.qingxwsbkvjgesrlldei.supabase.co:5432/postgres
CONNECTION_STRING="postgresql://postgres.qingxwsbkvjgesrlldei:pO8RKh5i1wkX60FR@aws-0-ap-south-1.pooler.supabase.com:5432/postgres"



OUTPUT_FILE=schema.ddl

# Run pg_dump to extract the schema (DDL) and save to a file
pg_dump --dbname="$CONNECTION_STRING" --schema-only --no-owner --no-privileges > $OUTPUT_FILE

# Check if the command was successful
if [ $? -eq 0 ]; then
  echo "DDL file generated successfully: $OUTPUT_FILE"
else
  echo "Error generating DDL file"
fi
