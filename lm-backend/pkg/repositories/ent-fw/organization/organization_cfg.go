package organization_repo

import "license-manager/pkg/repositories/ent-fw/ent"

func WithDriverAndURL(driver string, url string) func(*orgEntRepo) {
	return func(repo *orgEntRepo) {
		client, err := ent.Open(driver, url)
		if err != nil {
			client.Close()
			panic(err)
		}
		repo.client = client
	}
}

func WithEntClient(client *ent.Client) func(*orgEntRepo) {
	return func(org *orgEntRepo) {
		org.client = client
	}
}

func WithAutoMigration(migrate bool) func(*orgEntRepo) {
	return func(org *orgEntRepo) {
		org._runMig = migrate
	}
}