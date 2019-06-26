package mongodbatlas

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

//Provider returns the provider to be use by the code.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"public_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONGODB_ATLAS_PUBLIC_KEY", ""),
				Description: "MongoDB Atlas Programmatic Public Key",
			},
			"private_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONGODB_ATLAS_PRIVATE_KEY", ""),
				Description: "MongoDB Atlas Programmatic Private Key",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"mongodbatlas_database_user":  dataSourceMongoDBAtlasDatabaseUser(),
			"mongodbatlas_database_users": dataSourceMongoDBAtlasDatabaseUsers(),
			"mongodbatlas_project":        dataSourceMongoDBAtlasProject(),
			"mongodbatlas_projects":       dataSourceMongoDBAtlasProjects(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"mongodbatlas_database_user":           resourceMongoDBAtlasDatabaseUser(),
			"mongodbatlas_project_ip_whitelist":    resourceMongoDBAtlasProjectIPWhitelist(),
      "mongodbatlas_project":                 resourceMongoDBAtlasProject(),
			"mongodbatlas_cloud_provider_snapshot": resourceMongoDBAtlasCloudProviderSnapshot(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		PublicKey:  d.Get("public_key").(string),
		PrivateKey: d.Get("private_key").(string),
	}
	return config.NewClient(), nil
}
