package gitx

import (
	"e5Code-Service/common/sshx"
	"fmt"
)

type Cli struct {
  sshCli *sshx.Cli
}

func NewCliWithOpt(addr, user, pwd string) *Cli {
  return &Cli{
    sshx.NewCli(addr, user, pwd),
  }
}

func NewCli() *Cli {
  return &Cli{
    sshx.NewCli("git.byt0723.xyz:22", "git", "wangtao"),
  }
}

func (c *Cli) CreateUser(name string) (string, error) {
  // 添加Git下目录
  addDir := fmt.Sprintf("mkdir ./%s", name)
  // 添加Apache验证
  // addApache := fmt.Sprintf("htpasswd -mb /etc/httpd/conf.d/git-team.htpasswd %s %s", name, password)
  return c.sshCli.Run(addDir)
}

func (c *Cli) DestoryUser(name string) (string, error) {
  // 移除Git下目录
  rmDir := fmt.Sprintf("rm -rf ./%s", name)
  // 移除Apache验证
  // rmApache := fmt.Sprintf("htpasswd -D /etc/httpd/conf.d/git-team.htpasswd %s", name)
  return c.sshCli.Run(rmDir)
}

func (c *Cli) CreateRegistry(name, registry string) (string, error) {
  cmd := fmt.Sprintf("cd %s && git init --bare %s.git", name, registry)
  return c.sshCli.Run(cmd)
}

func (c *Cli) ForkRegistry(name, registry, url string) (string, error)  {
  cmd := fmt.Sprintf("cd %s && git clone --bare %s %s.git ", name,url,registry)
  return c.sshCli.Run(cmd)
}

func (c *Cli) DestoryRegistry(name, registry string) (string, error) {
  cmd := fmt.Sprintf("rm %s/%s", name, registry)
  return c.sshCli.Run(cmd)
}

func (c *Cli) AddSSHKey(key string) (string, error) {
  cmd := fmt.Sprintf("echo '%s' >> $HOME/.ssh/authorized_keys", key)
  return c.sshCli.Run(cmd)
}

func (c *Cli) DeleteSSHKey(key string) (string, error) {
  cmd := fmt.Sprintf("sed -ie \"/%s/d\" $HOME/.ssh/authorized_keys", key)
  return c.sshCli.Run(cmd)
}
