package circleci

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_TOKEN", nil),
				Description: "API Token",
			},
		},

		ResourcesMap: map[string]*schema.Resource{},

		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	token := d.Get("token").(string)

	//	client := &circleci.Client{
	//		Token: token,
	//	}

	//	log.Println(client)

	config := Config{
		ApiToken: token,
	}

	return config.Client()
}
