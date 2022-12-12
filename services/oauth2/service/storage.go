package service

import (
	"context"
	"github.com/ory/fosite"
	"gopkg.in/square/go-jose.v2"
	"oauth/persistence"
	"sync"
	"time"
)

type StoreAuthorizeCode struct {
	active bool
	fosite.Requester
}

type StoreRefreshToken struct {
	active bool
	fosite.Requester
}

type IssuerPublicKeys struct {
	Issuer    string
	KeysBySub map[string]SubjectPublicKeys
}

type SubjectPublicKeys struct {
	Subject string
	Keys    map[string]PublicKeyScopes
}

type PublicKeyScopes struct {
	Key    *jose.JSONWebKey
	Scopes []string
}

type Storage struct {
	AuthorizeCodes  map[string]StoreAuthorizeCode
	IDSessions      map[string]fosite.Requester
	AccessTokens    map[string]fosite.Requester
	RefreshTokens   map[string]StoreRefreshToken
	PKCES           map[string]fosite.Requester
	BlacklistedJTIs map[string]time.Time
	// In-memory request ID to token signatures
	AccessTokenRequestIDs  map[string]string
	RefreshTokenRequestIDs map[string]string
	// Public keys to check signature in auth grant jwt assertion.
	IssuerPublicKeys map[string]IssuerPublicKeys

	authorizeCodesMutex         sync.RWMutex
	idSessionsMutex             sync.RWMutex
	accessTokensMutex           sync.RWMutex
	refreshTokensMutex          sync.RWMutex
	pkcesMutex                  sync.RWMutex
	usersMutex                  sync.RWMutex
	blacklistedJTIsMutex        sync.RWMutex
	accessTokenRequestIDsMutex  sync.RWMutex
	refreshTokenRequestIDsMutex sync.RWMutex
	issuerPublicKeysMutex       sync.RWMutex
}

func (s Storage) GetPKCERequestSession(ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) CreatePKCERequestSession(ctx context.Context, signature string, requester fosite.Requester) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) DeletePKCERequestSession(ctx context.Context, signature string) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetPublicKey(ctx context.Context, issuer string, subject string, keyId string) (*jose.JSONWebKey, error) {
	s.issuerPublicKeysMutex.RLock()
	defer s.issuerPublicKeysMutex.RUnlock()

	if issuerKeys, ok := s.IssuerPublicKeys[issuer]; ok {
		if subKeys, ok := issuerKeys.KeysBySub[subject]; ok {
			if keyScopes, ok := subKeys.Keys[keyId]; ok {
				return keyScopes.Key, nil
			}
		}
	}

	return nil, fosite.ErrNotFound
}

func (s Storage) GetPublicKeys(ctx context.Context, issuer string, subject string) (*jose.JSONWebKeySet, error) {
	s.issuerPublicKeysMutex.RLock()
	defer s.issuerPublicKeysMutex.RUnlock()

	if issuerKeys, ok := s.IssuerPublicKeys[issuer]; ok {
		if subKeys, ok := issuerKeys.KeysBySub[subject]; ok {
			if len(subKeys.Keys) == 0 {
				return nil, fosite.ErrNotFound
			}

			keys := make([]jose.JSONWebKey, 0, len(subKeys.Keys))
			for _, keyScopes := range subKeys.Keys {
				keys = append(keys, *keyScopes.Key)
			}

			return &jose.JSONWebKeySet{Keys: keys}, nil
		}
	}

	return nil, fosite.ErrNotFound
}

func (s Storage) GetPublicKeyScopes(ctx context.Context, issuer string, subject string, keyId string) ([]string, error) {
	s.issuerPublicKeysMutex.RLock()
	defer s.issuerPublicKeysMutex.RUnlock()

	if issuerKeys, ok := s.IssuerPublicKeys[issuer]; ok {
		if subKeys, ok := issuerKeys.KeysBySub[subject]; ok {
			if keyScopes, ok := subKeys.Keys[keyId]; ok {
				return keyScopes.Scopes, nil
			}
		}
	}

	return nil, fosite.ErrNotFound
}

func (s Storage) IsJWTUsed(ctx context.Context, jti string) (bool, error) {
	err := s.ClientAssertionJWTValid(ctx, jti)
	if err != nil {
		return true, nil
	}

	return false, nil
}

func (s Storage) MarkJWTUsedForTime(ctx context.Context, jti string, exp time.Time) error {
	return s.SetClientAssertionJWT(ctx, jti, exp)
}

func (s Storage) Authenticate(ctx context.Context, name string, secret string) error {
	return persistence.User().CheckUser(ctx, name, secret)
}

func NewStorage() *Storage {
	return &Storage{
		IDSessions:             make(map[string]fosite.Requester),
		AuthorizeCodes:         map[string]StoreAuthorizeCode{},
		AccessTokens:           map[string]fosite.Requester{},
		RefreshTokens:          map[string]StoreRefreshToken{},
		PKCES:                  map[string]fosite.Requester{},
		AccessTokenRequestIDs:  map[string]string{},
		RefreshTokenRequestIDs: map[string]string{},
		IssuerPublicKeys:       map[string]IssuerPublicKeys{},
	}
}

func (s Storage) CreateOpenIDConnectSession(ctx context.Context, authorizeCode string, requester fosite.Requester) error {
	s.IDSessions[authorizeCode] = requester
	return nil
}

func (s Storage) GetOpenIDConnectSession(ctx context.Context, authorizeCode string, requester fosite.Requester) (fosite.Requester, error) {
	s.idSessionsMutex.RLock()
	defer s.idSessionsMutex.RUnlock()

	cl, ok := s.IDSessions[authorizeCode]
	if !ok {
		return nil, fosite.ErrNotFound
	}
	return cl, nil
}

func (s Storage) DeleteOpenIDConnectSession(ctx context.Context, authorizeCode string) error {
	s.idSessionsMutex.Lock()
	defer s.idSessionsMutex.Unlock()

	delete(s.IDSessions, authorizeCode)
	return nil
}

func (s Storage) RevokeRefreshToken(ctx context.Context, requestID string) error {
	s.refreshTokenRequestIDsMutex.Lock()
	defer s.refreshTokenRequestIDsMutex.Unlock()

	if signature, exists := s.RefreshTokenRequestIDs[requestID]; exists {
		rel, ok := s.RefreshTokens[signature]
		if !ok {
			return fosite.ErrNotFound
		}
		rel.active = false
		s.RefreshTokens[signature] = rel
	}
	return nil
}

func (s Storage) RevokeRefreshTokenMaybeGracePeriod(ctx context.Context, requestID string, signature string) error {
	return s.RevokeRefreshToken(ctx, requestID)
}

func (s Storage) RevokeAccessToken(ctx context.Context, requestID string) error {
	s.accessTokenRequestIDsMutex.RLock()
	defer s.accessTokenRequestIDsMutex.RUnlock()

	if signature, exists := s.AccessTokenRequestIDs[requestID]; exists {
		if err := s.DeleteAccessTokenSession(ctx, signature); err != nil {
			return err
		}
	}
	return nil
}

func (s Storage) GetClient(ctx context.Context, id string) (fosite.Client, error) {
	clientCommand := ClientCommand{ID: id}
	cl, err := clientCommand.FindClientById(ctx)
	if err != nil {
		return nil, fosite.ErrNotFound
	}
	return cl, nil
}

func (s Storage) ClientAssertionJWTValid(ctx context.Context, jti string) error {
	s.blacklistedJTIsMutex.RLock()
	defer s.blacklistedJTIsMutex.RUnlock()

	if exp, exists := s.BlacklistedJTIs[jti]; exists && exp.After(time.Now()) {
		return fosite.ErrJTIKnown
	}

	return nil
}

func (s Storage) SetClientAssertionJWT(ctx context.Context, jti string, exp time.Time) error {
	s.blacklistedJTIsMutex.Lock()
	defer s.blacklistedJTIsMutex.Unlock()

	// delete expired jtis
	for j, e := range s.BlacklistedJTIs {
		if e.Before(time.Now()) {
			delete(s.BlacklistedJTIs, j)
		}
	}

	if _, exists := s.BlacklistedJTIs[jti]; exists {
		return fosite.ErrJTIKnown
	}

	s.BlacklistedJTIs[jti] = exp
	return nil
}

func (s Storage) CreateAuthorizeCodeSession(ctx context.Context, code string, request fosite.Requester) (err error) {
	s.authorizeCodesMutex.Lock()
	defer s.authorizeCodesMutex.Unlock()

	s.AuthorizeCodes[code] = StoreAuthorizeCode{active: true, Requester: request}
	return nil
}

func (s Storage) GetAuthorizeCodeSession(ctx context.Context, code string, session fosite.Session) (request fosite.Requester, err error) {
	s.authorizeCodesMutex.RLock()
	defer s.authorizeCodesMutex.RUnlock()

	rel, ok := s.AuthorizeCodes[code]
	if !ok {
		return nil, fosite.ErrNotFound
	}
	if !rel.active {
		return rel, fosite.ErrInvalidatedAuthorizeCode
	}

	return rel.Requester, nil
}

func (s Storage) InvalidateAuthorizeCodeSession(ctx context.Context, code string) (err error) {
	s.authorizeCodesMutex.Lock()
	defer s.authorizeCodesMutex.Unlock()

	rel, ok := s.AuthorizeCodes[code]
	if !ok {
		return fosite.ErrNotFound
	}
	rel.active = false
	s.AuthorizeCodes[code] = rel
	return nil
}

func (s Storage) CreateRefreshTokenSession(ctx context.Context, signature string, request fosite.Requester) (err error) {
	s.refreshTokenRequestIDsMutex.Lock()
	defer s.refreshTokenRequestIDsMutex.Unlock()
	s.refreshTokensMutex.Lock()
	defer s.refreshTokensMutex.Unlock()

	s.RefreshTokens[signature] = StoreRefreshToken{active: true, Requester: request}
	s.RefreshTokenRequestIDs[request.GetID()] = signature
	return nil
}

func (s Storage) GetRefreshTokenSession(ctx context.Context, signature string, session fosite.Session) (request fosite.Requester, err error) {
	s.refreshTokensMutex.RLock()
	defer s.refreshTokensMutex.RUnlock()

	rel, ok := s.RefreshTokens[signature]
	if !ok {
		return nil, fosite.ErrNotFound
	}
	if !rel.active {
		return rel, fosite.ErrInactiveToken
	}
	return rel, nil
}

func (s Storage) DeleteRefreshTokenSession(ctx context.Context, signature string) (err error) {
	s.refreshTokensMutex.Lock()
	defer s.refreshTokensMutex.Unlock()

	delete(s.RefreshTokens, signature)
	return nil
}

func (s Storage) CreateAccessTokenSession(ctx context.Context, signature string, request fosite.Requester) (err error) {
	s.accessTokenRequestIDsMutex.Lock()
	defer s.accessTokenRequestIDsMutex.Unlock()
	s.accessTokensMutex.Lock()
	defer s.accessTokensMutex.Unlock()

	s.AccessTokens[signature] = request
	s.AccessTokenRequestIDs[request.GetID()] = signature
	return nil
}

func (s Storage) GetAccessTokenSession(ctx context.Context, signature string, session fosite.Session) (request fosite.Requester, err error) {
	s.accessTokensMutex.RLock()
	defer s.accessTokensMutex.RUnlock()

	rel, ok := s.AccessTokens[signature]
	if !ok {
		return nil, fosite.ErrNotFound
	}
	return rel, nil
}

func (s Storage) DeleteAccessTokenSession(ctx context.Context, signature string) (err error) {
	s.accessTokensMutex.Lock()
	defer s.accessTokensMutex.Unlock()

	delete(s.AccessTokens, signature)
	return nil
}
