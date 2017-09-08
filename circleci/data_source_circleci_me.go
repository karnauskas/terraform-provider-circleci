package circleci

import (
	//	"log"
	//	"github.com/jszwedko/go-circleci"
	//	circleci "github.com/ryanlower/go-circleci"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceCircleciMe() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCircleciMeRead,

		Schema: map[string]*schema.Schema{
			"login": {
				Type:     schema.TypeString,
				Optional: false,
				ForceNew: true,
			},
		},
	}
}

func dataSourceCircleciMeRead(d *schema.ResourceData, meta interface{}) error {

	//	log.Println(meta)

	//	client := meta.()

	//	return client.Me()

	d.Set("login", "meme")

	return nil
}
