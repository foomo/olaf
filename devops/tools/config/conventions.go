package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var errExtractNameAndGroupNameIsEmpty = errors.New("name is empty")
var errExtractNameAndGroupGroupIsEmpty = errors.New("group is empty")
var errExtractNameAndGroupTooManyParts = errors.New("too many name parts separated by _ there be only 2 parts max")

func extractNameAndGroup(n string) (name string, group GroupName, err error) {
	nameParts := strings.Split(n, "_")
	switch len(nameParts) {
	case 1:
		name = nameParts[0]
		if name == "" {
			return "", "", errExtractNameAndGroupNameIsEmpty
		}
		return name, GroupName(name), nil
	case 2:
		group = GroupName(nameParts[0])
		name = nameParts[1]
		if name == "" {
			return "", "", errExtractNameAndGroupNameIsEmpty
		}
		if group == "" {
			return "", "", errExtractNameAndGroupGroupIsEmpty
		}
		return name, group, nil
	default:
		return "", "", errExtractNameAndGroupTooManyParts
	}

}

func collectFrontends(path string, rw io.Writer) (frontends []Frontend, err error) {
	if rw == nil {
		rw = os.Stderr
	}
	path = filepath.Join(path, "frontend")
	fmt.Fprintln(rw, "scanning for frontends in", path)
	files, errReadDir := ioutil.ReadDir(path)
	if errReadDir != nil {
		return nil, errReadDir
	}

	type packagejson struct {
		Scripts map[string]string
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Fprintln(rw, "	descending into", f.Name())

			frontendPath := filepath.Join(path, f.Name())
			jsonPackagePath := filepath.Join(frontendPath, "package.json")
			jsonBytes, errRead := ioutil.ReadFile(jsonPackagePath)
			if errRead == nil {
				p := &packagejson{}
				errUnmarshal := json.Unmarshal(jsonBytes, &p)
				if errUnmarshal != nil {
					return nil, errors.New("could not parse " + jsonPackagePath + " : " + errUnmarshal.Error())
				}
				_, okBuild := p.Scripts["build"]
				_, okStart := p.Scripts["start"]

				if !okBuild || !okStart {
					fmt.Fprintln(rw, "		did not find build and start scripts, skipping this folder", p.Scripts)
					continue
				}

				pathDockerfile := filepath.Join(frontendPath, "Dockerfile")

				info, errDockerfileStat := os.Stat(pathDockerfile)
				if errDockerfileStat != nil || info.IsDir() {
					pathDockerfile = ""
				}

				name, group, errExtractNameAndGroup := extractNameAndGroup(f.Name())
				if errExtractNameAndGroup != nil {
					fmt.Fprintln(rw, "		could not extract name and group from", f.Name(), errExtractNameAndGroup)
					return nil, errExtractNameAndGroup
				}

				frontend := Frontend{
					Group: group,
					Build: Build{
						Name:       name,
						Path:       frontendPath,
						Dockerfile: pathDockerfile,
					},
				}
				frontends = append(frontends, frontend)
				fmt.Fprintln(rw, "		found a valid frontend", frontend)

			}
		}
	}
	return
}

func collectServices(path string, rw io.Writer) (services []Service, err error) {
	if rw == nil {
		rw = os.Stderr
	}

	path = filepath.Join(path, "backend", "cmd")

	fmt.Fprintln(rw, "scanning for services in", path)
	files, errReadDir := ioutil.ReadDir(path)
	if errReadDir != nil {
		return nil, errReadDir
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Fprintln(rw, "	descending into", f.Name())

			servicePath := filepath.Join(path, f.Name())

			programPath := filepath.Join(servicePath, f.Name()+".go")

			programInfo, errStat := os.Stat(programPath)
			if errStat != nil || programInfo.IsDir() {
				fmt.Fprintln(rw, "		expected ", programPath, " not found")
				continue
			}

			name, group, errExtractNameAndGroup := extractNameAndGroup(f.Name())
			if errExtractNameAndGroup != nil {
				fmt.Fprintln(rw, "		could not extract name and group from", f.Name(), errExtractNameAndGroup)
				return nil, errExtractNameAndGroup
			}
			service := Service{
				Group: group,
				Build: Build{
					Name: name,
				},
			}
			fmt.Fprintln(rw, "			found a valid service", service)
			services = append(services, service)

		}
	}
	return
}

func Collect(path string, w io.Writer) (groups Groups, err error) {
	if w == nil {
		w = os.Stderr
	}
	fmt.Fprintln(w, "collecting groups in:", path)
	// collect frontends
	frontends, errCollectFrontends := collectFrontends(path, w)
	if errCollectFrontends != nil {
		fmt.Fprintln(w, "	failed to collect frontends", errCollectFrontends)
		return nil, errCollectFrontends
	}
	// collect services
	services, errCollectServices := collectServices(path, w)
	if errCollectFrontends != nil {
		fmt.Fprintln(w, "	failed to collect services", errCollectServices)
		return nil, errCollectServices
	}
	// merge them into groups
	groups = Groups{}
	for _, service := range services {
		group, groupOK := groups[service.Group]
		if !groupOK {
			group = &Group{
				Name: service.Group,
			}
			groups[service.Group] = group
		}
		group.Services = append(group.Services, service)
	}
	for _, frontend := range frontends {
		group, groupOK := groups[frontend.Group]
		if !groupOK {
			group = &Group{Name: frontend.Group}
			groups[frontend.Group] = group
		}
		group.Frontends = append(group.Frontends, frontend)
	}
	return groups, nil
}
