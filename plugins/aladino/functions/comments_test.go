// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package plugins_aladino_functions_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/google/go-github/v42/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/reviewpad/reviewpad/v2/lang/aladino"
	mocks_aladino "github.com/reviewpad/reviewpad/v2/mocks/aladino"
	plugins_aladino "github.com/reviewpad/reviewpad/v2/plugins/aladino"
	"github.com/stretchr/testify/assert"
)

var comments = plugins_aladino.PluginBuiltIns().Functions["comments"].Code

func TestComments(t *testing.T) {
	wantedComments := aladino.BuildArrayValue(
		[]aladino.Value{
			aladino.BuildStringValue("hello world"),
		},
	)

	mockedEnv, err := mocks_aladino.MockDefaultEnv(
		mock.WithRequestMatch(
			mock.GetReposIssuesCommentsByOwnerByRepoByIssueNumber,
			[]*github.IssueComment{
				{
					Body: github.String("hello world"),
				},
			},
		),
	)

	if err != nil {
		log.Fatalf("mockDefaultEnv failed: %v", err)
	}

	args := []aladino.Value{}
	gotComments, err := comments(mockedEnv, args)

	assert.Nil(t, err)
	assert.Equal(t, wantedComments, gotComments)
}

func TestComments_WhenGetCommentsRequestFailed(t *testing.T) {
	failMessage := "GetCommentsRequestFailed"

	mockedEnv, err := mocks_aladino.MockDefaultEnv(
		mock.WithRequestMatchHandler(
			mock.GetReposIssuesCommentsByOwnerByRepoByIssueNumber,
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mock.WriteError(
					w,
					http.StatusInternalServerError,
					failMessage,
				)
			}),
		),
	)

	if err != nil {
		log.Fatalf("mockDefaultEnv failed: %v", err)
	}

	args := []aladino.Value{}
	gotComments, err := comments(mockedEnv, args)

	assert.Nil(t, gotComments)
	assert.Equal(t, err.(*github.ErrorResponse).Message, failMessage)
}