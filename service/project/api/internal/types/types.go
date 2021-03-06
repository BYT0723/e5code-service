// Code generated by goctl. DO NOT EDIT.
package types

type Project struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Url     string `json:"url"`
	Status  string `json:"status"`
	OwnerID string `json:"owner_id"`
	Owner   User   `json:"owner"`
}

type User struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Account string `json:"account"`
	Name    string `json:"name"`
}

type CommitInfo struct {
	Message string `json:"message"`
	Author  string `json:"author"`
	When    string `json:"when"`
}

type File struct {
	Name       string     `json:"name"`
	IsFile     bool       `json:"isFile"`
	Mode       string     `json:"mode"`
	Size       int64      `json:"size"`
	Path       string     `json:"path"`
	CommitHasg string     `json:"commitHash"`
	CommitInfo CommitInfo `json:"commitInfo"`
	Children   []File     `json:"children"`
}

type GetProjectReq struct {
	ID string `json:"id"`
}

type GetProjectReply struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Url     string `json:"url"`
	OwnerID string `json:"owner_id"`
}

type GetProjectAuthReply struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type AddProjectReq struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Url      string `json:"url"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type AddProjectReply struct {
	ID string `json:"id"`
}

type UpdateProjectReq struct {
	ID   string `json:"id"`
	Name string `json:"name,optional"`
	Desc string `json:"desc,optional"`
	Url  string `json:"url,optional"`
}

type UpdateProjectReply struct {
	Result bool `json:"result"`
}

type DeleteProjectReq struct {
	ID string `json:"id"`
}

type DeleteProjectReply struct {
	Result bool `json:"result"`
}

type AddUserReq struct {
	UserID    string `json:"user_id"`
	ProjectID string `json:"project_id"`
}

type AddUserReply struct {
	Result bool `json:"result"`
}

type RemoveUserReq struct {
	UserID    string `json:"user_id"`
	ProjectID string `json:"project_id"`
}

type RemoveUserReply struct {
	Result bool `json:"result"`
}

type ModifyPermissionReq struct {
	UserID    string `json:"user_id"`
	ProjectID string `json:"project_id"`
	Value     int64  `json:"value"`
}

type ModifyPermissionReply struct {
	Result bool `json:"result"`
}

type ListProjectReq struct {
	UserID string `json:"user_id"`
}

type ListProjectReply struct {
	Count  int64     `json:"count"`
	Result []Project `json:"result"`
}

type ListProjectFilesReq struct {
	ID     string `json:"id"`
	Path   string `json:"path"`
	IsWork bool   `json:"isWork"`
}

type ListProjectFilesReply struct {
	Count  int64  `json:"count"`
	Result []File `json:"result"`
}

type ListProjectAllFilesReq struct {
	ID     string `json:"id"`
	IsWork bool   `json:"isWork"`
}

type ListProjectAllFilesReply struct {
	Count  int64  `json:"count"`
	Result []File `json:"result"`
}

type CreateFileReq struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

type CreateFileReply struct {
	Result bool `json:"result"`
}

type ReadFileReq struct {
	ID     string `json:"id"`
	Path   string `json:"path"`
	IsWork bool   `json:"isWork"`
}

type ReadFileReply struct {
	Body string `json:"body"`
}

type UpdateFileReq struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Body string `json:"body"`
}

type UpdateFileReply struct {
	Result bool `json:"result"`
}

type DeleteFileReq struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

type DeleteFileReply struct {
	Result bool `json:"result"`
}

type MoveFileReq struct {
	ID      string `json:"id"`
	Oldpath string `json:"oldpath"`
	Newpath string `json:"newpath"`
}

type MoveFileReply struct {
	Result bool `json:"result"`
}

type MkDirReq struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

type MkDirReply struct {
	Result bool `json:"result"`
}

type FileStatus struct {
	Path   string `json:"path"`
	Status string `json:"status"`
}

type WorkStatusReq struct {
	ID string `json:"id"`
}

type WorkStatusReply struct {
	Result []FileStatus `json:"result"`
}

type CommitReq struct {
	ID        string   `json:"id"`
	Msg       string   `json:"msg"`
	FilePaths []string `json:"filePaths"`
}

type CommitReply struct {
	Result bool `json:"result"`
}
