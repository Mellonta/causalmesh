import argparse
import subprocess
from common import CLOUD, SERVERS, WS_DIR

def run_collect_output(cmd):
    res = subprocess.run(cmd, stdout=subprocess.PIPE)
    return res.stdout.decode('utf-8').strip()

def run_shell(cmd):
    res = subprocess.run(cmd, stdout=subprocess.PIPE, shell=True)
    return res.stdout.decode('utf-8').strip()

def cloudlab():
    for s in SERVERS:
        print("copying files to ", s)
        path = f'Haoran@{s}:~/'
        run_collect_output(['rsync', '--exclude', 'target', '--exclude', 'figures', '-r', f'{WS_DIR}/ccmesh', path])
        run_collect_output(['rsync', '-r', f'{WS_DIR}/ccmesh-go', path])

def aws_bak():
    for s in SERVERS:
        print("copying files to ", s)
        path = f'ubuntu@{s}:~/'
        run_shell(f'rsync -e "ssh -i /home/cavriola/.ssh/zhanghaorancs.pem" --exclude target -r {WS_DIR}/ccmesh {path}')
        run_shell(f'rsync -e "ssh -i /home/cavriola/.ssh/zhanghaorancs.pem" -r {WS_DIR}/ccmesh-go {path}')

def aws():
    for s in SERVERS:
        print("copying files to ", s)
        path = f'ubuntu@{s}:~/'
        run_shell(f'rsync --exclude target -r {WS_DIR}/ccmesh {path}')
        run_shell(f'rsync -r {WS_DIR}/ccmesh-go {path}')

def main():
    if CLOUD == "cloudlab":
        cloudlab()
    elif CLOUD == "aws":
        aws()
    else:
        exit(1)


if __name__ == '__main__':
    main()
