{ pkgs }: {
    deps = [
        pkgs.haskellPackages.concurrent-dns-cache
        pkgs.wget
        pkgs.go_1_17
        pkgs.gopls
    ];
}