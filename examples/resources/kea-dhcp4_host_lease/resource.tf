resource "kea-dhcp4_host_lease" "some_host" {
    name = "some_host.local"
    mac_address = "00:11:22:33:44:55"
    ip_address = "10.0.0.1"
}
