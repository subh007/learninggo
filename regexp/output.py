"""
Remote platform
This platform uses physical ethernet interfaces.
"""

# Update this dictionary to suit your environment.
remote_port_map = {
update lineupdate lineupdate lineupdate line}

def platform_config_update(config):
    """
    Update configuration for the remote platform
    @param config The configuration dictionary to use/update
    """
    global remote_port_map
    config["port_map"] = remote_port_map.copy()
    config["caps_table_idx"] = 0
