package data_employee

import "time"

type CreateEmployeeRequest struct {
	Id string `validate:"required" json:"id"`
}
type UpdateEmployeeRequest struct {
	Id string `validate:"required" json:"id"`
}

type EmployeeResponse struct {
	Id            int       `json:"id"`
	FirstName     *string   `json:"firstname"`
	LastName      *string   `json:"lastname"`
	ContactNo     *string   `json:"contact_no"`
	OfficialEmail *string   `json:"official_email"`
	PersonalEmail *string   `json:"Personal_email"`
	IdentityNo    *string   `json:"identity_no"`
	DateOfBirth   time.Time `json:"date_of_birth"`
}

/*

CREATE TABLE `employees` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `firstname` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `lastname` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `contact_no` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `official_email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `personal_email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `identity_no` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `date_of_birth` date DEFAULT NULL,
  `gender` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `emergency_contact_relationship` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `emergency_contact` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `emergency_contact_address` text COLLATE utf8mb4_unicode_ci,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `current_address` text COLLATE utf8mb4_unicode_ci,
  `permanent_address` text COLLATE utf8mb4_unicode_ci,
  `city` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `designation` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'office' COMMENT 'work from remote/office',
  `status` int NOT NULL DEFAULT '1',
  `employment_status` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `picture` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `joining_date` date DEFAULT NULL,
  `exit_date` date DEFAULT NULL,
  `gross_salary` int DEFAULT '0',
  `bonus` int DEFAULT '0',
  `branch_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `department_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `employees_official_email_unique` (`official_email`),
  UNIQUE KEY `employees_personal_email_unique` (`personal_email`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

*/

type Employees struct {
	Id                             int       `gorm:"type:int;primary_key"`
	Firstname                      string    `gorm:"type:varchar(255)"`
	Lastname                       string    `gorm:"type:varchar(255)"`
	Contact_no                     string    `gorm:"type:varchar(255)"`
	Official_email                 string    `gorm:"type:varchar(255)"`
	Personal_email                 string    `gorm:"type:varchar(255)"`
	Identity_no                    string    `gorm:"type:varchar(255)"`
	Date_of_birth                  time.Time `gorm:"type:date"`
	Gender                         string    `gorm:"type:varchar(255)"`
	Emergency_contact_relationship string    `gorm:"type:varchar(255)"`
	Emergency_contact              string    `gorm:"type:varchar(255)"`
	Emergency_contact_address      string    `gorm:"type:text"`
	City                           string    `gorm:"type:varchar(255)"`
	Designation                    string    `gorm:"type:varchar(255)"`
	Type                           string    `gorm:"type:varchar(255)"`
	Status                         int       `gorm:"type:int"`
	Employment_status              string    `gorm:"type:varchar(255)"`
	Picture                        string    `gorm:"type:varchar(255)"`
	Joining_date                   time.Time `gorm:"type:date"`
	Exit_date                      time.Time `gorm:"type:date"`
	Gross_salary                   int       `gorm:"type:int"`
	Bonus                          int       `gorm:"type:int"`
	Branch_id                      string    `gorm:"type:varchar(255)"`
	Department_id                  string    `gorm:"type:varchar(255)"`
	Remember_token                 string    `gorm:"type:varchar(100)"`
	Deleted_at                     time.Time `gorm:"type:timestamp"`
	Created_at                     time.Time `gorm:"type:timestamp"`
	Updated_at                     time.Time `gorm:"type:timestamp"`
}
