package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
)

func main() {
	ctx := context.Background()

	config := vault.DefaultConfig()
	vaultAddr := os.Getenv("VAULT_ADDR")
	if vaultAddr == "" {
		fmt.Println("vault server was not provided in VAULT_ADDR env var")
	}

	client, err := vault.NewClient(config)
	if err != nil {
		log.Printf("Unable to initialize AppRole auth method: %v", err)
	}

	config.Address = vaultAddr

	roleID := os.Getenv("APPROLE_ROLE_ID")
	if roleID == "" {
		fmt.Println("no role ID was provided in APPROLE_ROLE_ID env var")
	}

	secretID := &auth.SecretID{FromString: "77c82ed9-dc64-4812-324d-ac49bcab512a"}

	appRoleAuth, err := auth.NewAppRoleAuth(
		roleID,
		secretID,
		//auth.WithWrappingToken(), // Only required if the secret ID is response-wrapped.
	)
	if err != nil {
		fmt.Println("unable to initialize AppRole auth method: %w", err)
	}

	authInfo, err := client.Auth().Login(ctx, appRoleAuth)
	if err != nil {
		fmt.Println("unable to login to AppRole auth method: %w", err)
	}
	if authInfo == nil {
		fmt.Println("no auth info was returned after login")
	}

	username := "world"
	password := "aeiou"

	userpassConfig := map[string]interface{}{
		"username": username,
		"password": password,
	}

	path := "auth/userpass/users/" + username
	_, err = client.Logical().Write(path, userpassConfig)
	if err != nil {
		log.Fatal(err)
	}

	policyName := "developer"

	policyPath := fmt.Sprintf("auth/userpass/users/%s", username)
	_, err = client.Logical().Write(policyPath, map[string]interface{}{
		"policies": policyName,
	})
	if err != nil {
		fmt.Println("policy attach fail")
	}

	fmt.Println("Userpass created successfully!")

}
