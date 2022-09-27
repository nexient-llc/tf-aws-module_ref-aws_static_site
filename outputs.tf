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

output "bucket" {
  value = module.bucket
}

output "bucket_arn" {
  value = module.bucket.s3_bucket_arn
}

output "cloudfront" {
  value = module.cloudfront
}

output "cloudfront_arn" {
  value = module.cloudfront.cloudfront_distribution_arn
}
