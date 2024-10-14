-- poskita.auth_user definition
CREATE TABLE auth_user (
  id BIGSERIAL PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(250) NOT NULL,
  email_verified_at TIMESTAMP NULL DEFAULT NULL,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  is_superuser BOOLEAN NOT NULL,
  is_staff BOOLEAN NOT NULL,
  is_active BOOLEAN NOT NULL,
  photo VARCHAR(500) NOT NULL,
  date_joined TIMESTAMP NOT NULL,
  remember_token VARCHAR(100) DEFAULT NULL,
  created_at TIMESTAMP NULL DEFAULT NULL,
  created_by VARCHAR(100) DEFAULT NULL,
  updated_at TIMESTAMP NULL DEFAULT NULL,
  updated_by VARCHAR(100) DEFAULT NULL,
  deleted_at TIMESTAMP NULL DEFAULT NULL,
  deleted_by VARCHAR(100) DEFAULT NULL,
  CONSTRAINT auth_users_email_unique UNIQUE (email)
);

-- poskita.auth_audit definition
CREATE TABLE auth_audit (
  id BIGSERIAL PRIMARY KEY,
  created_by VARCHAR(50) DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL DEFAULT NULL,
  updated_by VARCHAR(100) DEFAULT NULL,
  user_id BIGINT NOT NULL,
  username VARCHAR(100) DEFAULT NULL,
  ip_address VARCHAR(50) DEFAULT NULL,
  activity_category VARCHAR(100) NOT NULL,
  activity_date TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',
  activity_name VARCHAR(100) NOT NULL,
  activity_ref_code VARCHAR(100) NOT NULL,
  activity_status VARCHAR(10) NOT NULL,
  device_info VARCHAR(100) DEFAULT NULL,
  device_id VARCHAR(100) DEFAULT NULL,
  error_code VARCHAR(10) NOT NULL,
  error_desc VARCHAR(1024) DEFAULT NULL
);

-- poskita.auth_company definition
CREATE TABLE auth_company (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) DEFAULT NULL,
  description VARCHAR(1000) DEFAULT NULL,
  address VARCHAR(500) DEFAULT NULL,
  phone VARCHAR(10) DEFAULT NULL,
  email VARCHAR(100) DEFAULT NULL,
  domain VARCHAR(100) DEFAULT NULL,
  village VARCHAR(100) DEFAULT NULL,
  district VARCHAR(100) DEFAULT NULL,
  city VARCHAR(100) DEFAULT NULL,
  province VARCHAR(100) DEFAULT NULL,
  postal_code VARCHAR(100) DEFAULT NULL,
  user_id BIGINT NOT NULL,
  created_at TIMESTAMP NULL DEFAULT NULL,
  created_by VARCHAR(100) DEFAULT NULL,
  updated_at TIMESTAMP NULL DEFAULT NULL,
  updated_by VARCHAR(100) DEFAULT NULL,
  deleted_at TIMESTAMP NULL DEFAULT NULL,
  deleted_by VARCHAR(100) DEFAULT NULL
);

-- poskita.auth_user_company_privilage definition
CREATE TABLE auth_user_company_privilage (
  user_id BIGINT NOT NULL,
  company_id BIGINT NOT NULL,
  privilage_id BIGINT NOT NULL
);

-- poskita.auth_user_company_role definition
CREATE TABLE auth_user_company_role (
  user_id BIGINT NOT NULL,
  company_id BIGINT NOT NULL,
  role_id BIGINT NOT NULL
);

-- poskita.auth_privilage definition
CREATE TABLE auth_privilage (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) DEFAULT NULL,
  desctription VARCHAR(1000) DEFAULT NULL,
  created_at TIMESTAMP NULL DEFAULT NULL,
  created_by VARCHAR(100) DEFAULT NULL,
  updated_at TIMESTAMP NULL DEFAULT NULL,
  updated_by VARCHAR(100) DEFAULT NULL,
  deleted_at TIMESTAMP NULL DEFAULT NULL,
  deleted_by VARCHAR(100) DEFAULT NULL
);

-- poskita.auth_role definition
CREATE TABLE auth_role (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) DEFAULT NULL,
  desctription VARCHAR(1000) DEFAULT NULL,
  created_at TIMESTAMP NULL DEFAULT NULL,
  created_by VARCHAR(100) DEFAULT NULL,
  updated_at TIMESTAMP NULL DEFAULT NULL,
  updated_by VARCHAR(100) DEFAULT NULL,
  deleted_at TIMESTAMP NULL DEFAULT NULL,
  deleted_by VARCHAR(100) DEFAULT NULL
);
