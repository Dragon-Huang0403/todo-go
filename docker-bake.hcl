variable "TAG" {
  default = "latest"
}

variable "REPOSITORY" {
  default = "dragon0huang"
}

variable "RELEASE" {
  default = true
}

variable "COMMIT" {}

group "default" {
  targets = ["todo"]
}

target "todo" {
  dockerfile = "./Dockerfile"
  platforms  = ["linux/amd64", "linux/arm64"]
  tags       = ["${REPOSITORY}/todo-go:${TAG}"]
  args = {
    COMMIT  = COMMIT
    VERSION = TAG
    RELEASE = RELEASE
    BASE    = RELEASE ? "scratch" : "alpine"
  }
}

