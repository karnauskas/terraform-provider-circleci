package circleci

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jszwedko/go-circleci"
)

func dataSourceCircleciMe() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCircleciMeRead,

		Schema: map[string]*schema.Schema{
			"login": &schema.Schema{
				Type:     schema.TypeString,
				Elem:     schema.TypeString,
				Computed: true,
			},
			"admin": &schema.Schema{
				Type:     schema.TypeBool,
				Elem:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceCircleciMeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*circleci.Client)

	me, err := client.Me()
	if err != nil {
		return err
	}

	d.SetId(me.Login)

	d.Set("login", me.Login)
	d.Set("admin", me.Admin)

	return nil
}
