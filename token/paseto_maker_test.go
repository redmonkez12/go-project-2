package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/redmonkez12/go-project-2/util"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidPasetoTokenAlgNone(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// 1. Test with a random string (not a token)
	payload, err := maker.VerifyToken("randomstring")
	require.Error(t, err)
	require.Nil(t, payload)

	// 2. Test with an empty token
	payload, err = maker.VerifyToken("")
	require.Error(t, err)
	require.Nil(t, payload)

	// 3. Test with a tampered token
	token, err := maker.CreateToken(util.RandomOwner(), time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Tamper the token
	tamperedToken := token + "tampered"
	payload, err = maker.VerifyToken(tamperedToken)
	require.Error(t, err)
	require.Nil(t, payload)

	// 4. Test with a token signed with a different key
	otherMaker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	otherToken, err := otherMaker.CreateToken(util.RandomOwner(), time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, otherToken)

	payload, err = maker.VerifyToken(otherToken)
	require.Error(t, err)
	require.Nil(t, payload)
}
