module.exports = {
  networks: {
    development: {
      host: "localhost",
      port: 9545,
      network_id: "*"
    },
    rinkeby: {
      host: "localhost", 
      port: 8545,
      from: "0xa46603651a444307c4a36dccce939ed922a51173",
      network_id: 4,
      gas: 470000
    }
  }
};