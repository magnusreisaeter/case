provider "google" {
  project = "project-01"
  region  = "us-central1"
}

resource "google_compute_network" "vpc_network" {
  name = "vpc-sql"
}

resource "google_sql_database_instance" "postgres_instance" {
  name             = "postgres-01"
  database_version = "POSTGRES_18"
  region           = "us-central1"

  settings {
    tier = "db-f1-micro"

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
  password = "NmJFUGdIeTlqWktVN2pO"
}