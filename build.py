"""
this script is used to prepare releases.

it builds fleck in full featured and bare builds for all supported operating systems and architectures

replaces makefile due to the fact, that it isn't operating system independent

requires a build.conf file including the following info:
    VERSION:x.x.x-x
    FEATURE:x.x
    FLAGS:x
"""
from typing import Dict, List
import datetime
import subprocess

arch = {
    "linux": [
        "386",
        "amd64",
        "arm",
        "arm64",
    ],
    "windows": [
        "386",
        "amd64",
        "arm",
        "arm64",
    ],
    "darwin": [
        "amd64",
        "arm64",
    ]
}


def run_cmd(s: str) -> str:
    """
    runs a command and returns its utf-8 decoded result
    """
    return subprocess.run(s.split(" "), stdout=subprocess.PIPE).stdout.decode("utf-8")


def get_config() -> Dict[str, str]:
    """
    parses the build config, returns it as a dict
    """
    l: List[str] = []
    r: Dict[str, str] = {}
    with open("build.conf", "r", encoding="utf-8") as f:
        l = f.readlines()
    for _ in l:
        if ':' not in _:
            continue
        k, v = _.split(":")
        r[k.replace("\n", "")] = v.replace("\n", "")
    return r

def build_for_arch(arch: str, os: str, flags: str, version: str, variables: Dict[str, str]):
    """
    dispatches a build command to the `go build` toolchain
    """
    for v in variables:
        flags += f" -X main.{v}={variables[v]}"
    cmd = f"go build -ldflags='{flags}' -o fleck_{version}_{os}_{arch}"
    env = f"CGO_ENABLED=0 GOOS={os} GOARCH={arch}"
    print(env, cmd)
    # TODO: set env variables
    # TODO: run command

def run():
    r = 0
    for a in arch:
        r += len(arch[a])

    print(f"I: detected {r} architecture operating system combinations, preparing build...")

    conf = get_config()
    print("I: read config from build.conf: \n",conf)

    variables = {
        "VERSION": f"{conf['VERSION']}+{conf['FEATURE']}",
        "BUILD_AT": datetime.datetime.now().isoformat(),
        "BUILD_BY": f"{run_cmd('git config --global user.name')}-{run_cmd('git config --global user.email')}".replace("\n", ""),
    }
    print("I: prepared variables: \n", variables)

    for a in arch:
        for o in arch[a]:
            build_for_arch(a, o, conf["FLAGS"], conf["VERSION"], variables)

run()
