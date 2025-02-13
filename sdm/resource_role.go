// This file was generated by protogen. DO NOT EDIT.

package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceRoleCreate),
		Read:   wrapCrudOperation(resourceRoleRead),
		Update: wrapCrudOperation(resourceRoleUpdate),
		Delete: wrapCrudOperation(resourceRoleDelete),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"access_rules": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: accessRulesJSONDiffSuppress,
				Description:      "AccessRules JSON encoded access rules data.",
			},
			"composite": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "True if the Role is a composite role.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique human-readable name of the Role.",
			},
			"tags": {
				Type: schema.TypeMap,

				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Tags is a map of key, value pairs.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func convertRoleFromResourceData(d *schema.ResourceData) *sdm.Role {
	return &sdm.Role{
		ID:          d.Id(),
		AccessRules: convertStringFromResourceData(d, "access_rules"),
		Composite:   convertBoolFromResourceData(d, "composite"),
		Name:        convertStringFromResourceData(d, "name"),
		Tags:        convertTagsFromResourceData(d, "tags"),
	}
}

func resourceRoleCreate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	localVersion := convertRoleFromResourceData(d)

	resp, err := cc.Roles().Create(ctx, localVersion)
	if err != nil {
		return fmt.Errorf("cannot create Role: %w", err)
	}
	d.SetId(resp.Role.ID)
	v := resp.Role
	d.Set("access_rules", (v.AccessRules))
	d.Set("composite", (v.Composite))
	d.Set("name", (v.Name))
	d.Set("tags", convertTagsToMap(v.Tags))
	return nil
}

func resourceRoleRead(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	localVersion := convertRoleFromResourceData(d)
	_ = localVersion

	resp, err := cc.Roles().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read Role %s: %w", d.Id(), err)
	}
	v := resp.Role
	d.Set("access_rules", (v.AccessRules))
	d.Set("composite", (v.Composite))
	d.Set("name", (v.Name))
	d.Set("tags", convertTagsToMap(v.Tags))
	return nil
}
func resourceRoleUpdate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	resp, err := cc.Roles().Update(ctx, convertRoleFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot update Role %s: %w", d.Id(), err)
	}
	d.SetId(resp.Role.ID)
	return resourceRoleRead(d, cc)
}
func resourceRoleDelete(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	var errNotFound *sdm.NotFoundError
	_, err := cc.Roles().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
