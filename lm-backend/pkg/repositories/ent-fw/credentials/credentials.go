package credentials_repo

import (
	"context"
	"errors"
	"fmt"
	"license-manager/pkg/domain"
	"license-manager/pkg/pkgerrors"
	"license-manager/pkg/repositories"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/credentials"
	ent_fw_common "license-manager/pkg/repositories/ent-fw/common"
)

type credentialsEntRepo struct {
	client *ent.Client
}

// *credentialsEntRepo implements repo interface
var _ repositories.CredentialsRepository = (*credentialsEntRepo)(nil)

func NewCredentialsEntRepo(client *ent.Client) *credentialsEntRepo {
	return &credentialsEntRepo{
		client: client,
	}
}

func (repo *credentialsEntRepo) DeleteByUserNameAndPasswordHash(name string, hash string) error {

	deleted, err := repo.client.Credentials.Delete().
		Where(credentials.And(
			credentials.UsernameEQ(name),
			credentials.PasswordHashEQ(hash),
		)).
		Exec(context.TODO())
	if err != nil {
		return err
	}
	
	if deleted <= 0 {
		return pkgerrors.ErrCredsNotFound
	}

	return nil
}

func (repo *credentialsEntRepo) FindByUserNameAndPasswordHash(name string, psswdHash string) (domain.Credentials, error) {
	
	creds, err := repo.client.Credentials.Query().
		Where(credentials.And(
			credentials.UsernameEQ(name),
			credentials.PasswordHashEQ(psswdHash),
		)).
		Only(context.TODO())
	if err != nil {
		return domain.Credentials{}, err
	}

	return toEntity(creds), nil
}

func (repo *credentialsEntRepo) MergeClaimsFor(name string, hash string, updates domain.Claims) (domain.Claims, error) {

	ctx := context.TODO()

	txClient, err := repo.client.Tx(ctx)
	if err != nil {
		return domain.Claims{}, err
	}

	// get current claims
	targetCreds, err := txClient.Credentials.Query().
		Where(credentials.And(
			credentials.UsernameEQ(name),
			credentials.PasswordHashEQ(hash),
		)).
		Only(ctx)
	if err != nil {
		if errors.Is(err, &ent.NotFoundError{}) {
			return domain.Claims{}, rollback(txClient, pkgerrors.ErrCredsNotFound)
		}
		return domain.Claims{}, rollback(txClient, err)
	}

	// merge
	currentClaims := toEntity(targetCreds).Claims
	resultClaims := ent_fw_common.MergeClaims(currentClaims, updates)
	
	// update creds
	_, err = txClient.Credentials.Update().
		Where(credentials.And(
			credentials.UsernameEQ(name),
			credentials.PasswordHashEQ(hash),
		)).
		SetClaims(resultClaims).
		Save(ctx)
	if err != nil {
		if errors.Is(err, &ent.NotFoundError{}) {
			return domain.Claims{}, rollback(txClient, pkgerrors.ErrCredsNotFound)
		}
		return domain.Claims{}, rollback(txClient, err)
	}

	err = txClient.Commit()
	if err != nil {
		return domain.Claims{}, err
	}

	return domain.Claims(resultClaims), nil
}

func (repo *credentialsEntRepo) Save(creds domain.Credentials) error {
	_, err := repo.client.Credentials.Create().
		SetUsername(creds.UserName).
		SetPasswordHash(creds.PasswordHash).
		SetClaims(creds.Claims).
		Save(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (repo *credentialsEntRepo) Update(creds domain.Credentials) error {
	_, err := repo.client.Credentials.Update().
		Where(credentials.And(
			credentials.UsernameEQ(creds.UserName),
			credentials.PasswordHashEQ(creds.PasswordHash),
		)).
		SetClaims(creds.Claims).
		Save(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func rollback(tx *ent.Tx, err error) error {
	if rbackErr := tx.Rollback(); rbackErr != nil {
		// wrap initial error with rollback error
		err = fmt.Errorf("%v: %w", rbackErr, err)
	}
	return err
}

func toEntity(dto *ent.Credentials) domain.Credentials {	
	return domain.Credentials{
		UserName: dto.Username,
		PasswordHash: dto.PasswordHash,
		Claims: dto.Claims,
	}
}
