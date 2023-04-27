"""
this script is used to prepare releases.

it builds fleck in full featured and bare builds for all supported operating systems and architectures

replaces makefile due to the fact, that it isn't operating system independent

requires a build.conf file including the following info:
    VERSION:x.x.x-x
    FEATURE:x.x
    FLAGS:x
"""
import os
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


def build_for_arch(bare: bool, arch: str, _os: str, flags: str, version: str, variables: Dict[str, str]):
    """
    dispatches a build command to the `go build` toolchain
    """
    for v in variables:
        val = variables[v]
        if v == "VERSION":
            val = "bare:"+val
        flags += f" -X main.{v}={val}"

    bTag = ""
    b = "_"
    if bare:
        bTag = "-tags=bare"
        b = "-bare_"

    cmd = f'go build {bTag} -ldflags="{flags}" -o ./out/fleck{b}{version}_{_os}_{arch}'
    subprocess.Popen(cmd, shell=True, env={
        **os.environ, 'CGO_ENABLED': '0', 'GOOS': _os, 'GOARCH': arch}
    )


def run():
    os.makedirs("out", exist_ok=True)
    print("created out dir: './out'")
    r = 0
    for a in arch:
        r += len(arch[a]) * 2

    print(
        f"I: detected {r} architecture operating system combinations, preparing build...")

    conf = get_config()
    print("I: read config from build.conf: \n", conf)

    variables = {
        "VERSION": f"{conf['VERSION']}+{conf['FEATURE']}",
        "BUILD_AT": datetime.datetime.now().isoformat(),
        "BUILD_BY": f"{run_cmd('git config --global user.name')}-{run_cmd('git config --global user.email')}".replace("\n", ""),
    }
    print("I: prepared variables: \n", variables)

    t = 0
    for a in arch:
        print("="*30)
        print(f"building for {a}")
        for o in arch[a]:
            print(f"building {t}/{r} [{conf['VERSION']}_{o}_{a}]")
            build_for_arch(
                False, o, a, conf["FLAGS"], conf["VERSION"], variables)
            t += 1
            print(f"building {t}/{r} [bare:{conf['VERSION']}_{o}_{a}]")
            build_for_arch(
                True, o, a, conf["FLAGS"], conf["VERSION"], variables)
            t += 1
    print("v"*30)
    print("done...")


run()
