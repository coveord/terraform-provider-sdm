// This file was generated by protogen. DO NOT EDIT.

package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func dataSourceSecretStore() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceSecretStoreList),
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ca_cert_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_cert_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_key_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret_stores": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aws": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the SecretStore.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the SecretStore.",
									},
									"region": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
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
							},
						},
						"vault_tls": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ca_cert_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"client_cert_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"client_key_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the SecretStore.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the SecretStore.",
									},
									"namespace": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"server_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
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
							},
						},
						"vault_token": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique identifier of the SecretStore.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Unique human-readable name of the SecretStore.",
									},
									"namespace": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"server_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
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
							},
						},
					},
				},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}

func convertSecretStoreFilterFromResourceData(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if v, ok := d.GetOk("type"); ok {
		filter += "type:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("CA_cert_path"); ok {
		filter += "cacertpath:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("client_cert_path"); ok {
		filter += "clientcertpath:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("client_key_path"); ok {
		filter += "clientkeypath:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("id"); ok {
		filter += "id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("name"); ok {
		filter += "name:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("namespace"); ok {
		filter += "namespace:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("region"); ok {
		filter += "region:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("server_address"); ok {
		filter += "serveraddress:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("tags"); ok {
		filter += "tags:? "
		args = append(args, v)
	}
	return filter, args
}

func dataSourceSecretStoreList(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	filter, args := convertSecretStoreFilterFromResourceData(d)
	resp, err := cc.SecretStores().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list SecretStores %s: %w", d.Id(), err)
	}
	ids := []string{}
	type entity = map[string]interface{}
	output := make([]map[string][]entity, 1)
	output[0] = map[string][]entity{
		"aws": {},
	}
	for resp.Next() {
		ids = append(ids, resp.Value().GetID())
		switch v := resp.Value().(type) {
		case *sdm.AWSStore:
			output[0]["aws"] = append(output[0]["aws"], entity{
				"id":     (v.ID),
				"name":   (v.Name),
				"region": (v.Region),
				"tags":   convertTagsToMap(v.Tags),
			})
		case *sdm.VaultTLSStore:
			output[0]["vault_tls"] = append(output[0]["vault_tls"], entity{
				"ca_cert_path":     (v.CACertPath),
				"client_cert_path": (v.ClientCertPath),
				"client_key_path":  (v.ClientKeyPath),
				"id":               (v.ID),
				"name":             (v.Name),
				"namespace":        (v.Namespace),
				"server_address":   (v.ServerAddress),
				"tags":             convertTagsToMap(v.Tags),
			})
		case *sdm.VaultTokenStore:
			output[0]["vault_token"] = append(output[0]["vault_token"], entity{
				"id":             (v.ID),
				"name":           (v.Name),
				"namespace":      (v.Namespace),
				"server_address": (v.ServerAddress),
				"tags":           convertTagsToMap(v.Tags),
			})
		}
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", resp.Err())
	}

	err = d.Set("ids", ids)
	if err != nil {
		return fmt.Errorf("cannot set ids: %w", err)
	}
	err = d.Set("secret_stores", output)
	if err != nil {
		return fmt.Errorf("cannot set output: %w", err)
	}
	d.SetId("SecretStore" + filter + fmt.Sprint(args...))
	return nil
}
