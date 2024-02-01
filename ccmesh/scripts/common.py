CLOUD = "cloudlab"
# CLOUD = "aws"
LOC = "home"
# LOC = "dsl"
# LOC = "mac"

NSERVERS = 14

if LOC == "home":
    WS_DIR = '/home/cavriola/WS'
    KEY_FILE = '/home/cavriola/.ssh/id_rsa'
elif LOC == "dsl":
    WS_DIR = '/home/haoran/WS'
    KEY_FILE = '/home/haoran/.ssh/id_ed25519'
else:
    WS_DIR = '/Users/cavriola/WS'
    KEY_FILE = '/Users/cavriola/.ssh/id_ed25519'

if CLOUD == "cloudlab":
    MSSERVERS = ["0942", 1009, "0820", "0929", 1130, 1040, "0919", 1104, "0917", "0908", "0941", "0939", "0903", 1010, 1036, 1018]
    MSSERVERS = MSSERVERS[:NSERVERS + 2]
    SERVERS = [f'ms{s}.utah.cloudlab.us' for s in MSSERVERS]
    HOST_SERVERS = [f'Haoran@{s}' for s in SERVERS]
elif CLOUD == "aws":
    SERVERS = ["3.82.245.250", "18.234.102.86", "54.211.138.210", "3.83.97.63", "3.83.107.40"]
    HOST_SERVERS = [f'ubuntu@{s}' for s in SERVERS]
    KEY_FILE = '/home/cavriola/.ssh/zhanghaorancs.pem'
    KEY_FILE = '/home/cavriola/.ssh/id_rsa'
else:
    exit(1)
