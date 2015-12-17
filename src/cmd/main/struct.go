package main

type GogsHookRequest struct {
	Secret      string     `json:"secret"`
	Ref         string     `json:"ref"`
	Before      string     `json:"before"`
	After       string     `json:"after"`
	Compare_url string     `json:"compare_url"`
	Commits     []Commit   `json:"commits"`
	Repository  Repository `json:"repository"`
	Pusher      Author     `json:"pusher"`
	Sender      Sender     `json:"sender"`
}

type Sender struct {
	Login     string `json:"login"`
	Id        int64  `json:"id"`
	AvatarUrl int64  `json:"avatar_url"`
}

type Repository struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	SshUrl      string `json:"ssh_url"`
	CloneUrl    string `json:"clone_url"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Watchers    int64  `json:"watchers"`
	Owner       Author `json:"owner"`
	Private     bool   `json:"private"`
}

type Commit struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Url     string `json:"url"`
	Author  Author `json:"author"`
}

type Author struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
