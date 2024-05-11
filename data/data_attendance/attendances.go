package data_attendance

type CreateAttendanceRequest struct {
	EmployeeId int `validate:"required" json:"employee_id"`
}
type UpdateAttendanceRequest struct {
	EmployeeId int `validate:"required" json:"employee_id"`
}

type AttendanceResponse struct {
	Id         int `json:"id"`
	EmployeeId int `json:"employee_id"`
}

/*

CREATE TABLE `attendances` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `employee_id` int NOT NULL,
  `time_in` time DEFAULT NULL,
  `time_out` time DEFAULT NULL,
  `timestamp_in` datetime DEFAULT NULL,
  `timestamp_out` datetime DEFAULT NULL,
  `comment` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `date` date NOT NULL,
  `status` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'present',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

*/

type Attendances struct {
	Id          int `gorm:"type:int;primary_key"`
	Employee_Id int `gorm:"type:int"`
}
