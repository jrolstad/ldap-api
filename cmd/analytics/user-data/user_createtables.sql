DROP TABLE IF EXISTS `user`;

CREATE EXTERNAL TABLE IF NOT EXISTS `user`(
    `Id` string,
    `Name` string,
    `Location` string,
    `Type` string,
    `ObjectType` string,
    `Upn` string,
    `GivenName` string,
    `Surname` string,
    `Email` string,
    `Manager` string,
    `Company` string,
    `Department` string,
    `Title` string
  )
ROW FORMAT SERDE 'org.apache.hive.hcatalog.data.JsonSerDe'
LOCATION 's3://jrolstad-identityobjects/user/';