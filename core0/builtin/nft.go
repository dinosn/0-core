package builtin

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"

	"github.com/zero-os/0-core/base/nft"
	"github.com/zero-os/0-core/base/pm"
	"github.com/zero-os/0-core/base/pm/core"
	"github.com/zero-os/0-core/base/pm/process"
)

type nftMgr struct {
	rules map[string]struct{}
	m     sync.RWMutex
}

func init() {
	b := &nftMgr{
		rules: make(map[string]struct{}),
	}
	pm.CmdMap["nft.open_port"] = process.NewInternalProcessFactory(b.openPort)
	pm.CmdMap["nft.drop_port"] = process.NewInternalProcessFactory(b.dropPort)
	pm.CmdMap["nft.list"] = process.NewInternalProcessFactory(b.listPorts)
	pm.CmdMap["nft.rule_exists"] = process.NewInternalProcessFactory(b.ruleExists)

}

type Port struct {
	Port      int    `json:"port"`
	Interface string `json:"interface,omitempty"`
	Subnet    string `json:"subnet,omitempty"`
}

func (b *nftMgr) parsePort(cmd *core.Command) (string, error) {
	var args Port
	if err := json.Unmarshal(*cmd.Arguments, &args); err != nil {
		return "", err
	}

	body := ""
	if args.Interface != "" {
		body += fmt.Sprintf(`iifname "%s" `, args.Interface)
	}
	if args.Subnet != "" {
		subnet := args.Subnet
		_, net, err := net.ParseCIDR(args.Subnet)
		if err == nil {
			subnet = net.String()
		}

		body += fmt.Sprintf(`ip saddr %s `, subnet)
	}

	body += fmt.Sprintf(`tcp dport %d accept`, args.Port)

	return body, nil
}

func (b *nftMgr) exists(rule string) bool {
	_, ok := b.rules[rule]
	return ok
}

func (b *nftMgr) register(rule string) error {
	if b.exists(rule) {
		return fmt.Errorf("exists")
	}

	b.rules[rule] = struct{}{}
	return nil
}

func (b *nftMgr) openPort(cmd *core.Command) (interface{}, error) {
	rule, err := b.parsePort(cmd)
	if err != nil {
		return nil, err
	}

	b.m.Lock()
	defer b.m.Unlock()

	if err := b.register(rule); err != nil {
		return nil, fmt.Errorf("rule exists")
	}

	n := nft.Nft{
		"filter": nft.Table{
			Family: nft.FamilyIP,
			Chains: nft.Chains{
				"input": nft.Chain{
					Rules: []nft.Rule{
						{Body: rule},
					},
				},
			},
		},
	}
	if err := nft.Apply(n); err != nil {
		delete(b.rules, rule)
		return nil, nft.Apply(n)
	}
	return nil, nil
}

func (b *nftMgr) dropPort(cmd *core.Command) (interface{}, error) {
	rule, err := b.parsePort(cmd)
	if err != nil {
		return nil, err
	}

	b.m.Lock()
	defer b.m.Unlock()

	if !b.exists(rule) {
		return nil, fmt.Errorf("rule does not exist")
	}

	n := nft.Nft{
		"filter": nft.Table{
			Family: nft.FamilyIP,
			Chains: nft.Chains{
				"input": nft.Chain{
					Rules: []nft.Rule{
						{Body: rule},
					},
				},
			},
		},
	}

	if err := nft.DropRules(n); err != nil {
		return nil, err
	}

	delete(b.rules, rule)
	return nil, nil
}

func (b *nftMgr) listPorts(cmd *core.Command) (interface{}, error) {
	b.m.RLock()
	defer b.m.RUnlock()

	ports := make([]string, 0, len(b.rules))
	for port, _ := range b.rules {
		ports = append(ports, port)
	}
	return ports, nil
}

func (b *nftMgr) ruleExists(cmd *core.Command) (interface{}, error) {
	rule, err := b.parsePort(cmd)
	if err != nil {
		return nil, err
	}

	b.m.RLock()
	defer b.m.RUnlock()

	return b.exists(rule), nil
}
