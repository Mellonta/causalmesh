# CausalMesh Server

# Prerequisites
Clone both the server code (ccmesh) and the client code (ccmesh-go) in the same directory.

Install Python 3 and fabric
```
pip install fabric
```

# Setup
Use the following profile to create a cluster on Cloudlab
```
https://www.cloudlab.us/p/a22d79e1ff06310e24e133d03acc1703ec3b09ad
```

Copy the hostname of the nodes to `scripts/common.py`.

Run on the host
```
python scripts/upload.py
python scripts/setup.py
```

# Run experiments
To run Causalmesh, HydroCache-Con and HydroCache-Opt

change feature flags in `Cargo.toml` to
```
default = ["ccmesh", "cloud"]
```
```
default = ["con", "cloud"]
```
```
default = ["opt", "cloud"]
```
respectively, then run
```
python scripts/upload.py
```

All following commands should be run on the host.

## Microbenchmarks
```
python scripts/run.py --alg ccmesh --exp micro
python scripts/run.py --alg con --exp micro
python scripts/run.py --alg opt --exp micro
```

## Real world application
```
python scripts/run.py --alg ccmesh --exp media
python scripts/run.py --alg opt --exp media
```

## Various length of workflow
```
python scripts/run.py --alg ccmesh --exp scale_lambda
```

## Various number of cache server
```
python scripts/run.py --alg ccmesh --exp scale_server
```

## Visibility
```
python scripts/run.py --alg ccmesh --exp vis
```