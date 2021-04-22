print("Wallet Notifications")

load("ext://restart_process", "docker_build_with_restart")

cfg = read_yaml(
    "tilt.yaml",
    default = read_yaml("tilt.yaml.sample"),
)

local_resource(
    "notifications-build-binary",
    "make fast_build",
    deps = ["./cmd", "./internal", "./rpc/cmd", "./rpc/internal"],
)
local_resource(
    "notifications-generate-protpbuf",
    "make gen-protobuf",
    deps = ["./rpc/proto/notifications/notifications.proto"],
)

docker_build(
    "velmie/wallet-notifications-db-migration",
    ".",
    dockerfile = "Dockerfile.migrations",
    only = "migrations",
)
k8s_resource(
    "wallet-notifications-db-migration",
    trigger_mode = TRIGGER_MODE_MANUAL,
    resource_deps = ["wallet-notifications-db-init"],
)

wallet_notifications_options = dict(
    entrypoint = "/app/service_notifications",
    dockerfile = "Dockerfile.prebuild",
    port_forwards = [],
    helm_set = [],
)

if cfg["debug"]:
    wallet_notifications_options["entrypoint"] = "$GOPATH/bin/dlv --continue --listen :%s --accept-multiclient --api-version=2 --headless=true exec /app/service_notifications" % cfg["debug_port"]
    wallet_notifications_options["dockerfile"] = "Dockerfile.debug"
    wallet_notifications_options["port_forwards"] = cfg["debug_port"]
    wallet_notifications_options["helm_set"] = ["containerLivenessProbe.enabled=false", "containerPorts[0].containerPort=%s" % cfg["debug_port"]]

docker_build_with_restart(
    "velmie/wallet-notifications",
    ".",
    dockerfile = wallet_notifications_options["dockerfile"],
    entrypoint = wallet_notifications_options["entrypoint"],
    only = [
        "./build",
        "zoneinfo.zip",
    ],
    live_update = [
        sync("./build", "/app/"),
    ],
)
k8s_resource(
    "wallet-notifications",
    resource_deps = ["wallet-notifications-db-migration"],
    port_forwards = wallet_notifications_options["port_forwards"],
)

yaml = helm(
    "./helm/wallet-notifications",
    # The release name, equivalent to helm --name
    name = "wallet-notifications",
    # The values file to substitute into the chart.
    values = ["./helm/values-dev.yaml"],
    set = wallet_notifications_options["helm_set"],
)

k8s_yaml(yaml)
