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

func resourceRoleAttachment() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceRoleAttachmentCreate),
		Read:   wrapCrudOperation(resourceRoleAttachmentRead),
		Delete: wrapCrudOperation(resourceRoleAttachmentDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"composite_role_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the composite role of this RoleAttachment.",
			},
			"attached_role_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the attached role of this RoleAttachment.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func convertRoleAttachmentFromResourceData(d *schema.ResourceData) *sdm.RoleAttachment {
	return &sdm.RoleAttachment{
		ID:              d.Id(),
		CompositeRoleID: convertStringFromResourceData(d, "composite_role_id"),
		AttachedRoleID:  convertStringFromResourceData(d, "attached_role_id"),
	}
}

func resourceRoleAttachmentCreate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	localVersion := convertRoleAttachmentFromResourceData(d)
	resp, err := cc.RoleAttachments().Create(ctx, localVersion)
	if err != nil {
		return fmt.Errorf("cannot create RoleAttachment: %w", err)
	}
	d.SetId(resp.RoleAttachment.ID)
	v := resp.RoleAttachment
	d.Set("composite_role_id", (v.CompositeRoleID))
	d.Set("attached_role_id", (v.AttachedRoleID))
	return nil
}

func resourceRoleAttachmentRead(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	localVersion := convertRoleAttachmentFromResourceData(d)
	_ = localVersion
	resp, err := cc.RoleAttachments().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read RoleAttachment %s: %w", d.Id(), err)
	}
	v := resp.RoleAttachment
	d.Set("composite_role_id", (v.CompositeRoleID))
	d.Set("attached_role_id", (v.AttachedRoleID))
	return nil
}
func resourceRoleAttachmentDelete(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	var errNotFound *sdm.NotFoundError
	_, err := cc.RoleAttachments().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
