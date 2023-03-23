package token_repo

import (
	"context"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/jwttoken"
	"license-manager/pkg/repositories/ent-fw/ent/user"
)

type jwtTokenEntRepo struct {
	client *ent.Client
}

func NewJwtTokenEntRepo(client *ent.Client) repositories.JwtTokenRepository {
	return &jwtTokenEntRepo {
		client: client,
	}
}

func (repo *jwtTokenEntRepo) Save(token domain.Token) (domain.Token, error) {
	t, err := repo.client.JwtToken.Create().
		SetToken(token.Value).
		SetRevoked(token.Revoked).
		SetClaims(token.Claims).
		SetIssuerID(token.IssuerID).
		Save(context.TODO())
	if err != nil {
		return domain.Token{}, err
	}
	
	return toEntity(t), nil
}

func (repo *jwtTokenEntRepo) FindByToken(tokenValue string) (domain.Token, error) {
	token, err := repo.client.JwtToken.Query().
		Where(
			jwttoken.TokenEQ(tokenValue),
		).
		Only(context.TODO())
	if err != nil {
		return domain.Token{}, err
	}
	return toEntity(token), nil
}

func (repo *jwtTokenEntRepo) FindByIssuer(userID string) ([]domain.Token, error) {
	
	tokens, err := repo.client.JwtToken.
		Query().
		Where(jwttoken.HasIssuerWith(
			user.IDEQ(userID),
		)).
		All(context.TODO())
	if err != nil {
		return nil, err
	}

	return toEntitySlice(tokens), nil
}

func (repo *jwtTokenEntRepo) FindClaimsByToken(tokenValue string) (domain.Claims, error) {
	token, err := repo.client.JwtToken.Query().
		Where(
			jwttoken.TokenEQ(tokenValue),
		).
		Only(context.TODO())
	if err != nil {
		return domain.Claims{}, err
	}
	return domain.Claims(token.Claims), nil
}

func (repo *jwtTokenEntRepo) IsRevoked(tokenValue string) (bool, error) {
	tok, err := repo.FindByToken(tokenValue)
	return tok.Revoked, err
}

func (repo *jwtTokenEntRepo) Delete(tokenValue string) error {
	deleted, err := repo.client.JwtToken.Delete().
		Where(
			jwttoken.TokenEQ(tokenValue),
		).
		Exec(context.TODO())
	if err != nil {
		return err
	}
	if deleted == 0 {
		return &ent.NotFoundError{}
	}
	return nil
}

func (repo *jwtTokenEntRepo) DeleteAllByIssuer(userID string) (int, error) {
	deleted, err := repo.client.JwtToken.Delete().
		Where(
			jwttoken.IssuerIDEQ(userID),
		).
		Exec(context.TODO())
	if err != nil {
		return 0, err
	}
	return deleted, nil
}


func toEntity(dto *ent.JwtToken) domain.Token {
	return domain.Token {
		Value: dto.Token,
		Revoked: dto.Revoked,
		Claims: dto.Claims,
		IssuerID: dto.IssuerID,
	}
}

func toEntitySlice(dtos []*ent.JwtToken) []domain.Token {
	
	tokens := make([]domain.Token, len(dtos))

	for i, dto := range dtos {
		tokens[i] = toEntity(dto)
	}

	return tokens
}