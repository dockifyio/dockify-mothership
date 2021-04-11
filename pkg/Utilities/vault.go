package Utilities

import (
	"fmt"
	"github.com/hashicorp/vault/api"
)

func InitVault(vaultAddr string) (*api.Client, error) {
	config := &api.Config{
		Address: vaultAddr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetValuesFromVaultV2Api(vaultClient *api.Client , token string, vaultPath string, keyName string)  (string, error) {
	vaultClient.SetToken(token)
	c := vaultClient.Logical()
	secret, err := c.Read(vaultPath)
	if err != nil {
		return "", err
	}
	m, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", err
	}
	str := fmt.Sprintf("%v", m[keyName])
	return str, nil
}


