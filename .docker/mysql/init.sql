SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN = 0

SET @@GLOBAL.GTID_PURGED='';

CREATE DATABASE IF NOT EXISTS `keep_note_database`;