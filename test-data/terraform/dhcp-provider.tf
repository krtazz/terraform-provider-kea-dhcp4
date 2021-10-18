# define kea-dhcp4 provider
provider "kea-dhcp4" {
    kea_server_address    = "http://localhost:8080"
    kea_server_username   = "test"
    kea_server_password   = "1234"
    kea_server_configfile = "/etc/kea/kea-dhcp4.conf"
}
