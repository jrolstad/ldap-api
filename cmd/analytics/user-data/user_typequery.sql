SET hive.cli.print.header=true;
SET hive.query.name=UserTypes;

SET hive.tez.input.format=org.apache.hadoop.hive.ql.io.CombineHiveInputFormat;
SET mapreduce.input.fileinputformat.split.maxsize=16777216;

WITH user_types AS (
    SELECT
        count(Id) as user_count,
        Type as user_type
    FROM user_data
    GROUP BY Type
)


SELECT
    user_type,
    user_count
FROM user_types;