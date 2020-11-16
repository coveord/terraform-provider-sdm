// This file was generated by protogen. DO NOT EDIT.

package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	sdm "github.com/strongdm/strongdm-sdk-go"
)

func resourceAccountGrant() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceAccountGrantCreate),
		Read:   wrapCrudOperation(resourceAccountGrantRead),
		Delete: wrapCrudOperation(resourceAccountGrantDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the composite role of this AccountGrant.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the attached role of this AccountGrant.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func convertAccountGrantFromResourceData(d *schema.ResourceData) *sdm.AccountGrant {
	return &sdm.AccountGrant{
		ID:         d.Id(),
		ResourceID: convertStringFromResourceData(d, "resource_id"),
		AccountID:  convertStringFromResourceData(d, "account_id"),
	}
}

func resourceAccountGrantCreate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	localVersion := convertAccountGrantFromResourceData(d)
	resp, err := cc.AccountGrants().Create(ctx, localVersion)
	if err != nil {
		return fmt.Errorf("cannot create AccountGrant: %w", err)
	}
	d.SetId(resp.AccountGrant.ID)
	v := resp.AccountGrant
	d.Set("resource_id", (v.ResourceID))
	d.Set("account_id", (v.AccountID))
	return nil
}

func resourceAccountGrantRead(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	localVersion := convertAccountGrantFromResourceData(d)
	_ = localVersion
	resp, err := cc.AccountGrants().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read AccountGrant %s: %w", d.Id(), err)
	}
	v := resp.AccountGrant
	d.Set("resource_id", (v.ResourceID))
	d.Set("account_id", (v.AccountID))
	return nil
}
func resourceAccountGrantDelete(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	var errNotFound *sdm.NotFoundError
	_, err := cc.AccountGrants().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
