provider "aws" {
  region = "us-west-2"
}

resource "aws_db_instance" "default" {
  identifier        = "sal-db"
  allocated_storage = 20
  storage_type      = "gp2"
  engine            = "postgres"
  engine_version    = "9.6.3"
  instance_class    = "db.t2.micro"
  name              = "saldb"
  username          = "dave"
  password          = "CorpseMaster6661992"
  port              = "5432"

  skip_final_snapshot = true
}
