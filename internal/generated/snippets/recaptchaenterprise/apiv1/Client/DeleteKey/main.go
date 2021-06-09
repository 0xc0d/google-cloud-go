// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START recaptchaenterprise_v1_generated_RecaptchaEnterpriseService_DeleteKey_sync]

package main

import (
	"context"

	recaptchaenterprise "cloud.google.com/go/recaptchaenterprise/apiv1"
	recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"
)

func main() {
	ctx := context.Background()
	c, err := recaptchaenterprise.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &recaptchaenterprisepb.DeleteKeyRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END recaptchaenterprise_v1_generated_RecaptchaEnterpriseService_DeleteKey_sync]