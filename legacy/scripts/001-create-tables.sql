create table IF not exists `report_task_table`
(
 `id`               INT(20) AUTO_INCREMENT,
 `task`             VARCHAR(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists `report_plugin_table`
(
 `id`               INT(20) AUTO_INCREMENT,
 `task`             INT(20),
 `start_time`       Datetime,
 `end_time`         Datetime,
 `memo`             VARCHAR(60),
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
  INDEX index_start_time(start_time),
  INDEX index_end_time(end_time),
  INDEX index_task(task),
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

