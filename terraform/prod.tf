provider "google" {
  project = "project-01"
  region  = "us-central1"
}

resource "google_compute_network" "vpc_network" {
  name = "vpc-sql"
}

resource "google_sql_database_instance" "postgres_instance" {
  name             = "postgres-02"
  database_version = "POSTGRES_17"
  region           = "us-central1"

  settings {
    tier = "db-g1-small"

    ip_configuration {
      ipv4_enabled = true
    }
  }
}

resource "google_sql_database" "postgres_database" {
  name     = "db-01"
  instance = google_sql_database_instance.postgres_instance.name
}

resource "google_sql_user" "postgres_user" {
  name     = "login_user"
  instance = google_sql_database_instance.postgres_instance.name
  password = "6bEPgHy9jZKU7jN"
}