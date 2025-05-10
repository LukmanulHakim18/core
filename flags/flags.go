package flags

//bb_discovery_mode
const BB_DISCOVERY_ENV_NAME = "bb_discovery_mode"
const BB_DISCOVERY_MODE_ETCD = "etcd"
const BB_DISCOVERY_MODE_ZK = "zk"

const SERVICE_PATH = "/service/"
const REGISTRY_NODE = "registry"
const ETCD_GLOBALS_CONFIG_PATH = "/globals/"
const ZK_GLOBALS_CONFIG_PATH = "/globals"

//monitor Prometheus
const BB_PROMETHEUS_NAMESPACE = "BBService"
const BB_PROMETHEUS_SWITCH_COUNTER_ON = "Counter"
const BB_PROMETHEUS_SWITCH_Histogram_ON = "Histogram"
const BB_PROMETHEUS_SWITCH_ALL_ON = "Enable"
const BB_PROMETHEUS_SWITCH_ALL_OFF = "Disable"
