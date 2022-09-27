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

locals {
  name_prefix = "${var.application}-${replace(var.region, "-", "_")}-${var.env}-${var.env_instance}"
  bucket_name = "${replace(local.name_prefix, "_", "")}-s3-000"
  path_403    = "" == var.path_403 ? "/${var.default_root_object}" : var.path_403
  path_404    = "" == var.path_404 ? "/${var.default_root_object}" : var.path_404
}

module "bucket" {
  source  = "terraform-aws-modules/s3-bucket/aws"
  version = "3.4.0"

  bucket = local.bucket_name

  # Ensure it's easy to replace
  force_destroy = true

  # Keep everything private
  acl                     = "private"
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

module "cloudfront" {
  source  = "terraform-aws-modules/cloudfront/aws"
  version = "~> 2"

  create_origin_access_identity = true

  origin_access_identities = {
    s3_bucket = local.bucket_name
  }

  is_ipv6_enabled = true

  price_class = "PriceClass_100"

  origin = {
    "S3-${local.bucket_name}" = {
      domain_name = "${local.bucket_name}.s3.amazonaws.com"
      s3_origin_config = {
        origin_access_identity = "s3_bucket"
      }
    }
  }

  origin_group = {}

  default_cache_behavior = {
    viewer_protocol_policy = "redirect-to-https"

    allowed_methods = ["GET", "HEAD"]
    cached_methods  = ["GET", "HEAD"]

    target_origin_id = "S3-${local.bucket_name}"
  }

  default_root_object = var.default_root_object

  custom_error_response = [
    {
      error_caching_min_ttl = "0"
      error_code            = "403"
      response_code         = "200"
      response_page_path    = local.path_403
    },
    {
      error_caching_min_ttl = "0"
      error_code            = "404"
      response_code         = "200"
      response_page_path    = local.path_404
    }
  ]
}
