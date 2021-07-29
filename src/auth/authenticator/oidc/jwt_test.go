package oidc

import (
	"context"
	"crypto/x509"
	"strings"
	"testing"
	"time"

	"github.com/consensys/quorum-key-manager/pkg/jwt"
	"github.com/consensys/quorum-key-manager/pkg/tls/certificate"
	"github.com/consensys/quorum-key-manager/pkg/tls/testutils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJWTChecker_RSAToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	claimsCfg := &ClaimsConfig{
		Username: "test.username",
		Group:    "test.groups",
	}

	cert, _ := certificate.X509KeyPair([]byte(testutils.RSACertPEM), []byte(testutils.RSAKeyPEM))
	checker := NewJWTChecker([]*x509.Certificate{cert.Leaf}, claimsCfg, false)
	generator, _ := jwt.NewTokenGenerator(cert.PrivateKey)

	t.Run("should accept token and extract claims successfully", func(t *testing.T) {
		username := "username1"
		groups := []string{"group1", "group2"}
		token, _ := generator.GenerateAccessToken(map[string]interface{}{
			"test.username": username,
			"test.groups":   strings.Join(groups, ","),
		}, time.Second)
		claims, err := checker.Check(ctx, token)
		require.NoError(t, err)
		assert.Equal(t, username, claims.Username)
		assert.Equal(t, groups, claims.Groups)
	})

	t.Run("should accept token and only username claims successfully", func(t *testing.T) {
		username := "username2"
		token, _ := generator.GenerateAccessToken(map[string]interface{}{
			"test.username": username,
		}, time.Second)
		claims, err := checker.Check(ctx, token)
		require.NoError(t, err)
		assert.Equal(t, username, claims.Username)
		assert.Empty(t, claims.Groups)
	})

	t.Run("should accept token and only groups claims successfully", func(t *testing.T) {
		groups := []string{"group1", "group2"}
		token, _ := generator.GenerateAccessToken(map[string]interface{}{
			"test.groups": strings.Join(groups, ","),
		}, time.Second)
		claims, err := checker.Check(ctx, token)
		require.NoError(t, err)
		assert.Equal(t, groups, claims.Groups)
		assert.Empty(t, claims.Username)
	})

	t.Run("should reject invalid token successfully", func(t *testing.T) {
		token := "invalid-token"
		claims, err := checker.Check(ctx, token)
		assert.Error(t, err)
		assert.Empty(t, claims)
	})
}