// Copyright 2022 Nexient LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

variable "application" {
  type        = string
  description = "The name of the application"
}

variable "region" {
  type        = string
  description = "The AWS region to deploy to (eg us-west-2)"
}

variable "env" {
  type        = string
  description = "The environment to deploy to (eg dev)"
}

variable "env_instance" {
  type        = string
  description = "The environment instance number (eg 123)"
}

variable "default_root_object" {
  type        = string
  description = "The default item in the bucket to point to"
  default     = "index.html"
}

variable "path_403" {
  type        = string
  description = "An alternative object path to point 403s to. If empty, defaults to the path of the var.default_root_object."
  default     = ""
}

variable "path_404" {
  type        = string
  description = "An alternative object path to point 404s to. If empty, defaults to the path of the var.default_root_object."
  default     = ""
}
