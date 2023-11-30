const avalanche = {
  id: 43114,
  name: "avalanche",
  network: "avalanche",
  nativeCurrency: {
    decimals: 18,
    name: "AVAX",
    symbol: "AVAX",
  },
  rpcUrls: {
    public: { http: ["https://api.avax.network/ext/bc/C/rpc"] },
    default: { http: ["https://api.avax.network/ext/bc/C/rpc"] },
  },
  blockExplorers: {
    default: { name: "Snowtrace", url: "https://snowtrace.io/" },
  },
  testnet: true,
};

const fuji = {
  id: 43113,
  name: "fuji",
  network: "fuji",
  nativeCurrency: {
    decimals: 18,
    name: "AVAX",
    symbol: "AVAX",
  },
  rpcUrls: {
    public: { http: ["https://api.avax-test.network/ext/bc/C/rpc"] },
    default: { http: ["https://api.avax-test.network/ext/bc/C/rpc"] },
  },
  blockExplorers: {
    default: { name: "Snowtrace", url: "https://testnet.snowtrace.dev/" },
  },
  testnet: true,
};

const anvil = {
  id: 31337,
  name: "anvil",
  network: "anvil",
  nativeCurrency: {
    decimals: 18,
    name: "AVAX",
    symbol: "AVAX",
  },
  rpcUrls: {
    public: { http: ["http://localhost:9650"] },
    default: { http: ["http://localhost:9650"] },
  },
  blockExplorers: {
    default: { name: "Blockscout", url: "https://todo.com" },
  },
  testnet: true,
};

const fork = {
  id: 43114,
  name: "anvil",
  network: "anvil",
  nativeCurrency: {
    decimals: 18,
    name: "AVAX",
    symbol: "AVAX",
  },
  rpcUrls: {
    public: { http: ["http://localhost:9650"] },
    default: { http: ["http://localhost:9650"] },
  },
  blockExplorers: {
    default: { name: "Blockscout", url: "https://todo.com" },
  },
  testnet: true,
};

const chains = { avalanche, fuji, anvil, fork };

export { chains };
