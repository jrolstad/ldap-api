CREATE EXTERNAL TABLE IF NOT EXISTS `user_data`(
    `Id` string,
    `Name` string,
    `Location` string,
    `Type` string
  )
ROW FORMAT SERDE 'org.apache.hadoop.hive.serde2.JsonSerDe'
LOCATION 's3://jrolstad-identityobjects/user/';