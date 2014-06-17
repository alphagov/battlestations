package github

import (
	"net/http"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

type User struct {
	Details github.User
	Teams   []github.Team
	Token   oauth.Token
}

type Service interface {
	AuthURL() string
	UserFromCode(string) (User, error)
}

type OAuthService struct {
	Organisation string
	config       *oauth.Config
}

func (s *OAuthService) AuthURL() string {
	return s.config.AuthCodeURL("")
}

func (s *OAuthService) UserFromCode(code string) (User, error) {
	var userStruct User
	var err error

	t := &oauth.Transport{Config: s.config}

	if _, err = t.Exchange(code); err != nil {
		return User{}, err
	}

	if userStruct, err = s.getUserStruct(t); err != nil {
		return User{}, err
	}

	return userStruct, nil
}

func (s *OAuthService) getUserStruct(oauthTransport *oauth.Transport) (User, error) {
	var user *github.User
	var isOrgMember bool
	var teams []github.Team
	var err error

	githubClient := github.NewClient(oauthTransport.Client())

	if user, _, err = githubClient.Users.Get(""); err != nil {
		return User{}, err
	}

	if isOrgMember, _, err = githubClient.Organizations.IsMember(s.Organisation, *user.Login); err != nil {
		return User{}, err
	}

	if isOrgMember {
		// get the users teams
		var request *http.Request

		if request, err = githubClient.NewRequest("GET", "https://api.github.com/user/teams", nil); err != nil {
			return User{}, err
		}

		if _, err = githubClient.Do(request, &teams); err != nil {
			return User{}, err
		}

		return User{*user, teams, *oauthTransport.Token}, nil
	} else {
		return User{}, &MembershipError{s.Organisation}
	}
}

type Config struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Organisation string `json:"organisation"`
}

func NewOAuthService(redirectURL string, config Config) *OAuthService {
	oauthConfig := &oauth.Config{
		ClientId:     config.ClientId,
		ClientSecret: config.ClientSecret,
		Scope:        "user",
		AuthURL:      "https://github.com/login/oauth/authorize",
		TokenURL:     "https://github.com/login/oauth/access_token",
		RedirectURL:  redirectURL,
	}
	return &OAuthService{config.Organisation, oauthConfig}
}
