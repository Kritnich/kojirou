{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = {self, nixpkgs, flake-utils }: 
    flake-utils.lib.eachDefaultSystem (system:
      let 
        pkgs = import nixpkgs { inherit system; };
      in {
        packages = rec {
          default = kojirou;
          kojirou = with pkgs; buildGoModule rec {
            name = "kojirou";
            version = "0.3.0";
            src = fetchFromGitHub {
              owner = "leotaku";
              repo = "kojirou";
              rev = "v${version}";
              sha256 = "sha256-kqdvqIbt/p7VksRSZw6ZoZqlwtlnFhfWcyJHXp8+BfU=";
            };
            proxyVendor = true;
            preBuild = ''
              go mod edit -go=1.20
              go mod tidy
            '';
            vendorSha256 = "sha256-/1LQg186DIm70gnQD6/RQiUhSsXfzOaTv1b6VDojTNU=";
          };
        };
        devShell = import ./shell.nix { inherit pkgs; };
      });
}